package rest

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin_ok(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("username=testUser&password=1234"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	testee := LoginEndpoint{"abc"}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.Login(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Len(t, rec.Body.String(), 147)
	}
}

func TestLogin_invalidCredentials(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("username=invalid&password=1234"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	testee := LoginEndpoint{"abc"}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.Login(context)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}

func TestLogin_emptyBody(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	rec := httptest.NewRecorder()
	testee := LoginEndpoint{"abc"}

	context := e.NewContext(req, rec)

	if assert.NoError(t, testee.Login(context)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}
