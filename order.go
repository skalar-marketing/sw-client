package sw

const (
	createOrderEndpoint                = "/checkout/order"
	cancelOrderEndpoint                = "/order/state/cancel"
	listOrdersEndpoint                 = "/order"
	updatePaymentMethodOfOrderEndpoint = "/order/payment"
)

// OrderClient handles all requests regarding orders.
type OrderClient struct {
	client *Client
}
