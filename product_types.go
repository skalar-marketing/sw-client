package sw

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
	Sorting  string     `json:"sorting"`
	Elements []Category `json:"elements"`
}
