package apiserver

import (
	"avito/internal/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	defaultReadTimeout       = 5 * time.Second
	defaultReadHeaderTimeout = 6 * time.Second
	defaultWriteTimeout      = 10 * time.Second
	defaultIdleimeout        = 7 * time.Second
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

func newServer() *server {
	srv := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

	srv.configureRouter()

	return srv
}

func (s *server) configureRouter() {
}

func (s *server) Registration(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "error with decode", http.StatusBadRequest)
			return
		}

		u := model.User{
			Email:        req.Email,
			HashPassword: req.Password,
		}

		
	}
}
