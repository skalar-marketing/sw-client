package sw

import (
	"time"
)

// Order is an order.
type Order struct {
	//Id                      string       `json:"id"`
	//VersionId               string       `json:"versionId"`
	//OrderNumber             string       `json:"orderNumber"`
	BillingAddressId string `json:"billingAddressId"`
	//BillingAddressVersionId string       `json:"billingAddressVersionId"`
	CurrencyId     string    `json:"currencyId"`
	LangaugeId     string    `json:"langaugeId"`
	SalesChannelId string    `json:"salesChannelId"`
	OrderDateTime  time.Time `json:"orderDateTime"`
	//OrderDate               string       `json:"orderDate"`
	//PriceOrder              []PriceOrder `json:"price"`
	//AmountTotal   float64         `json:"amountTotal"`
	//AmountNet     float64         `json:"amountNet"`
	//PositionPrice float64         `json:"positionPrice"`
	//TaxStatus     string          `json:"taxStatus"`
	//ShippingCosts []ShippingCosts `json:"shippingCosts"`
	//ShippingTotal float64	`json:"shippingTotal"`
	CurrencyFactor float64 `json:"currencyFactor"`
	//DeepLinkCode   string  `json:"deepLinkCode"`
	//AffiliateCode string	`json:"affiliateCode"`
	//CampaignCode	string	`json:"campaignCode"`
	//CustomerComment string	`json:"customerComment"`
	//CustomFields	[]CustomFields `json:"customFields"`
	//CreatedById	string	`json:"createdById"`
	//UpdateById	string	`json:"updateById"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt time.Time	`json:"updatedAt"`
	//StateMachineState []StateMachineState `json:"stateMachineState"`
	//OrderCustomer []OrderCustomer `json:"shippingOrderAddress"`
	//Currency      []Currency      `json:"currency"`
	//Language      []Language      `json:"language"`
	//Addresses []Addresses	`json:"addresses"`
	//BillingAddress []BillingAddress	`json:"billingAddress"`
	//Deliveries []Deliveries	`json:"deliveries"`
	//LineItems	[]LineItems	`json:"lineItems"`
	//Transactions []Transactions	`json:"transactions"`
	//Documents []Documents `json:"documents"`
	//Tags []Tags	`json:"tags"`
}

/*Nichts ist "required"*/
/* TODO Vernünftigen Namen aussuchen. */

type CreateOrderBody struct {
	CustomerComment string `json:"customerComment"`
	AffiliateCode   string `json:"affiliateCode"`
	CampaignCode    string `json:"campaignCode"`
}

/* TODO Vernünftigen Namen aussuchen. */

type CancelOrderBody struct {
	OrderId string `json:"orderId"`
}

/* TODO Passt der Name? Muss hier auch No Aggregations und OnlyAggregations. Required überprüfen. */

type ListOrderFilter struct {
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
	Filter     []Filter `json:"filter"` /* Ab hier sind die structs in category_types. */
	Sort       []Sort   `json:"sort"`
	PostFilter []Filter `json:"post-filter"`
	//associations
	Aggregations   []Aggregation `json:"aggregations"`
	Grouping       []string      `json:"grouping"`
	CheckPromotion bool          `json:"checkPromotion"`
}

/* TODO Vernünftigen Namen aussuchen. */

type UpdatePaymentMethodOfOrderBody struct {
	PaymentMethodOfId string `json:"paymentMethodOfId"`
	OrderId           string `json:"orderId"`
}

//CancelOrderResult is the result of a cancelled order.
//Returns the state of the state machine. The order state will be set to 'cancelled'.
type CancelOrderResult struct {
	//Id            string `json:"id"`
	TechnicalName string `json:"technicalName"`
	Name          string `json:"name"`
	//CustomFields    []CustomFields `json:"customFields"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt       time.Time      `json:"updatedAt"`
	//Translated      []Translated   `json:"translated"`
}

//ListOrderResult is an array of orders and an indicator if the payment of the order can be changed.
type ListOrderResult struct {
	//Orders            []Orders            `json:"orders"`
	//PaymentChangeable []PaymentChangeable `json:"paymentChangeable"`
}

//UpdatePaymentMethodOfOrderResult is true if the payment method of the order is successful.
type UpdatePaymentMethodOfOrderResult struct {
	//Success bool	`json:"success"`
}

//Language is the struct for Language in Order.
type Language struct {
	//Id                string `json:"id"`
	//ParentId          string `json:"parentId"`
	LocaleId string `json:"localeId"`
	//TranslationCodeId string `json:"translationCodeId"`
	Name string `json:"name"`
	//CustomFields      []CustomFields `json:"customFields"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt         time.Time         `json:"updatedAt"`
	//Parent            []Parent          `json:"parent"`
	//Locale            []Locale          `json:"locale"`
	//TranslationCode   []TranslationCode `json:"translationCode"`
	//Children          []Children        `json:"children"`
}

//Currency is the struct for Currency in Order.
type Currency struct {
	//Id              string         `json:"id"`
	Factor    float64 `json:"factor"`
	Symbol    string  `json:"symbol"`
	IsoCode   string  `json:"isoCode"`
	ShortName string  `json:"shortName"`
	Name      string  `json:"name"`
	//Position        int64          `json:"position"`
	//IsSystemDefault bool           `json:"isSystemDefault"`
	//CustomFields    []CustomFields `json:"customFields"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt       time.Time      `json:"updatedAt"`
	//Translated      []Translated   `json:"translated"`
}

//OrderCustomer is the struct for OrderCustomer in Order.
type OrderCustomer struct {
	//Id             string         `json:"id"`
	//VersionId      string         `json:"versionId"`
	Email        string `json:"email"`
	SalutationId string `json:"salutationId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	//Company        string         `json:"company"`
	//Title          string         `json:"title"`
	//VatIds         []string       `json:"vatIds"`
	//CustomerNumber string         `json:"customerNumber"`
	//CustomFields   []CustomFields `json:"customFields"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt      time.Time      `json:"updatedAt"`
	//Salutation     []Salutation   `json:"salutation"`
}

//StateMachineState is the struct for StateMachineState in Order.
type StateMachineState struct {
	//Id            string         `json:"id"`
	TechnicalName string `json:"technicalName"`
	Name          string `json:"name"`
	//CustomFields  []CustomFields `json:"customFields"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt     time.Time      `json:"updatedAt"`
	//Translated    []Translated   `json:"translated"`
}

//ShippingCosts is the struct for ShippingCosts in Order.
type ShippingCosts struct {
	UnitPrice  float64 `json:"unitPrice"`
	TotalPrice float64 `json:"totalPrice"`
	Quantity   int64   `json:"quantity"`
	//CalculatedTaxes []CalculatedTaxes `json:"calculatedTaxes"`
	//TaxRules        []TaxRules        `json:"taxRules"`
	//ReferencePrice  []ReferencePrice  `json:"referencePrice"`
	//ListPrice       []ListPrice       `json:"listPrice"`
}

//PriceOrder is the struct PriceOrder in Order.
type PriceOrder struct {
	NetPrice   float64 `json:"netPrice"`
	TotalPrice float64 `json:"totalPrice"`
	//CalculatedTaxes []CalculatedTaxes `json:"calculatedTaxes"`
	//TaxRules        []TaxRules        `json:"taxRules"`
	PositionPrice float64 `json:"positionPrice"`
	RawTotal      float64 `json:"rawTotal"`
	TaxStatus     string  `json:"taxStatus"`
}

type Children struct {
	//Id                string         `json:"id"`
	//ParentId          string         `json:"parentId"`
	LocaleId string `json:"localeId"`
	//TranslationCodeId string         `json:"translationCodeId"`
	Name string `json:"name"`
	//CustomFields      []CustomFields `json:"customFields"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt         time.Time      `json:"updatedAt"`
	//Parent            []Parent       `json:"parent"`
	//Locale            []Locale          `json:"locale"`
	//TranslationCode   []TranslationCode `json:"translationCode"`
	//Children          []Children        `json:"children"`
}

//TranslationCode is the struct TranslationCode in Language.
type TranslationCode struct {
	//Id        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Territory string `json:"territory"`
	//CustomFields  []CustomFileds `json:"customFields"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt     time.Time      `json:"updatedAt"`
	//Translated    []Translated   `json:"translated"`
}

//Locale is the struct Locale in Language.
type Locale struct {
	//Id        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Territory string `json:"territory"`
	//CustomFields  []CustomFileds `json:"customFields"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt     time.Time      `json:"updatedAt"`
	//Translated    []Translated   `json:"translated"`
}

//Parent is the struct Parent in Language.
type Parent struct {
	//Id                string         `json:"id"`
	//ParentId          string         `json:"parentId"`
	LocaleId string `json:"localeId"`
	//TranslationCodeId string         `json:"translationCodeId"`
	Name string `json:"name"`
	//CustomFields      []CustomFields `json:"customFields"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt         time.Time      `json:"updatedAt"`
	//Parent            []Parent       `json:"parent"`
	//Locale            []Locale          `json:"locale"`
	//TranslationCode   []TranslationCode `json:"translationCode"`
	//Children          []Children        `json:"children"`
}

//Salutation is the struct of Salutation in OrderCustomer.
type Salutation struct {
	//id            string         `json:"id"`
	SalutationKey string `json:"salutationKey"`
	DisplayName   string `json:"displayName"`
	LetterName    string `json:"letterName"`
	//CustomFields  []CustomFileds `json:"customFields"`
	CreatedAt time.Time `json:"createdAt"`
	//updatedAt     time.Time      `json:"updatedAt"`
	//Translated    []Translated   `json:"translated"`
}

//ListPrice is the struct ListPrice in ShippingCosts.
//type ListPrice struct {
//	Price      float64 `json:"price"`
//	Discount   float64 `json:"discount"`
//	Percentage float64 `json:"percentage"`
//}
