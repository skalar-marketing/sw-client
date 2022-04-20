package sw

import "fmt"

const (
	listCategoriesEndpoint = "/category"
	getCategoryEndpoint    = "/category/%s"
)

// CategoryClient handles all requests regarding categories.
type CategoryClient struct {
	client *Client
}

// List performs a filtered search for categories.
func (client *CategoryClient) List(filter ListCategoryFilter) (*ListCategoryResult, error) {
	result := new(ListCategoryResult)

	if err := client.client.performPost(listCategoriesEndpoint, filter, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Get returns a single category.
func (client *CategoryClient) Get(id string, filter GetCategoryFilter) (*Category, error) {
	result := new(Category)

	if err := client.client.performPost(fmt.Sprintf(getCategoryEndpoint, id), filter, &result); err != nil {
		return nil, err
	}

	return result, nil
}
