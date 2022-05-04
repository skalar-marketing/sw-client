/* sharedStructs.go is a file that contains structs that are used in different files */

package sw

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
