package api

import (
	"film-library/internal/api/handlers"
	"film-library/internal/api/middleware"
	"film-library/internal/api/router"
	"film-library/internal/config"
	"net/http"
)

func Run(c config.Config, m chan error, impl handlers.Handler) {
	router := router.CreateRouter(impl)
	http.Handle("/", middleware.AuthMiddleware(router))
	go func() {
		if err := http.ListenAndServe(c.Server.Port, nil); err != nil {
			m <- err
		}
	}()
	go func() {
		if err := http.ListenAndServeTLS(c.Server.SecurePort, "server.crt", "server.key", nil); err != nil {
			m <- err
		}
	}()
}
