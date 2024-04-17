package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (a *Handler) DeleteFilm(writer http.ResponseWriter, request *http.Request) {
	var queryParam string
	urlPath := request.URL.Path
	for i := 1; i < len(urlPath); i++ {
		if urlPath[i] == '/' {
			queryParam = urlPath[i+1 : len(urlPath)]
		}
	}
	queryid, err := strconv.ParseInt(queryParam, 10, 64)
	_, err = a.Serv.DeleteFilm(request.Context(), queryid)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	fmt.Println("film has been deleted")
}
