package main

import (
	"github.com/Pelissaroo/api-students/api"
	_ "github.com/Pelissaroo/api-students/docs"
	"github.com/rs/zerolog/log"
)

func main() {
	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

}
