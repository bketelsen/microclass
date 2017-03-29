package main

import (
	"log"

	"github.com/bketelsen/microclass/module4/4main/handler"
	"github.com/micro/go-micro"

	account "github.com/bketelsen/microclass/module4/4main/proto/account"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Register Handler
	account.RegisterAccountHandler(service.Ser ver(), new(handler.Account))

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
