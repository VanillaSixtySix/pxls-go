package model

type ServerConfig struct {
	HttpServer HttpServerConfig `toml:"httpserver"`
	Database   DatabaseConfig   `toml:"database"`
}

type HttpServerConfig struct {
	Address string `toml:"address"`
}

type DatabaseConfig struct {
	Url string `toml:"url"`
}
