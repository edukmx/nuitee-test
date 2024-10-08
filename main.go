package main

import (
	"net/http"

	ioc "github.com/edukmx/nuitee/container"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		panic("cannot load .env file")
	}

	container := ioc.BuildContainer()

	err = container.Invoke(func(server *http.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
