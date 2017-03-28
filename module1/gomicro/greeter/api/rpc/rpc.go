package main

import (
	"log"

	proto "github.com/bketelsen/microclass/module1/gomicro/greeter/api/rpc/proto/hello"
	hello "github.com/bketelsen/microclass/module1/gomicro/greeter/srv/proto/hello"
	"github.com/micro/go-micro"

	"golang.org/x/net/context"
)

type Greeter struct {
	Client hello.SayClient
}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Greeter.Hello API request")

	// make the request
	response, err := g.Client.Hello(ctx, &hello.Request{Name: req.Name})
	if err != nil {
		return err
	}

	// set api response
	rsp.Msg = response.Msg
	return nil
}

func main() {
	// Create service
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
	)

	// Init to parse flags
	service.Init()

	// Register Handlers
	proto.RegisterGreeterHandler(service.Server(), &Greeter{
		// Create Service Client
		Client: hello.NewSayClient("go.micro.srv.greeter", service.Client()),
	})

	// for handler use

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
