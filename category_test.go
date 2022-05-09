package sw

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCategory_List(t *testing.T) {
	data := new(ListCategoryFilter)
	var req *http.Request
	ts, client := shopwareMock(t, new(ListCategoryResult), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Category.List(ListCategoryFilter{
		Page:  1,
		Limit: 10,
	})
	checkResponse(t, req, "/store-api/category")
	assert.Equal(t, 1, data.Page)
	assert.Equal(t, 10, data.Limit)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCategoryClient_Get(t *testing.T) {
	data := new(GetCategoryFilter)
	var req *http.Request
	ts, client := shopwareMock(t, new(Category), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Category.Get("id", GetCategoryFilter{
		Page:  1,
		Limit: 10,
	})
	checkResponse(t, req, "/store-api/category/id")
	assert.Equal(t, 1, data.Page)
	assert.Equal(t, 10, data.Limit)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
