package sw

const (
	createOrFetchCartEndpoint = "/checkout/cart"
	deleteCartEndpoint        = "/checkout/cart"
	addItemCartEndpoint       = "/checkout/cart/line-item"
	removeItemCartEndpoint    = "/checkout/cart/line-item"
	updateItemCartEndpoint    = "/checkout/cart/line-item"
)

// CartClient handles all requests regarding the cart.
type CartClient struct {
	client *Client
}

// CreateOrFetch creates or fetches a new cart.
func (client *CartClient) CreateOrFetch() (*Cart, error) {
	result := new(Cart)

	if err := client.client.performGet(createOrFetchCartEndpoint, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Delete deletes a cart.
func (client *CartClient) Delete() (*DeleteCartResult, error) {
	result := new(DeleteCartResult)

	if err := client.client.performDelete(deleteCartEndpoint, result); err != nil {
		return nil, err
	}

	return result, nil
}

// AddItems adds new items into the cart.
func (client *CartClient) AddItems(add AddItemCart) (*Cart, error) {
	result := new(Cart)

	if err := client.client.performPost(addItemCartEndpoint, add, result); err != nil {
		return nil, err
	}

	return result, nil
}

// RemoveItems removes items from the cart and recalculates it.
func (client *CartClient) RemoveItems() (*Cart, error) {
	result := new(Cart)

	if err := client.client.performDelete(removeItemCartEndpoint, result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateItems updates items inside the cart.
func (client *CartClient) UpdateItems(update UpdateItemCart) (*Cart, error) {
	result := new(Cart)

	if err := client.client.performPost(updateItemCartEndpoint, update, result); err != nil {
		return nil, err
	}

	return result, nil
}
