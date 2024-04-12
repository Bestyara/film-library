package router

import (
	"film-library/internal/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateRouter(impl api.Api) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		//impl.AddFilm(writer, request)
		switch request.Method {
		case http.MethodPost:
			impl.AddFilm(writer, request)
		case http.MethodPut:
			impl.AddFilm(writer, request)
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
			log.Println("wrong method")
		}
	})
	router.HandleFunc("/sort/{(name|description|rating|releasedate|default)}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			impl.SortFilms(writer, request)
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
			log.Println("wrong method")
		}
	})
	router.HandleFunc("/delete/{/[0-9]+/}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodDelete:
			impl.DeleteFilm(writer, request)
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
			log.Println("wrong method")
		}
	})
	router.HandleFunc("/select/{/[0-9]+/}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			impl.SelectFilm(writer, request)
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
			log.Println("wrong method")
		}
	})
	router.HandleFunc("/update/{/[0-9]+/}", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPatch:
			impl.UpdateFilm(writer, request)
		case http.MethodPut:
			impl.UpdateFilm(writer, request)
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
			log.Println("wrong method")
		}
	})
	return router
}
