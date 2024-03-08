package main

import (
	"pxls-go/internal/httpserver"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	httpserver.RegisterHandlers(e)
	e.Logger.Fatal(e.Start(":4567"))
}
