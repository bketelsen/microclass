package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/bketelsen/microclass/module4/1generate/handler"
	"github.com/bketelsen/microclass/module4/1generate/subscriber"

	example "github.com/bketelsen/microclass/module4/1generate/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.1generate"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	service.Server().Subscribe(
		service.Server().NewSubscriber("topic.go.micro.srv.1generate", new(subscriber.Example)),
	)

	// Register Function as Subscriber
	service.Server().Subscribe(
		service.Server().NewSubscriber("topic.go.micro.srv.1generate", subscriber.Handler),
	)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
