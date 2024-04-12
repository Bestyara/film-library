package api

import (
	"encoding/json"
	"film-library/internal/model"
	"film-library/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Api struct {
	Serv *service.FilmServ
}

func CreateRouter(impl Api) *mux.Router {
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

func (a *Api) AddFilm(writer http.ResponseWriter, request *http.Request) {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
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
	//filmwrite := model.Film{
	//	Name:        film.Name,
	//	Description: film.Description,
	//	Rating:      film.Rating,
	//	ReleaseDate: film.ReleaseDate,
	//}
	//ans, err := json.Marshal(filmwrite)
	//writer.Write(ans)
}

func (a *Api) DeleteFilm(writer http.ResponseWriter, request *http.Request) {
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
		return
	}
	fmt.Println("film has been deleted")
}

func (a *Api) SelectFilm(writer http.ResponseWriter, request *http.Request) {
	//var queryParam string
	//urlPath := request.URL.Path
	//for i := 1; i < len(urlPath); i++ {
	//	if urlPath[i] == '/' {
	//		queryParam = urlPath[i+1 : len(urlPath)]
	//	}
	//}
	data, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var film model.Film
	if err := json.Unmarshal(data, &film); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	f, err := a.Serv.SelectFilm(request.Context(), film)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	ans, err := json.Marshal(f)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(ans)
}

func (a *Api) SortFilms(writer http.ResponseWriter, request *http.Request) {
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
	fmt.Println("Sorted list of films has been returned")
}

func (a *Api) UpdateFilm(writer http.ResponseWriter, request *http.Request) {

}
