package main

import (
	ioc "github.com/edukmx/nuitee/container"
	"github.com/edukmx/nuitee/httpx"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		panic("cannot load .env file")
	}

	container := ioc.BuildContainer()

	err = container.Invoke(func(server *httpx.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
