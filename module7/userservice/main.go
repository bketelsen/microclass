package main

import (
	"log"

	"github.com/bketelsen/microclass/module7/userservice/db"
	"github.com/bketelsen/microclass/module7/userservice/handler"
	account "github.com/bketelsen/microclass/module7/userservice/proto/account"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-os/event"
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
	ev := event.NewEvent()
	act := handler.NewAccount(ev)

	account.RegisterAccountHandler(service.Server(), act)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
