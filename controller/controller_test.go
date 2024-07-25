package controller

import (
	"api_key_test/oapi"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetHello(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	a := &APIController{}

	if assert.NoError(t, a.GetHello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestPostNumber(t *testing.T) {
	e := echo.New()

	val := `{"num": 1}`

	req := httptest.NewRequest(http.MethodPost, "/number", strings.NewReader(val))
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, rec)
	a := &APIController{}
	params := oapi.PostNumberParams{
		XAPIKEY: "testkey",
	}

	if assert.NoError(t, a.PostNumber(c, params)) {
		// ステータスコード確認
		assert.Equal(t, http.StatusOK, rec.Code, rec.Body.String())

		// レスポンスボディ確認
		var actualRes map[string]int
		expectedRes := map[string]int{"num": 2}
		err := json.Unmarshal(rec.Body.Bytes(), &actualRes)
		if assert.NoError(t, err) {
			assert.Equal(t, expectedRes, actualRes)
		}
	}
}

func TestDeleteNumber(t *testing.T) {
	e := echo.New()

	val := `{"num": 1}`

	req := httptest.NewRequest(http.MethodPost, "/number", strings.NewReader(val))
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, rec)
	a := &APIController{}
	params := oapi.DeleteNumberParams{
		XAPIKEY: "testkey",
	}

	if assert.NoError(t, a.DeleteNumber(c, params)) {
		// ステータスコード確認
		assert.Equal(t, http.StatusOK, rec.Code, rec.Body.String())

		// レスポンスボディ確認
		var actualRes map[string]int
		expectedRes := map[string]int{"num": 0}
		err := json.Unmarshal(rec.Body.Bytes(), &actualRes)
		if assert.NoError(t, err) {
			assert.Equal(t, expectedRes, actualRes)
		}
	}
}
