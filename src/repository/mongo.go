package repository

import "github.com/shadev/recipescator/src/model"

type Repository interface {
	FindAll() ([]model.Recipe, error)
}

type MongoRepo struct {
}

func (repo *MongoRepo) FindAll() ([]model.Recipe, error) {
	return []model.Recipe{}, nil
}
