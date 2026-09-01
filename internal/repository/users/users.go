package users

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Repo for users create.
type Repo struct {
	pool   *pgxpool.Pool
	logger *logrus.Logger
}

// New repo for users create.
func New(pool *pgxpool.Pool, logger *logrus.Logger) Repo {
	return Repo{
		pool:   pool,
		logger: logger,
	}
}

