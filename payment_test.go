package sw

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//func TestPaymentClient_InitiatePayment(t *testing.T) {}

func TestPaymentClient_ListShippingMethods(t *testing.T) {
	data := new(ListShippingMethodsFilter)
	var req *http.Request
	ts, client := shopwareMock(t, new(ListShippingMethodsResult), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Payment.ListShippingMethods(ListShippingMethodsFilter{
		Page:  1,
		Limit: 10,
	})
	checkResponse(t, req, "/store-api/shipping-method")
	assert.Equal(t, 1, data.Page)
	assert.Equal(t, 10, data.Limit)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestPaymentClient_AvailablePayment(t *testing.T) {
	data := new(ListAvailablePaymentFilter)
	var req *http.Request
	ts, client := shopwareMock(t, new(ListAvailablePaymentResult), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Payment.AvailablePayment(ListAvailablePaymentFilter{
		Page:  1,
		Limit: 10,
	})
	checkResponse(t, req, "/store-api/payment-method")
	assert.Equal(t, 1, data.Page)
	assert.Equal(t, 10, data.Limit)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
