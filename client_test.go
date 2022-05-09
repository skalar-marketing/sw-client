package sw

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestClient_Context(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Context()
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func shopwareMock(t *testing.T, result, data any, callback func(*http.Request)) (*httptest.Server, *Client) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(body, data))
		resp, err := json.Marshal(result)
		assert.NoError(t, err)
		_, err = w.Write(resp)
		assert.NoError(t, err)
		callback(r)
	}))
	client, err := NewClient(ts.URL, "key")
	assert.NoError(t, err)
	return ts, client
}

func checkResponse(t *testing.T, req *http.Request, path string) {
	assert.Equal(t, path, req.URL.Path)
	assert.Equal(t, "key", req.Header.Get("sw-access-key"))
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
	assert.Equal(t, "application/json", req.Header.Get("Accept"))
}
