package handlers

import (
	"encoding/json"
	"film-library/internal/model"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (a *Handler) AddFilm(writer http.ResponseWriter, request *http.Request) {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var film model.Film
	if err := json.Unmarshal(data, &film); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	_, err = a.Serv.AddFilm(request.Context(), film)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	fmt.Println("film has been added")
}
