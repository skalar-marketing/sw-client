package sw

// Context contains some general information about the store and the user.
type Context struct {
	Token string `json:"token"`
	//currentCustomerGroup TODO
	//fallbackCustomerGroup
	//currency
	//salesChannel
	//taxRules
	//customer
	//paymentMethod
	//shippingMethod
	//shippingLocation
	//rulesIds
	RulesLocked bool `json:"rulesLocked"`
	//permissions
	PermissionsLocked bool   `json:"permissionsLocked"`
	ApiAlias          string `json:"apiAlias"`
}
