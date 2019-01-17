package rest

import (
	"github.com/labstack/echo"
	"github.com/shadev/recipescator/src/repository"
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
