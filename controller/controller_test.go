package controller

import (
	"api_key_test/oapi"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetHello(t *testing.T) {
	e := echo.New()
	a := &APIController{}

	data := []struct {
		name           string
		expectedStatus int
	}{
		{"正常系", http.StatusOK},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/hello", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if err := a.GetHello(c); err != nil {
				t.Errorf("得られたエラー: %v", err)
			}
			// ステータスコード確認
			gotStatus := rec.Code
			if d.expectedStatus != gotStatus {
				t.Errorf("case: %s, got: %d, want: %d", d.name, gotStatus, d.expectedStatus)
			}
		})
	}
}

func TestPostNumber(t *testing.T) {
	e := echo.New()
	a := &APIController{}

	data := []struct {
		name           string
		reqHeader      string
		reqBody        string
		expectedStatus int
	}{
		{
			"正常系",
			"testkey",
			`{"num": 1}`,
			http.StatusOK,
		},
		{
			"異常系_リクエストボディ不足",
			"testkey",
			`{}`,
			http.StatusBadRequest,
		},
		{
			"異常系_リクエストボディの型不一致",
			"testkey",
			`{"num": "1"}`,
			http.StatusBadRequest,
		},
		{
			"異常系_API_KEY不一致",
			"",
			`{"num": 1}`,
			http.StatusUnauthorized,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/number", strings.NewReader(d.reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			reqHeader := oapi.PostNumberParams{
				XAPIKEY: d.reqHeader,
			}

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if err := a.PostNumber(c, reqHeader); err != nil {
				t.Errorf("得られたエラー: %v", err)
			}

			// ステータスコード確認
			gotStatus := rec.Code
			if d.expectedStatus != gotStatus {
				t.Errorf("case: %s, got: %d, want: %d", d.name, gotStatus, d.expectedStatus)
			}
		})
	}
}

func TestDeleteNumber(t *testing.T) {
	e := echo.New()
	a := &APIController{}

	data := []struct {
		name           string
		reqHeader      string
		reqBody        string
		expectedStatus int
	}{
		{
			"正常系",
			"testkey",
			`{"num": 1}`,
			http.StatusOK,
		},
		{
			"異常系_リクエストボディ不足",
			"testkey",
			`{}`,
			http.StatusBadRequest,
		},
		{
			"異常系_リクエストボディの型不一致",
			"testkey",
			`{"num": "1"}`,
			http.StatusBadRequest,
		},
		{
			"異常系_API_KEY不一致",
			"",
			`{"num": 1}`,
			http.StatusUnauthorized,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/number", strings.NewReader(d.reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			reqHeader := oapi.DeleteNumberParams{
				XAPIKEY: d.reqHeader,
			}

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if err := a.DeleteNumber(c, reqHeader); err != nil {
				t.Errorf("得られたエラー: %v", err)
			}

			// ステータスコード確認
			gotStatus := rec.Code
			if d.expectedStatus != gotStatus {
				t.Errorf("case: %s, got: %d, want: %d", d.name, gotStatus, d.expectedStatus)
			}
		})
	}
}
