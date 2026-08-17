package apiserver

import (
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
	s.router.HandleFunc("/", s.handleHello)
}

func (s *server) handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
