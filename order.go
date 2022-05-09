package sw

const (
	createOrderEndpoint              = "/checkout/order"
	cancelOrderEndpoint              = "/order/state/cancel"
	listOrdersEndpoint               = "/order"
	updatePaymentMethodOrderEndpoint = "/order/payment"
)

// OrderClient handles all requests regarding orders.
type OrderClient struct {
	client *Client
}

//Create adds new items into the cart.
func (client *OrderClient) Create(create CreateOrder) (*Order, error) {
	result := new(Order)

	if err := client.client.performPost(createOrderEndpoint, create, result); err != nil {
		return nil, err
	}

	return result, nil
}

//Cancel cancels items from the cart.
func (client *OrderClient) Cancel(cancel CancelOrder) (*CancelOrderResult, error) {
	result := new(CancelOrderResult)

	if err := client.client.performPost(cancelOrderEndpoint, cancel, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (client *OrderClient) List(filter ListOrderFilter) (*ListOrderResult, error) {
	result := new(ListOrderResult)

	if err := client.client.performPost(listOrdersEndpoint, filter, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (client *OrderClient) UpdatePaymentMethod(update UpdatePaymentMethodOrderBody) (*UpdatePaymentMethodOrderResult, error) {
	result := new(UpdatePaymentMethodOrderResult)

	if err := client.client.performPost(updatePaymentMethodOrderEndpoint, update, result); err != nil {
		return nil, err
	}

	return result, nil
}
