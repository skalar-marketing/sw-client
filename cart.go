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
