package api

import (
	"encoding/json"
	"film-library/internal/model"
	"film-library/internal/service"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Api struct {
	Serv *service.FilmServ
}

func CreateRouter(impl Api) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			impl.AddFilm(writer, request)
		default:
			log.Println("wrong method")
		}
	})
	router.HandleFunc("/sort", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			impl.SortFilms(writer, request)
		default:
			log.Println("wrong method")
		}
	})
	router.HandleFunc("/delete", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodDelete:
			impl.DeleteFilm(writer, request)
		default:
			log.Println("wrong method")
		}
	})
	router.HandleFunc("/select", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			impl.SelectFilm(writer, request)
		default:
			log.Println("wrong method")
		}
	})
	router.HandleFunc("/update", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPatch:
			impl.UpdateFilm(writer, request)
		case http.MethodPut:
			impl.UpdateFilm(writer, request)
		default:
			log.Println("wrong method")
		}
	})
	return router
}

func (a *Api) AddFilm(writer http.ResponseWriter, request *http.Request) {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	var film model.Film
	if err := json.Unmarshal(data, &film); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	id, err := a.Serv.AddFilm(request.Context(), film)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	filmwrite := model.Film{
		ID:          id,
		Name:        film.Name,
		Description: film.Description,
		Rating:      film.Rating,
		ReleaseDate: film.ReleaseDate,
	}

	ans, err := json.Marshal(filmwrite)
	writer.Write(ans)
}

func (a *Api) DeleteFilm(writer http.ResponseWriter, request *http.Request) {

}

func (a *Api) SelectFilm(writer http.ResponseWriter, request *http.Request) {

}

func (a *Api) SortFilms(writer http.ResponseWriter, request *http.Request) {

}

func (a *Api) UpdateFilm(writer http.ResponseWriter, request *http.Request) {

}
