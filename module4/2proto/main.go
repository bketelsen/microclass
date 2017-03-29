package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/bketelsen/microclass/module4/2proto/handler"
	"github.com/bketelsen/microclass/module4/2proto/subscriber"

	example "github.com/bketelsen/microclass/module4/2proto/proto/account"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.generate"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	service.Server().Subscribe(
		service.Server().NewSubscriber("topic.go.micro.srv.generate", new(subscriber.Example)),
	)

	// Register Function as Subscriber
	service.Server().Subscribe(
		service.Server().NewSubscriber("topic.go.micro.srv.generate", subscriber.Handler),
	)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
