package main

import (
	"context"
	"film-library/internal/api"
	"film-library/internal/db"
	"film-library/internal/repository"
	"film-library/internal/service"
	"log"
	"net/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.NewDb(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool(ctx).Close()

	r := repository.NewRepository(database)
	s := service.NewService(r)
	impl := api.Api{Serv: s}

	http.Handle("/", api.CreateRouter(impl))
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}

}
