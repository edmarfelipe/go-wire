package db

import "github.com/edmarfelipe/go-wire/internal/interfaces/models"

type LaunchRepository interface {
	Create(model models.Launch) error
	FindAll() ([]models.Launch, error)
	FindOne(userId int) (models.Launch, error)
}
