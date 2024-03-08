package main

import (
	"context"
	"log"
	"os"

	"pxls-go/internal/httpserver"
	"pxls-go/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/pelletier/go-toml/v2"
)

var Config model.ServerConfig

func main() {
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatalln("Could not load config.toml:", err)
	}

	err = toml.Unmarshal(configFile, &Config)
	if err != nil {
		log.Fatalln("Could not unmarshal config.toml:", err)
	}

	conn, err := pgx.Connect(context.Background(), Config.Database.Url)
	if err != nil {
		log.Fatalln("Could not connect to database:", err)
	}
	defer conn.Close(context.Background())

	e := echo.New()
	e.HideBanner = true
	httpserver.RegisterHandlers(e)
	e.Logger.Fatal(e.Start(Config.HttpServer.Address))
}
