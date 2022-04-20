package sw

import "time"

// Category is a product category.
type Category struct {
	Id                     string   `json:"id"`
	VersionId              string   `json:"versionId"`
	ParentId               string   `json:"parentId"`
	ParentVersionId        string   `json:"parentVersionId"`
	AfterCategoryId        string   `json:"afterCategoryId"`
	AfterCategoryVersionId string   `json:"afterCategoryVersionId"`
	MediaId                string   `json:"mediaId"`
	DisplayNestedProducts  bool     `json:"displayNestedProducts"`
	Breadcrumb             []string `json:"breadcrumb"`
	Level                  int64    `json:"level"`
	Path                   string   `json:"path"`
	ChildCount             int64    `json:"childCount"`
	Type                   string   `json:"type"`
	ProductAssignmentType  string   `json:"productAssignmentType"`
	Visible                bool     `json:"visible"`
	Active                 bool     `json:"active"`
	Name                   string   `json:"name"`
	//customFields TODO
	LinkType         string    `json:"linkType"`
	InternalLink     string    `json:"internalLink"`
	ExternalLink     string    `json:"externalLink"`
	LinkNewTab       bool      `json:"linkNewTab"`
	Description      string    `json:"description"`
	MetaTitle        string    `json:"metaTitle"`
	MetaDescription  string    `json:"metaDescription"`
	Keywords         string    `json:"keywords"`
	CmsPageId        string    `json:"cmsPageId"`
	CmsPageVersionId string    `json:"cmsPageVersionId"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	//translated
	Parent   *Category  `json:"parent"`
	Children []Category `json:"children"`
	//media
	//cmsPage
	//seoUrls
}

// Filter is used to filter results.
type Filter struct {
	Type  string `json:"type"`
	Field string `json:"field"`
	Value string `json:"value"`
}

// Sort is used to sort results.
type Sort struct {
	Field          string `json:"field"`
	Order          string `json:"order"`
	NaturalSorting bool   `json:"naturalSorting"`
}

// Aggregation is used to aggregate results.
type Aggregation struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Field string `json:"field"`
}

// ListCategoryFilter is used to filter categories.
type ListCategoryFilter struct {
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
	Filter     []Filter `json:"filter"`
	Sort       []Sort   `json:"sort"`
	PostFilter []Filter `json:"post-filter"`
	//associations TODO
	Aggregations []Aggregation `json:"aggregations"`
	Grouping     []string      `json:"grouping"`
}

// GetCategoryFilter is used to get a single category.
type GetCategoryFilter struct {
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
	Filter     []Filter `json:"filter"`
	Sort       []Sort   `json:"sort"`
	PostFilter []Filter `json:"post-filter"`
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
	ShippingFreeFilter bool          `json:"	shipping-free-filter"`
	PropertyFilter     bool          `json:"property-filter"`
	PropertyWhitelist  string        `json:"property-whitelist"`
	ReduceAggregations string        `json:"reduce-aggregations"`
}

// ListCategoryResult contains a list of categories.
type ListCategoryResult struct {
	ApiAlias string `json:"apiAlias"`
	Entity   string `json:"entity"`
	Total    int    `json:"total"`
	//aggregations TODO
	Page     int        `json:"page"`
	Limit    int        `json:"limit"`
	Elements []Category `json:"elements"`
}
