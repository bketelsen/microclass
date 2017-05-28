package actions

import (
	"context"
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/micro/go-micro/client"
	"github.com/micro/user-srv/proto/account"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// AuthHandler attempts authorization against the login microservice with the generic client
func AuthHandler(c buffalo.Context) error {
	// Create new request to service go.micro.srv.user, method Account.Login
	req := client.NewRequest("go.micro.srv.user", "Account.Login", &account.LoginRequest{
		Username: "bketelsen",
		Password: "ncc1701c",
	})

	rsp := &account.LoginResponse{}

	// Call service
	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return c.Render(400, r.JSON(rsp))
	}

	fmt.Println("rsp:", rsp.Session)

	return c.Render(200, r.JSON(rsp))
}

// Auth2Handler attempts authorization against the login microservice with the generated client
func Auth2Handler(c buffalo.Context) error {
	cl := account.NewAccountClient("go.micro.srv.user", client.DefaultClient)
	req := &account.LoginRequest{
		Username: "bketelsen",
		Password: "ncc1701c",
	}
	fmt.Println(cl)
	rsp, err := cl.Login(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return c.Render(400, r.JSON(rsp))
	}
	return c.Render(200, r.JSON(rsp))
}
