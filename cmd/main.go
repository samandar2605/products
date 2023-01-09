package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/samandar2605/products/api"
	"github.com/samandar2605/products/config"
	"github.com/samandar2605/products/storage"
)

func main() {
	cfg := config.Load(".")
	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostConfig.Host,
		cfg.PostConfig.Port,
		cfg.PostConfig.User,
		cfg.PostConfig.Password,
		cfg.PostConfig.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	strg := storage.NewStoragePg(psqlConn)

	apiServer := api.New(&api.RouterOptions{
		Cfg:     &cfg,
		Storage: strg,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

	log.Print("Server stopped")

}
