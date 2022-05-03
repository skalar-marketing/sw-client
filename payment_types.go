package sw

import "time"

// todo Initiate a payment for an order

type InitiateBody struct {
	OrderId string `json:"orderId"`
	//FinishUrl string	`json:"finishUrl"`
	//ErrorUrl string	`json:"errorUrl"`
}

/*todo in fetch shipping methods is onlyAvailable a var?*/

// ListShippingMethodsFilter is used to filter available shipping methods.
type ListShippingMethodsFilter struct {
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
	Filter     []Filter `json:"filter"`
	Sort       []Sort   `json:"sort"`
	PostFilter []Filter `json:"post-filter"`
	//associations
	Aggregations []Aggregation `json:"aggregations"`
	Grouping     []string      `json:"grouping"`
}

// ListAvailablePaymentFilter is used to filter available payments methods.
type ListAvailablePaymentFilter struct {
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
	Filter     []Filter `json:"filter"`
	Sort       []Sort   `json:"sort"`
	PostFilter []Filter `json:"post-filter"`
	//associations
	Aggregations []Aggregation `json:"aggregations"`
	Grouping     []string      `json:"grouping"`
}

// ListShippingMethodsResult contains a list of available payments methods.
type ListShippingMethodsResult struct {
	Total int `json:"total"`
	//Aggregations []Aggregations `json:"aggregations"`
	Elements []ElementsShipping `json:"elements"`
}

// ListAvailablePaymentResult contains a list of available payments methods.
type ListAvailablePaymentResult struct {
	Total int `json:"total"`
	//Aggregations []Aggregations `json:"aggregations"`
	Elements []ElementsPayment `json:"elements"`
}

//ElementsShipping is the struct for Elements in ListShippingMethodsResult.
type ElementsShipping struct {
	//Id             string         `json:"id"`
	Name string `json:"name"`
	//Active         bool           `json:"active"`
	//CustomFields   []CustomFields `json:"customFields"`
	//MediaId        string         `json:"mediaId"`
	DeliveryTimeId string `json:"deliveryTimeId"`
	TaxType        string `json:"taxType"`
	//Description    string         `json:"description"`
	//TrackingUrl string	`json:"trackingUrl"`
	CreatedAt time.Time `json:"createdAt"`
	//UpdatedAt           time.Time      `json:"updatedAt"`
	//Translated          []Translated   `json:"translated"`
	//DeliveryTime []DeliveryTime `json:"deliveryTime"`
	//AvailabilityRule []AvailabilityRule `json:"availabilityRule"`
	//Prices []Prices	`json:"prices"`
	//Media []Media	`json:"media"`
	//Tags []Tags `json:"tags"`
	//Tax []Tax `json:"tax"`
}

//ElementsPayment is the struct for Elements in ListAvailablePaymentResult.
type ElementsPayment struct {
	//Id                  string         `json:"id"`
	Name string `json:"name"`
	//DistinguishableName string         `json:"distinguishableName"`
	//Description         string         `json:"description"`
	//Position            int64          `json:"position"`
	//Active              bool           `json:"active"`
	//AfterOrderEnabled   bool           `json:"afterOrderEnabled"`
	//CustomFields        []CustomFields `json:"customFields"`
	//MediaId             string         `json:"mediaId"`
	CreatedAt time.Time `json:"createdAt"`
	//UpdatedAt           time.Time      `json:"updatedAt"`
	//Translated          []Translated   `json:"translated"`
	//Media []Media	`json:"media"`
}
