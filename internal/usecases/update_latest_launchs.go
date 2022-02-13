package usecases

import (
	"github.com/edmarfelipe/go-wire/internal/interfaces/repositories/db"
	"github.com/edmarfelipe/go-wire/internal/interfaces/repositories/http"
)

type UpdateLatestLaunchsUsecase struct {
	dbRepository   db.LaunchRepository
	httpRepository http.LaunchRepository
}

func NewUpdateLatestLaunchsUsecase(dbRepo db.LaunchRepository, httpRepo http.LaunchRepository) *UpdateLatestLaunchsUsecase {
	return &UpdateLatestLaunchsUsecase{
		dbRepository:   dbRepo,
		httpRepository: httpRepo,
	}
}

func (usecase UpdateLatestLaunchsUsecase) Update() error {
	launchs, err := usecase.httpRepository.GetAll()
	if err != nil {
		return err
	}

	for _, item := range launchs {
		err = usecase.dbRepository.Create(item)
		if err != nil {
			return err
		}
	}

	return nil
}
