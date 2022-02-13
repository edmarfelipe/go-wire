package http

import "github.com/edmarfelipe/go-wire/internal/interfaces/models"

type LaunchRepository interface {
	GetAll() ([]models.Launch, error)
}
