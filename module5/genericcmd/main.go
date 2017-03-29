package main

import (
	"context"
	"fmt"

	"github.com/bketelsen/microclass/module4/userservice/proto/account"
	"github.com/micro/go-micro/client"
)

func main() {
	// Create new request to service go.micro.srv.user, method Account.Login
	req := client.NewRequest("go.micro.srv.user", "Account.Login", &account.LoginRequest{
		Username: "gopher",
		Password: "password1",
	})

	rsp := &account.LoginResponse{}

	// Call service
	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("response:", rsp.String())
}
