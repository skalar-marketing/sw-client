package sw

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCategory_List(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Category.List(ListCategoryFilter{
		Page:  1,
		Limit: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCategoryClient_Get(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	categories, err := client.Category.List(ListCategoryFilter{
		Page:  1,
		Limit: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, categories)
	result, err := client.Category.Get(categories.Elements[0].Id, GetCategoryFilter{
		Page:  1,
		Limit: 10,
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
