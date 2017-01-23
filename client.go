package main

import (
	"context"
	"fmt"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/siadat/swagger-eg/swagger/client"
	"github.com/siadat/swagger-eg/swagger/client/operations"

	"github.com/go-openapi/strfmt"
)

func main() {
	c := newSwaggerClient()
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	_, err := c.Operations.PostKeysKey(&operations.PostKeysKeyParams{
		Key:     "k",
		Value:   "v",
		Context: ctx,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func newSwaggerClient() *client.Testman {
	rt := httptransport.New(
		"localhost:8001",
		client.DefaultBasePath,
		client.DefaultSchemes,
	)
	rt.Debug = true
	return client.New(rt, strfmt.NewFormats())
}
