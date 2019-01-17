package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shadev/recipescator/src/repository"
	"github.com/shadev/recipescator/src/rest"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.RemoveTrailingSlash())

	repo := new(repository.MongoRepo)

	endpoint := &rest.Endpoint{Repo: repo}

	e.GET("/recipes", endpoint.GetAllRecipes)

	serverError := e.Start(":1323")
	e.Logger.Fatal(serverError)
}
