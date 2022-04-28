package sw

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCartClient_AddItemsCart(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Cart.AddItemsCart(AddNUpdateItemCart{})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCartClient_UpdateItemsCart(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Cart.UpdateItemsCart(AddNUpdateItemCart{})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
