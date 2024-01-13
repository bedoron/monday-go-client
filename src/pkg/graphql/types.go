package graphql

import "time"

const (
	ValueTypeDateCreation = "creation_log" // graphql only
	ValueTypeStatus       = "color"        // webhook only
	ValueTypeStatusGQL    = "status"       // graphql only

	ValueTypeNumericGQL = "numbers" // graphql only
	ValueTypeText       = "text"    // graphql only
	ValueTypePeopleGQL  = "people"  // graphql only
)

type GQLUpdateResponse struct {
	Errors *[]string       `mapstructure:"errors,omitempty"`
	Data   *map[string]any `mapstructure:"data,omitempty"`
}

type GQLResponse struct {
	Data   *GQLData `mapstructure:"data,omitempty"`
	Errors *[]GQLError
}

type GQLError struct {
	Message   string
	Locations []GQLErrorLocation
}

type GQLErrorLocation struct {
	Line   int
	Column int
}

type GQLData struct {
	Boards []GQLBoard `mapstructure:"boards"`
}

type GQLBoard struct {
	Columns   *[]GQLColumnMetadata `mapstructure:"columns"`
	BoardId   *int64               `mapstructure:"id"`
	ItemsPage GQLItemsPage         `mapstructure:"items_page"`
}

type GQLItemsPage struct {
	Cursor *string   `mapstructure:"cursor"`
	Items  []GQLItem `mapstructure:"items"`
}

type GQLColumnMetadata struct {
	Id    string `mapstructure:"id"`
	Title string `mapstructure:"title"`
}

type GQLItem struct {
	Created *time.Time  `mapstructure:"created_at"`
	Id      *uint64     `mapstructure:"id"`
	Url     string      `mapstructure:"name"`
	Columns []GQLColumn `mapstructure:"column_values"`
}

type GQLColumn struct {
	Type        string  `mapstructure:"type"`
	Id          string  `mapstructure:"id"`
	Text        *string `mapstructure:"text"`
	Value       *string `mapstructure:"value"`
	*GQLText    `mapstructure:",omitempty"`
	*GQLNumeric `mapstructure:",omitempty"`
	*GQLStatus  `mapstructure:",omitempty"`
	*GQLDate    `mapstructure:",omitempty"`
}

func (gqlC GQLColumn) Empty() bool {
	return gqlC.GQLText == nil &&
		gqlC.GQLNumeric == nil &&
		gqlC.GQLStatus == nil &&
		gqlC.GQLDate == nil
}

type GQLText struct {
	Text string
}

type GQLNumeric struct {
	Value float64
}

type GQLStatus struct {
	Text string
}

type GQLDate struct {
	Date time.Time
}
