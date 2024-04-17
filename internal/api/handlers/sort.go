package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (a *Handler) SortFilms(writer http.ResponseWriter, request *http.Request) {
	var queryParam string
	urlPath := request.URL.Path
	for i := 1; i < len(urlPath); i++ {
		if urlPath[i] == '/' {
			queryParam = urlPath[i+1 : len(urlPath)]
		}
	}
	fList, err := a.Serv.SortFilms(request.Context(), queryParam)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	for i := 0; i < len(fList); i++ {
		d, err := json.Marshal(fList[i])
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		writer.Write(d)
	}
}
