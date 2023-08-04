package server

import (
	"github.com/micro/go-micro/v3/api"
	"github.com/micro/go-micro/v3/api/handler/api"
	"github.com/micro/go-micro/v3/api/router"
	"github.com/micro/go-micro/v3/api/router/gorilla"
	"github.com/micro/go-micro/v3/util/log"
	"github.com/micro/go-micro/v3/web"
	"github.com/micro/go-micro/v3/web/handler"
	"github.com/micro/go-micro/v3/web/router/static"
	"github.com/micro/go-micro/v3/errors"
	"net/http"
)

func Run() {
	// Create a new service
	service := web.NewService(
		web.Name("go.micro.api.supplychain"),
		web.Version("latest"),
	)

	// Initialize service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// Create API handler
	apiHandler := api.NewHandler(
		api.WithNamespace(apiNamespace),
		api.WithRouter(gorilla.NewRouter(router.WithHandler(handler.Default))),
	)

	// Register API handler
	if err := apiHandler.Register(apiEndpoint); err != nil {
		log.Fatal(err)
	}

	// Create a new static file router
	staticRouter := static.NewRouter(
		static.WithHandler(http.FileServer(http.Dir("html"))),
		static.WithPrefix("/html"),
	)

	// Register static file handler
	if err := service.Handle(staticRouter); err != nil {
		log.Fatal(err)
	}

	// Run the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
