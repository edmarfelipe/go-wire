// +build wireinject

package main

import (
	"github.com/edmarfelipe/go-wire/internal/drivers/db"
	"github.com/edmarfelipe/go-wire/internal/drivers/http"
	dbInterfaces "github.com/edmarfelipe/go-wire/internal/interfaces/repositories/db"
	httpInterfaces "github.com/edmarfelipe/go-wire/internal/interfaces/repositories/http"
	"github.com/edmarfelipe/go-wire/internal/usecases"
	"github.com/google/wire"
)

func InitializeUsecase() (*usecases.UpdateLatestLaunchsUsecase, error) {
	wire.Build(
		http.NewLaunchRepository,
		db.NewLaunchRepository,
		wire.Bind(new(httpInterfaces.LaunchRepository), new(*http.LaunchRepository)),
		wire.Bind(new(dbInterfaces.LaunchRepository), new(*db.LaunchRepository)),
		usecases.NewUpdateLatestLaunchsUsecase,
	)

	return nil, nil
}
