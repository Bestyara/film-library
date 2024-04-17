//go:generate mockgen -source ./handler_model.go -destination=./handler_mocks_test.go -package=handlers
package handlers

import "film-library/internal/service"

type Handler struct {
	Serv *service.FilmServ
}
