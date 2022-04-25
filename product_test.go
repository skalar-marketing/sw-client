package sw

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestProductClient_ListCategory(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	products, err := client.Product.ListCategory("", ListProductsCategoryFilter{
		Page:  1,
		Limit: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, products)
	result, err := client.Product.ListCategory(products.Elements[0].Id, ListProductsCategoryFilter{})
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
