package rest

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLogin_ok(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", buildFormData("testUser", "1234"))
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
	req := httptest.NewRequest(http.MethodPost, "/login", buildFormData("invalid", "1234"))
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

func buildFormData(username, password string) *strings.Reader {
	f := make(url.Values)
	f.Set("username", username)
	f.Set("password", password)
	return strings.NewReader(f.Encode())
}
