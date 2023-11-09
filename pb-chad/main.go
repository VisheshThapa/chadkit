package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var app *pocketbase.PocketBase

func main() {
	app = pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodPost,
			Path:    "/api/stripe-webhook",
			Handler: StripePaymentHandler,
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
