package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	// server.Default() creates a Hertz with recovery middleware.
	// If you need a pure hertz, you can use server.New()
	h := server.Default()

	RegisterRoute(h)

	h.Spin()
}
