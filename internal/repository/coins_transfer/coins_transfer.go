package coinstransfer

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Repo for transfers create.
type Repo struct {
	pool   *pgxpool.Pool
	logger *logrus.Logger
}

// New repo for transfers create.
func New(pool *pgxpool.Pool, logger *logrus.Logger) Repo {
	return Repo{
		pool:   pool,
		logger: logger,
	}
}