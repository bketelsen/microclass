package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/bketelsen/microclass/module4/5db/handler"

	account "github.com/bketelsen/microclass/module4/5db/proto/account"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Register Handler
	account.RegisterAccountHandler(service.Server(),new(handler.Account))


	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
