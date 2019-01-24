package rest

import (
	"github.com/labstack/echo"
	"github.com/shadev/recipescator/repository"
	"net/http"
)

type Endpoint struct {
	Repo repository.Repository
}

func (ep *Endpoint) GetAllRecipes(context echo.Context) error {
	recipes, e := ep.Repo.FindAll()
	if e != nil {
		return context.String(http.StatusInternalServerError, e.Error())
	}
	return context.JSONPretty(http.StatusOK, recipes, " ")
}

func (ep *Endpoint) GetSingleRecipe(context echo.Context) error {
	recipe, e := ep.Repo.FindOne(context.Param("rid"))
	if e != nil {
		return context.String(http.StatusInternalServerError, e.Error())
	}
	if recipe == nil {
		return context.String(http.StatusNotFound, "")
	}
	return context.JSONPretty(http.StatusOK, recipe, " ")
}
