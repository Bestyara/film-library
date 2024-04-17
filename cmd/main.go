package main

import (
	"context"
	"film-library/internal/api"
	"film-library/internal/api/handlers"
	"film-library/internal/config"
	"film-library/internal/db"
	"film-library/internal/repository"
	"film-library/internal/service"
	"log"
)

func main() {
	c, err := config.ConfigInit()
	if err != nil {
		log.Println(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database, err := db.NewDb(c, ctx)
	if err != nil {
		log.Println(err)
	}
	defer database.GetPool(ctx).Close()

	r := repository.NewRepository(database)
	s := service.NewService(r)
	impl := handlers.Handler{Serv: s}

	m := make(chan error)
	api.Run(c, m, impl)
	select {
	case err := <-m:
		if err != nil {
			log.Println(err)
		}
	}
}
