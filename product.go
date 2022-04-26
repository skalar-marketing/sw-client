package sw

import "fmt"

const (
	listProductsCategoryEndpoint = "/product-listing/%s"
	getProductEndpoint           = "/product/%s"
	searchProductsEndpoint       = "/search"
)

// ProductClient handles all requests regarding products.
type ProductClient struct {
	client *Client
}

// ListCategory returns the products for a category.
func (client *ProductClient) ListCategory(categoryId string, filter ListProductsFilter) (*ListProductsResult, error) {
	result := new(ListProductsResult)

	if err := client.client.performPost(fmt.Sprintf(listProductsCategoryEndpoint, categoryId), filter, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Get returns a single product.
func (client *ProductClient) Get(id string) (*GetProductResult, error) {
	result := new(GetProductResult)

	if err := client.client.performPost(fmt.Sprintf(getProductEndpoint, id), nil, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Search searches for products.
func (client *ProductClient) Search(filter ListProductsFilter) (*ListProductsResult, error) {
	result := new(ListProductsResult)

	if err := client.client.performPost(searchProductsEndpoint, filter, result); err != nil {
		return nil, err
	}

	return result, nil
}
