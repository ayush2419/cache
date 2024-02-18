package main

import (
	"github.com/ayush2419/cache/internal/cache"
	"github.com/ayush2419/cache/pkg"
	"github.com/ayush2419/cache/server"
)

func main() {
	serverConfig := server.ServerConfig{
		ListenAddr: ":3000",
		IsLeader:   true,
	}


	server := server.NewServer(serverConfig, pkg.NewHandler(cache.NewCache()))
	server.Start()
}
