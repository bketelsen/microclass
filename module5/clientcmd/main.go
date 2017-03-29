package main

import (
	"fmt"
	"time"
	"context"
	"github.com/bketelsen/microclass/module4/userservice/proto/account"
	"github.com/micro/go-micro/client"
)

func main() {
	cl := account.NewAccountClient("go.micro.srv.user",client.DefaultClient)
	req := &account.LoginRequest{
		Username: "gopher",
		Password: "password1",
	}
	var rsp *account.LoginResponse
	var err error

	ctx, cxl:= context.WithTimeout(context.Background(),1*time.Second)
	defer cxl()
	// Call service
	if rsp, err = cl.Login(ctx,req); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("response:", rsp.String())
}
