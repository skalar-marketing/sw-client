package sw

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCartClient_CreateOrFetch(t *testing.T) {
	var req *http.Request
	ts, client := shopwareMock(t, new(Cart), nil, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Cart.CreateOrFetch()
	checkResponse(t, req, "/store-api/checkout/cart")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCartClient_Delete(t *testing.T) {
	var req *http.Request
	ts, client := shopwareMock(t, new(DeleteCartResult), nil, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Cart.Delete()
	checkResponse(t, req, "/store-api/checkout/cart")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCartClient_AddItems(t *testing.T) {
	data := new(AddItemCart)
	var req *http.Request
	ts, client := shopwareMock(t, new(Cart), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Cart.AddItems(AddItemCart{})
	checkResponse(t, req, "/store-api/checkout/cart/line-item")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCartClient_RemoveItems(t *testing.T) {
	var req *http.Request
	ts, client := shopwareMock(t, new(Cart), nil, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Cart.RemoveItems()
	checkResponse(t, req, "/store-api/checkout/cart/line-item")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCartClient_UpdateItems(t *testing.T) {
	data := new(UpdateItemCart)
	var req *http.Request
	ts, client := shopwareMock(t, new(Cart), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Cart.UpdateItems(UpdateItemCart{})
	checkResponse(t, req, "/store-api/checkout/cart/line-item")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
