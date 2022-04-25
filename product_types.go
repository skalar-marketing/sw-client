package sw

import "time"

// Product is a product.
type Product struct {
	Id                           string    `json:"id"`
	VersionId                    string    `json:"versionId"`
	ParentId                     string    `json:"parentId"`
	ParentVersionId              string    `json:"parentVersionId"`
	ManufacturerId               string    `json:"manufacturerId"`
	ProductManufacturerVersionId string    `json:"productManufacturerVersionId"`
	UnitId                       string    `json:"unitId"`
	TaxId                        string    `json:"taxId"`
	CoverId                      string    `json:"coverId"`
	ProductMediaVersionId        string    `json:"productMediaVersionId"`
	DeliveryTimeId               string    `json:"deliveryTimeId"`
	CanonicalProductId           string    `json:"canonicalProductId"`
	CmsPageId                    string    `json:"cmsPageId"`
	CmsPageVersionId             string    `json:"cmsPageVersionId"`
	ProductNumber                string    `json:"productNumber"`
	Stock                        int64     `json:"stock"`
	RestockTime                  int64     `json:"restockTime"`
	Active                       bool      `json:"active"`
	AvailableStock               int64     `json:"availableStock"`
	Available                    bool      `json:"available"`
	IsCloseout                   bool      `json:"isCloseout"`
	DisplayGroup                 string    `json:"displayGroup"`
	MainVariantId                string    `json:"mainVariantId"`
	ManufacturerNumber           string    `json:"manufacturerNumber"`
	Ean                          string    `json:"ean"`
	PurchaseSteps                int64     `json:"purchaseSteps"`
	MaxPurchase                  int64     `json:"maxPurchase"`
	MinPurchase                  int64     `json:"minPurchase"`
	PurchaseUnit                 float64   `json:"purchaseUnit"`
	ReferenceUnit                float64   `json:"referenceUnit"`
	ShippingFree                 bool      `json:"shippingFree"`
	MarkAsTopseller              bool      `json:"markAsTopseller"`
	Weight                       float64   `json:"weight"`
	Width                        float64   `json:"width"`
	Height                       float64   `json:"height"`
	Length                       float64   `json:"length"`
	ReleaseDate                  time.Time `json:"releaseDate"`
	RatingAverage                float64   `json:"ratingAverage"`
	CategoryTree                 []string  `json:"categoryTree"`
	PropertyIds                  []string  `json:"propertyIds"`
	OptionIds                    []string  `json:"optionIds"`
	CategoryIds                  []string  `json:"categoryIds"`
	ChildCount                   int64     `json:"childCount"`
	Sales                        int64     `json:"sales"`
	MetaDescription              string    `json:"metaDescription"`
	Name                         string    `json:"name"`
	Keywords                     string    `json:"keywords"`
	Description                  string    `json:"description"`
	MetaTitle                    string    `json:"metaTitle"`
	PackUnit                     string    `json:"packUnit"`
	PackUnitPlural               string    `json:"packUnitPlural"`
	//customFields
	//calculatedPrice
	//calculatedPrices
	CalculatedMaxPurchase int64 `json:"calculatedMaxPurchase"`
	//calculatedCheapestPrice
	IsNew     bool      `json:"isNew"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	//translated
	Parent *Product `json:"parent"`
	//children
	//deliveryTime
	//tax
	//manufacturer
	//unit
	//cover
	//cmsPage
	//canonicalProduct
	//media
	//crossSellings
	//configuratorSettings
	//productReviews
	//mainCategories
	//seoUrls
	//options
	//properties
	//categories
	//streams
	//categoriesRo
	//seoCategory
}

// ListProductsCategoryFilter is used to filter products for a category.
type ListProductsCategoryFilter struct {
	NoAggregations   string   `json:"no-aggregations"`
	OnlyAggregations string   `json:"only-aggregations"`
	Page             int      `json:"page"`
	Limit            int      `json:"limit"`
	Filter           []Filter `json:"filter"`
	Sort             []Sort   `json:"sort"`
	PostFilter       []Filter `json:"post-filter"`
	//associations TODO
	Aggregations       []Aggregation `json:"aggregations"`
	Grouping           []string      `json:"grouping"`
	Order              string        `json:"order"`
	P                  int           `json:"p"`
	Manufacturer       string        `json:"manufacturer"`
	MinPrice           int           `json:"min-price"`
	MaxPrice           int           `json:"max-price"`
	Rating             int           `json:"rating"`
	ShippingFree       bool          `json:"shipping-free"`
	Properties         string        `json:"properties"`
	ManufacturerFilter bool          `json:"manufacturer-filter"`
	PriceFilter        bool          `json:"price-filter"`
	RatingFilter       bool          `json:"rating-filter"`
	ShippingFreeFilter bool          `json:"shipping-free-filter"`
	PropertyFilter     bool          `json:"property-filter"`
	PropertyWhitelist  string        `json:"property-whitelist"`
	ReduceAggregations string        `json:"reduce-aggregations"`
}

// ListProductsCategoryResult is the result type for products in a category.
type ListProductsCategoryResult struct {
	ApiAlias string `json:"apiAlias"`
	Entity   string `json:"entity"`
	Total    int    `json:"total"`
	//aggregations TODO
	Page  int `json:"page"`
	Limit int `json:"limit"`
	//currentFilters
	//availableSortings
	Sorting  string    `json:"sorting"`
	Elements []Product `json:"elements"`
}
