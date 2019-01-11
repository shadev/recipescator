package rest

import (
	"github.com/labstack/echo"
	"github.com/shadev/recipescator/model"
	"github.com/shadev/recipescator/repository"
	"log"
	"net/http"
)

// RecipeEndpoint provides CRUD functionality for recipes
type RecipeEndpoint struct {
	Repo repository.Repository
}

// GetAllRecipes returns all recipes found
func (ep *RecipeEndpoint) GetAllRecipes(context echo.Context) error {
	recipes, e := ep.Repo.FindAll()
	if e != nil {
		log.Println(e)
		return context.NoContent(http.StatusInternalServerError)
	}
	return context.JSONPretty(http.StatusOK, recipes, " ")
}

// GetSingleRecipe returns a single recipe, selected by recipeId (rid) if found
func (ep *RecipeEndpoint) GetSingleRecipe(context echo.Context) error {
	recipe, e := ep.Repo.FindOne(context.Param("rid"))
	if e != nil {
		log.Println(e)
		return context.NoContent(http.StatusInternalServerError)
	}
	if recipe == nil {
		return context.NoContent(http.StatusNotFound)
	}
	return context.JSONPretty(http.StatusOK, recipe, " ")
}

// PostNewRecipe inserts a new recipe in the repository and returns 201 Created with Location-Header on success
func (ep *RecipeEndpoint) PostNewRecipe(context echo.Context) error {
	var newRecipe model.Recipe
	e := context.Bind(&newRecipe)

	if e != nil {
		log.Println(e)
		return context.NoContent(http.StatusBadRequest)
	}
	rid, e := ep.Repo.Insert(newRecipe)

	if e != nil {
		log.Println(e)
		return context.NoContent(http.StatusInternalServerError)
	}

	context.Response().Header().Add("Location", "/recipes/"+rid)
	return context.NoContent(http.StatusCreated)
}
