package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {

	listenAddr := flag.String("listenAddr", ":8080", "Listen address of API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1/")

	app.Listen(*listenAddr)

}
