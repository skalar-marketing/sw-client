package sw

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestProductClient_ListCategory(t *testing.T) {
	data := new(ListProductsFilter)
	var req *http.Request
	ts, client := shopwareMock(t, new(ListProductsResult), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Product.ListCategory("id", ListProductsFilter{
		Page:  1,
		Limit: 10,
	})
	checkResponse(t, req, "/store-api/product-listing/id")
	assert.Equal(t, 1, data.Page)
	assert.Equal(t, 10, data.Limit)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestProductClient_Get(t *testing.T) {
	var req *http.Request
	ts, client := shopwareMock(t, new(GetProductResult), nil, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Product.Get("id")

	checkResponse(t, req, "/store-api/product/id")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestProductClient_Search(t *testing.T) {
	data := new(SearchProductsFilter)
	var req *http.Request
	ts, client := shopwareMock(t, new(ListProductsResult), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Product.Search(SearchProductsFilter{
		Page:   1,
		Limit:  10,
		Search: "test",
	})
	checkResponse(t, req, "/store-api/search")
	assert.Equal(t, 1, data.Page)
	assert.Equal(t, 10, data.Limit)
	assert.Equal(t, "test", data.Search)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
