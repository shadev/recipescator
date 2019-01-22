package repository

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/shadev/recipescator/model"
	"log"
)

type Repository interface {
	FindAll() ([]*model.Recipe, error)
}

type MongoRepo struct {
	Client     *mongo.Client
	Db         string
	Collection string
}

func (repo *MongoRepo) FindAll() ([]*model.Recipe, error) {

	collection := repo.Client.Database(repo.Db).Collection(repo.Collection)

	cursor, e := collection.Find(context.Background(), bson.M{})

	if e != nil {
		log.Println("Could not get the collection from the database; ", e)
		return nil, e
	}
	var recipes []*model.Recipe

	for cursor.Next(context.Background()) {
		var recipe model.Recipe
		e := cursor.Decode(&recipe)

		if e != nil {
			log.Println("Could not decode recipe; ", e)
			return nil, e
		}
		recipes = append(recipes, &recipe)
	}

	e = cursor.Close(context.Background())
	if e != nil {
		log.Println("Could not close cursor; ", e)
	}

	return recipes, nil
}
