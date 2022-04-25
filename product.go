package sw

import "fmt"

const (
	listProductsCategoryEndpoint = "/product-listing/%s"
)

// ProductClient handles all requests regarding products.
type ProductClient struct {
	client *Client
}

// ListCategory returns the products for a category.
func (client *ProductClient) ListCategory(categoryId string, filter ListProductsCategoryFilter) (*ListProductsCategoryResult, error) {
	result := new(ListProductsCategoryResult)

	if err := client.client.performPost(fmt.Sprintf(listProductsCategoryEndpoint, categoryId), filter, &result); err != nil {
		return nil, err
	}

	return result, nil
}
