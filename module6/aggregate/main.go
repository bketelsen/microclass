package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bketelsen/microclass/module4/userservice/proto/account"
	"github.com/micro/go-micro/client"
	"github.com/micro/profile-srv/proto/record"
)

func main() {
	cl := account.NewAccountClient("go.micro.srv.user", client.DefaultClient)
	req := &account.LoginRequest{
		Username: "gopher",
		Password: "password1",
	}
	var rsp *account.LoginResponse
	var err error

	ctx, cxl := context.WithTimeout(context.Background(), 1*time.Second)
	defer cxl()
	// Call service
	if rsp, err = cl.Login(ctx, req); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Found account:", rsp.Session.GetUsername())

	sreq := &account.SearchRequest{
		Username: "gopher",
		Limit:    10,
	}
	var srsp *account.SearchResponse
	if srsp, err = cl.Search(ctx, sreq); err != nil {
		fmt.Println("call err:", err, srsp)
	}
	var foundUser *account.User
	if srsp.Users != nil {
		if len(srsp.Users) > 0 {
			foundUser = srsp.Users[0]
		} else {
			fmt.Println("no user found")
			return
		}
	} else {
		fmt.Println("Got no users")
		return
	}

	profcl := record.NewRecordClient("go.micro.srv.profile", client.DefaultClient)

	profreq := &record.ReadRequest{
		Id: foundUser.GetId(),
	}
	var profrsp *record.ReadResponse
	if profrsp, err = profcl.Read(ctx, profreq); err != nil {
		fmt.Println("call err:", err, profrsp)
		return
	}
	fmt.Println(profrsp.Profile)

}
