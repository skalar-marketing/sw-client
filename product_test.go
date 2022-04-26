package sw

import (
	"github.com/stretchr/testify/assert"
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
	result, err := client.Product.ListCategory(categories.Elements[0].Id, ListProductsFilter{
		Page:  1,
		Limit: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestProductClient_Get(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	categories, err := client.Category.List(ListCategoryFilter{
		Page:  1,
		Limit: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, categories)
	products, err := client.Product.ListCategory(categories.Elements[0].Id, ListProductsFilter{
		Page:  1,
		Limit: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, products)
	product, err := client.Product.Get(products.Elements[0].Id)
	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestProductClient_Search(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Product.Search(SearchProductsFilter{
		Page:   1,
		Limit:  10,
		Search: "test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
