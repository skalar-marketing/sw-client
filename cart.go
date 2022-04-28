package sw

const (
	createOrFetchCartEndpoint = "/checkout/cart"
	deleteCartEndpoint        = "/checkout/cart"
	addItemCartEndpoint       = "/checkout/cart/line-item"
	updateItemCartEndpoint    = "/checkout/cart/line-item"
	removeItemCartEndpoint    = "/checkout/cart/line-item"
)

// CartClient handles all requests regarding the cart.
type CartClient struct {
	client *Client
}

//AddItemsCart adds new items into the cart.
func (client *CartClient) AddItemsCart(add AddNUpdateItemCart) (*Cart, error) {
	result := new(Cart)

	if err := client.client.performPost(addItemCartEndpoint, add, result); err != nil {
		return nil, err
	}

	return result, nil
}

//UpdateItemsCart updates items inside the cart.
func (client *CartClient) UpdateItemsCart(update AddNUpdateItemCart) (*Cart, error) {
	result := new(Cart)

	if err := client.client.performPost(updateItemCartEndpoint, update, result); err != nil {
		return nil, err
	}

	return result, nil
}
