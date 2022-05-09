package sw

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCartClient_CreateOrFetch(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Cart.CreateOrFetch()

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

/* Delete */

func TestCartClient_AddItems(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Cart.AddItems(AddItemCart{})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

/* RemoveItems */

func TestCartClient_UpdateItems(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Cart.UpdateItems(UpdateItemCart{})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
