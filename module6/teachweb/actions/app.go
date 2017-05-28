package actions

import (
	"github.com/gobuffalo/buffalo"

	"github.com/bketelsen/microclass/module6/teachweb/models"
	"github.com/gobuffalo/buffalo/middleware"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_teachweb_session",
		})

		app.Use(middleware.PopTransaction(models.DB))

		app.GET("/", HomeHandler)
		app.GET("/auth", AuthHandler)
		app.GET("/auth2", Auth2Handler)
		app.ServeFiles("/assets", packr.NewBox("../assets"))
	}

	return app
}
