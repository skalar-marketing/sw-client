package sw

import (
	"os"
	"testing"
)

func TestProductClient_ListCategory(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	categories, err := client.Category.List(ListCategoryFilter{
		Page:  1,
		Limit: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, categories)
	result, err := client.Product.ListCategory(categories.Elements[0].Id, ListProductsCategoryFilter{})
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
