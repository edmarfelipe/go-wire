package db

import (
	"github.com/edmarfelipe/go-wire/internal/interfaces/models"
	"github.com/guregu/dynamo"
)

type LaunchRepository struct {
	Table *dynamo.Table
}

func NewLaunchRepository() *LaunchRepository {
	return &LaunchRepository{
		Table: NewDynamoTable("launchs"),
	}
}

func (rep LaunchRepository) Create(model models.Launch) error {
	err := rep.Table.Put(model).Run()

	if err != nil {
		return err
	}

	return nil
}

func (rep LaunchRepository) FindOne(userId int) (models.Launch, error) {
	var result models.Launch

	err := rep.Table.Get("UserID", userId).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (rep LaunchRepository) FindAll() ([]models.Launch, error) {
	var results []models.Launch

	err := rep.Table.Scan().All(&results)
	if err != nil {
		return results, err
	}

	return results, nil
}
