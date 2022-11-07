package controller

import (
	"crypto/rsa"

	"github.com/Gurpreet-Bacancy/bcl/dbconn"
	"github.com/Gurpreet-Bacancy/bcl/postgres"
)

type Application struct {
	db         *dbconn.Postgres
	models     *postgres.Models
	privateKey *rsa.PrivateKey
}

func New(db *dbconn.Postgres, models *postgres.Models, privateKey *rsa.PrivateKey) *Application {
	app := &Application{
		db:     db,
		models: models,
	}
	if privateKey != nil {
		app.privateKey = privateKey
	}
	return app
}
