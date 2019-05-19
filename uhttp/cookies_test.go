package uhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRelURL(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/abc/xyz", nil)
	require.Nil(t, err)
	assert.Equal(t, "/abc/xyz", RelURL(req))
}

func TestUseCors(t *testing.T) {
	var res http.ResponseWriter = httptest.NewRecorder()
	UseCors(&res, &CorsHeaders{
		Origin:  "*",
		Headers: "*",
		Methods: "*",
	})
	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Methods"))
}

func TestUseUTF8Json(t *testing.T) {
	var res http.ResponseWriter = httptest.NewRecorder()

	UseUTF8Json(&res, "")
	exp := "application/json; charset=utf-8"
	assert.Equal(t, exp, res.Header().Get("Content-Type"))

	UseUTF8Json(&res, "abc")
	exp = "application/abc+json; charset=utf-8"
	assert.Equal(t, exp, res.Header().Get("Content-Type"))
}
