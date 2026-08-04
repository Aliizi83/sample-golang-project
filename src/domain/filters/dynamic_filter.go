package filters

type Sort struct {
	ColumnId string `json:"columnId"`
	Sort     string `json:"sort"`
}

type Filter struct {
	Type string `json:"type"`
	From string `json:"from"`
	To   string `json:"to"`

	FilterType string `json:"filterType"`
}

type DynamicFilter struct {
	Sorts   *[]Sort           `json:"sorts"`
	Filters map[string]Filter `json:"filters"`
}
