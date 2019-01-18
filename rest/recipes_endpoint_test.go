package rest

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/shadev/recipescator/model"
	"github.com/shadev/recipescator/testresources"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRepo struct {
	mock.Mock
}

func (repo *MockRepo) FindAll() ([]*model.Recipe, error) {
	args := repo.Called()
	recipes, ok := args.Get(0).([]*model.Recipe)
	e := args.Error(1)

	if ok {
		return recipes, e
	} else {
		return nil, e
	}
}

func TestGetAllRecipes_ok(t *testing.T) {
	resultAsBytes, _ := ioutil.ReadFile("../testresources/testGetAllRecipes_ok.json")
	expectedResult := string(resultAsBytes)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/recipes", nil)
	rec := httptest.NewRecorder()
	mockRepo := new(MockRepo)
	mockRepo.On("FindAll").Return(testresources.SampleRecipes(), nil)
	testee := Endpoint{mockRepo}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.GetAllRecipes(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, expectedResult, rec.Body.String())
	}
	mockRepo.AssertExpectations(t)
}

func TestGetAllRecipes_empty(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/recipes", nil)
	rec := httptest.NewRecorder()
	mockRepo := new(MockRepo)
	mockRepo.On("FindAll").Return([]*model.Recipe{}, nil)
	testee := Endpoint{mockRepo}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.GetAllRecipes(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, "[]", rec.Body.String())
	}
	mockRepo.AssertExpectations(t)
}

func TestGetAllRecipes_serverError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/recipes", nil)
	rec := httptest.NewRecorder()
	mockRepo := new(MockRepo)
	mockRepo.On("FindAll").Return(nil, errors.New("Database offline"))
	testee := Endpoint{mockRepo}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.GetAllRecipes(context)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
	mockRepo.AssertExpectations(t)
}
