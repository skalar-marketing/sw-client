package sw

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOrderClient_Create(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Order.Create(CreateOrder{})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestOrderClient_List(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Order.List(ListOrderFilter{
		Page:  1,
		Limit: 10})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
