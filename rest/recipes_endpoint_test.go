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
	"strings"
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

func (repo *MockRepo) FindOne(rid string) (*model.Recipe, error) {
	args := repo.Called()
	recipe, ok := args.Get(0).(*model.Recipe)
	e := args.Error(1)

	if ok {
		return recipe, e
	} else {
		return nil, e
	}
}

func (repo *MockRepo) Insert(toBeInserted model.Recipe) (string, error) {
	args := repo.Called()
	rid := args.String(0)
	e := args.Error(1)

	if e != nil {
		return "", e
	} else {
		return rid, e
	}
}

func TestGetSingleRecipe_ok(t *testing.T) {
	resultAsBytes, _ := ioutil.ReadFile("../testresources/testGetSingleRecipe_ok.json")
	expectedResult := string(resultAsBytes)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/recipes/123456789", nil)
	rec := httptest.NewRecorder()
	mockRepo := new(MockRepo)
	mockRepo.On("FindOne").Return(testresources.SampleRecipes()[0], nil)
	testee := Endpoint{mockRepo}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.GetSingleRecipe(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, expectedResult, rec.Body.String())
	}
	mockRepo.AssertExpectations(t)
}

func TestGetSingleRecipe_notFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/recipes/invalidId", nil)
	rec := httptest.NewRecorder()
	mockRepo := new(MockRepo)
	mockRepo.On("FindOne").Return(nil, nil)
	testee := Endpoint{mockRepo}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.GetSingleRecipe(context)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
	mockRepo.AssertExpectations(t)
}

func TestGetSingleRecipe_serverError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/recipes/123456789", nil)
	rec := httptest.NewRecorder()
	mockRepo := new(MockRepo)
	mockRepo.On("FindOne").Return(nil, errors.New("Database offline"))
	testee := Endpoint{mockRepo}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.GetSingleRecipe(context)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
	mockRepo.AssertExpectations(t)
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

func TestPostSingleRecipe_ok(t *testing.T) {
	resultAsBytes, _ := ioutil.ReadFile("../testresources/newRecipe.json")
	expectedResult := string(resultAsBytes)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/recipes", strings.NewReader(expectedResult))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockRepo := new(MockRepo)
	mockRepo.On("Insert").Return("1a2b3c4e5f6g7h", nil)
	testee := Endpoint{mockRepo}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.PostNewRecipe(context)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "/recipes/1a2b3c4e5f6g7h", rec.Header().Get("Location"))
	}
	mockRepo.AssertExpectations(t)
}

func TestPostSingleRecipe_BadRequest(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/recipes", strings.NewReader("{elem:"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockRepo := new(MockRepo)
	testee := Endpoint{mockRepo}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.PostNewRecipe(context)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
	mockRepo.AssertExpectations(t)
}

func TestPostSingleRecipe_serverError(t *testing.T) {
	resultAsBytes, _ := ioutil.ReadFile("../testresources/newRecipe.json")
	expectedResult := string(resultAsBytes)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/recipes", strings.NewReader(expectedResult))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	mockRepo := new(MockRepo)
	mockRepo.On("Insert").Return("", errors.New("Something's wrong with Echo"))
	testee := Endpoint{mockRepo}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.PostNewRecipe(context)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
	mockRepo.AssertExpectations(t)
}
