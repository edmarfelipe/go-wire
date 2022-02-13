package http

import (
	"github.com/edmarfelipe/go-wire/internal/interfaces/models"
	"github.com/go-resty/resty/v2"
)

const URL = "https://ll.thespacedevs.com/2.2.0/launch/?format=json&search=SpaceX"

type LaunchData struct {
	Results []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Failreason string `json:"failreason"`
	} `json:"results"`
}

type LaunchRepository struct {
	http *resty.Client
}

func NewLaunchRepository() *LaunchRepository {
	return &LaunchRepository{
		http: resty.New(),
	}
}

func (rep LaunchRepository) GetAll() ([]models.Launch, error) {
	result := new(LaunchData)
	lauchs := make([]models.Launch, 0)

	_, err := rep.http.R().
		SetResult(result).
		Get(URL)

	if err != nil {
		return lauchs, err
	}

	for _, item := range result.Results {
		lauchs = append(lauchs, models.Launch{
			Id:         item.ID,
			Name:       item.Name,
			FailReason: item.Failreason,
		})
	}

	return lauchs, nil
}
