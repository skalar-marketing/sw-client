package sw

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//func TestPaymentClient_InitiatePayment(t *testing.T) {
//	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
//	assert.NoError(t, err)
//	result, err := client.Payment.InitiatePayment(InitiatePayment{})
//
//	assert.NoError(t, err)
//	assert.NotNil(t, result)
//}

func TestPaymentClient_ListShippingMethods(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Payment.ListShippingMethods(ListShippingMethodsFilter{
		Page:  1,
		Limit: 10,
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestPaymentClient_AvailablePayment(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Payment.AvailablePayment(ListAvailablePaymentFilter{
		Page:  1,
		Limit: 10,
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
