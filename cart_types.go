package sw

/* Die API Doku für cart hat keine required. Die hab ich mir selbst ausgesucht. */

// Cart is a cart.
type Cart struct { /*EVtl. CartResult? */
	ApiAlias string `json:"apiAlias"`
	Name     string `json:"name"`
	Token    string `json:"token"`
	//Price     []Price      `json:"price"`		/* Unsichtbar weil ansonsten der Test nicht funktioniert. */
	LineItems []LineItems  `json:"lineItems"`
	Errors    []CartErrors `json:"errors"` /* Struct hat anderen Namen, weil es Errors auch bei customer_types gibt. */
	//Transactions []Transactions `json:"transactions"`
	//Modified        bool           `json:"modified"`
	//CustomerComment string         `json:"customerComment"`
	//AffiliateCode   string         `json:"affiliateCode"`
	//CampaignCode    string         `json:"campaignCode"`
}

////DeleteCartResult is the successfully deleted cart.
//type DeleteCartResult struct {
//	Success bool `json:"success"`
//}

//AddNUpdateItemCart adds or updates items to the cart.
type AddNUpdateItemCart struct {
	//ApiAlias string  `json:"apiAlias"`
	Items []Items `json:"items"`
}

//Price is the struct Price for Cart.
type Price struct {
	NetPrice      float64 `json:"netPrice"`
	TotalPrice    float64 `json:"totalPrice"`
	PositionPrice float64 `json:"positionPrice"`
	TaxStatus     string  `json:"taxStatus"`
}

//LineItems lists all items in the cart. It's the struct for LineItems in Cart.
type LineItems struct {
	//Id           string `json:"id"`
	//ReferencedId string `json:"referencedId"`
	//Label        string `json:"label"`
	Quantity int32 `json:"quantity"`
	//Type         string `json:"type"`
	//Good         bool   `json:"good"`
	//Description  string `json:"description"`
	//Removable    bool   `json:"removable"`
	//Stackable    bool   `json:"stackable"`
	//Modified     bool   `json:"modified"`
}

//CartErrors is a list of all cart errors, such as insufficient stocks, invalid addresses or vouchers. It's the struct for CartErrors in Cart.
type CartErrors struct {
	//Key     string `json:"key"`
	//Level   string `json:"level"`
	//Message string `json:"message"`
}

//Transactions is a list of all payment transactions associated with the current cart.
//type Transactions struct {
//	PaymentMethodId string `json:"paymentMethodId"`
//}

//Items is the struct for Items in AddUpdateItemCart.
type Items struct {
	//Id          string `json:"id"`
	//ReferenceId string `json:"referenceId"`
	//Label       string `json:"label"`
	Quantity int32 `json:"quantity"`
	//Type         string `json:"type"`
	//Good         bool   `json:"good"`
	//Description  string `json:"description"`
	//Removable    bool   `json:"removable"`
	//Stackable    bool   `json:"stackable"`
	//Modified     bool   `json:"modified"`
}
