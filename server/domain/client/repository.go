package client

import (
	"api/server/infrastructure/persistence/client"
	"api/server/infrastructure/persistence/client/postgres"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type repository struct {
	pg *postgres.PGClient
}

func newRepository(db *sql.DB) *repository {
	return &repository{
		pg: &postgres.PGClient{DB: db},
	}
}

func (r *repository) Create(ctx *gin.Context, c *client.Client) (err error) {
	return r.pg.Create(ctx, c)
}

func (r *repository) List(ctx *gin.Context) (clients []client.Clients, err error) {
	return r.pg.List(ctx)
}
