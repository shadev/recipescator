package main

import (
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/shadev/recipescator/repository"
	"github.com/shadev/recipescator/rest"
	"log"
	"os"
)

func main() {
	e := initEcho()

	endpoint := rest.Endpoint{Repo: mongoRepo()}

	e.GET("/recipes", endpoint.GetAllRecipes)

	serverError := e.Start(":1323")
	e.Logger.Fatal(serverError)
}

func initEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.RemoveTrailingSlash())
	return e
}

func mongoRepo() *repository.MongoRepo {
	client, err := mongo.Connect(context.Background(), os.Getenv("RECIPESCATOR_MONGO_URL"))
	if err != nil {
		log.Fatal("Could not connect to MongoDB ", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB; ", err)
	}
	return &repository.MongoRepo{Client: client, Db: "recipescator-db", Collection: "recipes"}
}
