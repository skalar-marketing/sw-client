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
