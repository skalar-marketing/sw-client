package sw

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func shopwareMock(t *testing.T, result, data any, callback func(*http.Request)) (*httptest.Server, *Client) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if data != nil {
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(body, data))
		}
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

func TestClient_Context(t *testing.T) {
	var req *http.Request
	ts, client := shopwareMock(t, new(Context), nil, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Context()

	checkResponse(t, req, "/store-api/context")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
