package main

import (
	"log"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/siadat/swagger-eg/swagger/restapi"
	"github.com/siadat/swagger-eg/swagger/restapi/operations"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTestmanAPI(swaggerSpec)

	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.PostKeysKeyHandler = operations.PostKeysKeyHandlerFunc(f)
	api.ServerShutdown = func() {}
	api.Logger = log.Printf

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8001
	log.Fatal(server.Serve())
}

func f(params operations.PostKeysKeyParams) middleware.Responder {
	return operations.NewPostKeysKeyOK()
}
