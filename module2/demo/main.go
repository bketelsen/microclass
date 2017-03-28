package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/bketelsen/microclass/module2/demo/handler"
	"github.com/bketelsen/microclass/module2/demo/subscriber"

	example "github.com/bketelsen/microclass/module2/demo/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.brianketelsen.srv.demo"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	example.RegisterSubscriber("topic.com.brianketelsen.srv.demo", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	example.RegisterSubscriber("topic.com.brianketelsen.srv.demo", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
