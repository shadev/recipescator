package rest

import (
	"github.com/labstack/echo"
	"github.com/shadev/recipescator/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockRepo struct {
	mock.Mock
}

func (repo *MockRepo) FindAll() ([]model.Recipe, error) {
	args := repo.Called()
	return args.Get(0).([]model.Recipe), args.Error(1)
}

func TestGetAllRecipes_ok(t *testing.T) {
	resultAsBytes, _ := ioutil.ReadFile("../testresources/testGetAllRecipes_ok.json")

	expectedResult := string(resultAsBytes)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/recipes", nil)
	rec := httptest.NewRecorder()
	mockRepo := new(MockRepo)
	mockRepo.On("FindAll").Return(createRecipe(), nil)
	testee := Endpoint{mockRepo}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.GetAllRecipes(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, expectedResult, rec.Body.String())
	}
	mockRepo.AssertExpectations(t)
}

func createRecipe() []model.Recipe {
	recipes := []model.Recipe{
		{
			Rid:      "123456789",
			Title:    "Tiramisu Bars",
			Rating:   5,
			Servings: 4,
			Tags:     []string{"Dessert", "Quick", "Vegan"},
			Ingredients: []model.Ingredient{
				{Name: "Cashew", Amount: "1 Cup"},
				{Name: "Cocoa", Amount: "0.5 Cup"},
			},
			Source:     model.Source{SourceType: model.COOKBOOK, Title: "Richa's Everyday Kitchen", Ref: "Page 214"},
			Time:       model.Time{Active: 10 * time.Minute, Inactive: 1 * time.Hour, Prep: 10 * time.Minute},
			PreparedOn: []time.Time{time.Date(2018, 12, 25, 0, 0, 0, 0, time.UTC)},
			Comments:   []string{"Tastes great", "One of my favourite desserts"},
		},
		{
			Rid:      "abcdefgh",
			Title:    "Cholent",
			Rating:   5,
			Servings: 4,
			Tags:     []string{"Stew", "Quick", "Vegan"},
			Ingredients: []model.Ingredient{
				{Name: "TVP", Amount: "1 Cup"},
				{Name: "Onion", Amount: "1 large", Preparation: "Diced"},
			},
			Source:     model.Source{SourceType: model.INTERNET, Title: "Veganomicon", Ref: "https://postpunkkitchen.de/cholent"},
			Time:       model.Time{Active: 50 * time.Minute, Inactive: 40 * time.Minute, Prep: 20 * time.Minute},
			PreparedOn: []time.Time{time.Date(2019, 1, 6, 0, 0, 0, 0, time.UTC)},
			Comments:   []string{"Great winter dish", "So yummy"},
		},
	}
	return recipes
}
