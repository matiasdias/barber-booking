package service

import (
	"api/server/infrastructure/persistence/service"
	"api/server/infrastructure/persistence/service/postgres"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type repository struct {
	pg *postgres.PGService
}

func newRepository(db *sql.DB) *repository {
	return &repository{
		pg: &postgres.PGService{DB: db},
	}
}

func (r *repository) Create(ctx *gin.Context, s *service.Services) (err error) {
	return r.pg.Create(ctx, s)
}

func (r *repository) List(ctx *gin.Context) (services []service.ListService, err error) {
	return r.pg.List(ctx)
}
