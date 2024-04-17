package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (a *Handler) SelectFilm(writer http.ResponseWriter, request *http.Request) {
	var queryParam string
	urlPath := request.URL.Path
	for i := 1; i < len(urlPath); i++ {
		if urlPath[i] == '/' {
			queryParam = urlPath[i+1 : len(urlPath)]
		}
	}
	queryid, err := strconv.ParseInt(queryParam, 10, 64)
	f, err := a.Serv.SelectFilm(request.Context(), queryid)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ans, err := json.Marshal(f)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	writer.Write(ans)
}
