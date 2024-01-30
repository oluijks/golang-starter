package main

import (
	"github.com/oluijks/golang-starter/server/api"
	"github.com/oluijks/golang-starter/server/api/handlers"
	"github.com/oluijks/golang-starter/server/internal/config"
	"github.com/oluijks/golang-starter/server/internal/database"
	"github.com/oluijks/golang-starter/server/internal/storage"
	"github.com/oluijks/golang-starter/server/internal/token"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.NewConfig,
			api.NewGinEngine,
			storage.NewDBStore,
			token.NewPasetoTokenMaker,
			database.NewMongoConnection,
		),
		fx.Provide(
			handlers.NewPingHandlers,
			handlers.NewAuthHandlers,
			handlers.NewAccountHandlers,
		),
		fx.Invoke(api.NewGinEngine),
	).Run()
}
