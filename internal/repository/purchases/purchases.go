package purchases

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Repo for purchase create.
type Repo struct {
	pool   *pgxpool.Pool
	logger *logrus.Logger
}

// New repo for merch create.
func New(pool *pgxpool.Pool, logger *logrus.Logger) Repo {
	return Repo{
		pool:   pool,
		logger: logger,
	}
}
