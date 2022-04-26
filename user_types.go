package sw

import "time"

// Login logs in customers given their credentials.
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ConfirmRegistration confirms a customer registration.
type ConfirmRegistration struct {
	Hash string `json:"hash"`
	Em   string `json:"em"`
}

// Register registers a customer. Used both for normal customers and guest customers.
type Register struct {
	Email                  string           `json:"email"`
	Password               string           `json:"password"`
	SalutationId           string           `json:"salutationId"`
	FirstName              string           `json:"firstName"`
	LastName               string           `json:"lastName"`
	AcceptedDataProtection string           `json:"acceptedDataProtection"`
	StorefrontUrl          string           `json:"storefrontUrl"`
	BillingAddress         []BillingAddress `json:"billingAddress"` //Das Objekt ist unten
	//ShippingAddress        []ShippingAddress `json:"shippingAddress"` //Das Objekt ist unten
	//AccountType   string `json:"accountType"`
	//Guest         string `json:"guest"`
	//BirthdayDay   string `json:"birthdayDay"`
	//BirthdayMonth string `json:"birthdayMonth"`
	//BirthdayYear  string `json:"birthdayYear"`
	//Title         string `json:"title"`
	//AffiliateCode string	`json:"affiliateCode"`
	//CampaignCode	string	`json:"campaignCode"`
}

// LoginResult returns a context token which is associated with the logged in user.
type LoginResult struct {
	ContextToken string `json:"contextToken"`
} //TODO Response 401

// SuccessfulLogout returns a context token for the anonymous user.
type SuccessfulLogout struct {
	ContextToken string `json:"contextToken"`
} //TODO Response 403

//TODO Response 200, 404, 412

// RegisterResult is the result of a successful registration.
type RegisterResult struct {
	//Id string `json:"id"`
	GroupId                string `json:"groupId"`
	DefaultPaymentMethodId string `json:"defaultPaymentMethodId"`
	SalesChannelId         string `json:"salesChannelId"`
	LanguageId             string `json:"languageId"`
	//LastPaymentMethodId string `json:"lastPaymentMethodId"`
	DefaultBillingAddressId  string `json:"defaultBillingAddressId"`
	DefaultShippingAddressId string `json:"defaultShippingAddressId"`
	CustomerNumber           string `json:"customerNumber"`
	SalutationId             string `json:"salutationId"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	//Company string `json:"company"`
	Email string `json:"email"`
	//Title                   string   `json:"title"`
	//VatIds                  []string `json:"vatIds"`
	//AffiliateCode           string   `json:"affiliateCode"`
	//CampaignCode            string   `json:"campaignCode"`
	//Active                  bool     `json:"active"`
	//DoubleOptInRegistration bool     `json:"doubleOptInRegistration"`
	//DoubleOptInEmailSentDate time.Time `json:"doubleOptInEmailSentDate"`
	//DoubleOptInConfirmDate   time.Time `json:"doubleOptInConfirmDate"`
	//Hash                     string    `json:"hash"`
	//Guest bool	`json:"guest"`
	//FirstLogin	time.Time	`json:"firstLogin"`
	//LastLogin     time.Time `json:"lastLogin"`
	//Newsletter    bool      `json:"newsletter"`
	//Birthday      string    `json:"birthday"`
	//lastOrderDate time.Time `json:"lastOrderDate"`
	//OrderCount    int64     `json:"orderCount"`
	//CustomFields	[]CustomFields `json:"customFields"`
	//TagIds []string `json:"tagIds"`
	CreatedAt time.Time `json:"createdAt"`
	//UpdatedAt time.Time `json:"updatedAt"`
	//Group                []Group                `json:"group"`
	//DefaultPaymentMethod []DefaultPaymentMethod `json:"defaultPaymentMethod"`
	//Language             []Language             `json:"language"`
	//LastPaymentMethod []LastPaymentMethod `json:"lastPaymentMethod"`
	//DefaultBillingAddress []DefaultBillingAddress `json:"defaultBillingAddress"`
	//DefaultShippingAddress []DefaultShippingAddress `json:"defaultShippingAddress"`
	//Salutation []Salutation `json:"salutation"`
	//Addresses []Addresses `json:"addresses"`
}

// BillingAddress is the billing address.
type BillingAddress struct {
	//Id         string `json:"id"`
	CustomerId string `json:"customerId"`
	CountryId  string `json:"countryId"`
	//CountryStateId string `json:"countryStateId"`
	SalutationId string `json:"salutationId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Zipcode      string `json:"zipcode"`
	City         string `json:"city"`
	//Company
	Street string `json:"street"`
	//Department
	//Title
	//PhoneNumber
	//additionalAddressLine1
	//additionalAddressLine2
	//CustomFields
	//Country []Country
	//CountryState []CountryState
	//Salutation []Salutation
}

//// ShippingAddress is the shipping address.
//type ShippingAddress struct {
//	//Id         string `json:"id"`
//	CustomerId string `json:"customerId"`
//	CountryId      string `json:"countryId"`
//	//CountryStateId string `json:"countryStateId"`
//	SalutationId string `json:"salutationId"`
//	FirstName    string `json:"firstName"`
//	LastName     string `json:"lastName"`
//	Zipcode      string `json:"zipcode"`
//	City         string `json:"city"`
//	//Company
//	Street string `json:"street"`
//	//Department
//	//Title
//	//PhoneNumber
//	//additionalAddressLine1
//	//additionalAddressLine2
//	//CustomFields
//	//Country []Country
//	//CountryState []CountryState
//	//Salutation []Salutation
//}
