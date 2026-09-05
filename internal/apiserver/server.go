package apiserver

import (
	"avito/internal/repository/users"
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
	router    *mux.Router
	usersRepo *users.Repo
	logger    *logrus.Logger
}

func newServer(router *mux.Router, usersRepo *users.Repo, logger *logrus.Logger) *server {
	srv := &server{
		router:    router,
		usersRepo: usersRepo,
		logger:    logger,
	}

	srv.configureRouter()

	return srv
}

func (s *server) configureRouter() {
}
