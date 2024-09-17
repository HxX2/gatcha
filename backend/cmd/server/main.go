package main

import (
	"gatcha/internal/config"
	"gatcha/internal/router"
	"gatcha/internal/server"
)

func main() {
	cfg := config.Load()
	r := router.NewRouter(cfg)
	s := server.NewServer(cfg, r)
	s.Start()
}
