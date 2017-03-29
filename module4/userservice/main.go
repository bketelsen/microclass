package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/cli"
	"github.com/bketelsen/microclass/module4/userservice/handler"
	"github.com/bketelsen/microclass/module4/userservice/db"
	account "github.com/bketelsen/microclass/module4/userservice/proto/account"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Flags(
			cli.StringFlag{
				Name:   "database_url",
				EnvVar: "DATABASE_URL",
				Usage:  "The database URL e.g root@tcp(127.0.0.1:3306)/user",
			},
		),

		micro.Action(func(c *cli.Context) {
			if len(c.String("database_url")) > 0 {
				db.Url = c.String("database_url")
			}
		}),
	)

	service.Init()
	db.Init()

	account.RegisterAccountHandler(service.Server(), new(handler.Account))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
