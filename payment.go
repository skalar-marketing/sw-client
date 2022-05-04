package sw

const (
	initiatePaymentEndpoint         = "/handle-payment"
	shippingMethodsEndpoint         = "/shipping-method"
	availablePaymentMethodsEndpoint = "/payment-method"
)

// PaymentClient handles all requests regarding a payment flow after an order has been created.
type PaymentClient struct {
	client *Client
}

// InitiatePayment is to initiate a payment flow after an order has been created.
func (client *ProductClient) InitiatePayment(initiate InitiatePayment, filter ListProductsFilter) (*_, error) {
	result := new(_)

	if err := client.client.performPost(updateItemCartEndpoint, initiate, result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListShippingMethods performs a filtered search for shipping methods.
func (client *ProductClient) ListShippingMethods(filter ListShippingMethodsFilter) (*ListShippingMethodsResult, error) {
	result := new(ListShippingMethodsResult)

	if err := client.client.performPost(shippingMethodsEndpoint, filter, result); err != nil {
		return nil, err
	}

	return result, nil
}

// AvailablePayment performs a filtered search for shipping methods.
func (client *ProductClient) AvailablePayment(filter ListAvailablePaymentFilter) (*ListAvailablePaymentResult, error) {
	result := new(ListAvailablePaymentResult)

	if err := client.client.performPost(availablePaymentMethodsEndpoint, filter, result); err != nil {
		return nil, err
	}

	return result, nil
}
