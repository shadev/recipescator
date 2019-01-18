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
}

func (repo *MongoRepo) FindAll() ([]*model.Recipe, error) {

	client, e := mongo.Connect(context.TODO(), "<url>")

	if e != nil {
		log.Fatal("Could not connect to MongoDB; ", e)
		return nil, e
	}
	e = client.Ping(context.TODO(), nil)

	if e != nil {
		log.Fatal("Could not ping MongoDB; ", e)
		return nil, e
	}

	collection := client.Database("recipescator-db").Collection("recipes")

	cursor, e := collection.Find(context.TODO(), bson.M{})

	if e != nil {
		log.Fatal("Could not get the collection from the database; ", e)
		return nil, e
	}
	var recipes []*model.Recipe

	for cursor.Next(context.TODO()) {
		var recipe model.Recipe
		e := cursor.Decode(&recipe)

		if e != nil {
			log.Fatal("Could not decode recipe; ", e)
			return nil, e
		}
		recipes = append(recipes, &recipe)
	}

	e = cursor.Close(context.TODO())
	if e != nil {
		log.Fatal("Could not close cursor; ", e)
	}

	return recipes, nil
}
