package sw

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestOrderClient_Create(t *testing.T) {
	data := new(CreateOrder)
	var req *http.Request
	ts, client := shopwareMock(t, new(Order), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Order.Create(CreateOrder{})
	checkResponse(t, req, "/store-api/checkout/order")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestOrderClient_Cancel(t *testing.T) {
	data := new(CancelOrder)
	var req *http.Request
	ts, client := shopwareMock(t, new(CancelOrderResult), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Order.Cancel(CancelOrder{})
	checkResponse(t, req, "/store-api/order/state/cancel")

	assert.NoError(t, err)
	assert.NotNil(t, result)
} // todo braucht noch den Test für den ContextToken

func TestOrderClient_List(t *testing.T) {
	data := new(ListOrderFilter)
	var req *http.Request
	ts, client := shopwareMock(t, new(ListOrderResult), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Order.List(ListOrderFilter{
		Page:  1,
		Limit: 10,
	})
	checkResponse(t, req, "/store-api/order")
	assert.Equal(t, 1, data.Page)
	assert.Equal(t, 10, data.Limit)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestOrderClient_UpdatePaymentMethod(t *testing.T) {
	data := new(UpdatePaymentMethodOrderBody)
	var req *http.Request
	ts, client := shopwareMock(t, new(UpdatePaymentMethodOrderResult), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Order.UpdatePaymentMethod(UpdatePaymentMethodOrderBody{})
	checkResponse(t, req, "/store-api/order/payment")

	assert.NoError(t, err)
	assert.NotNil(t, result)
} // todo braucht noch den Test für den ContextToken
