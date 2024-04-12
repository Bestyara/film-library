package middleware

import (
	"errors"
	"film-library/internal/config"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if b, _ := isAuthorised(username, password); !b {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}

func isAuthorised(login string, password string) (bool, error) {
	y, err := config.ConfigInit()
	if err != nil {
		return false, errors.New("can not get config parameters")
	}
	for _, val := range y.Users {
		if val.Login == login && val.Password == password {
			return true, nil
		}
	}
	return false, errors.New("not authorized")
}
