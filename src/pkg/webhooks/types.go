package webhooks

import (
	"time"

	"github.com/bedoron/monday-go-client/src/pkg/entities"
)

const (
	EventCreate       = "create_pulse"
	EventDelete       = "delete_pulse"
	EventUpdateName   = "update_name"
	EventUpdateColumn = "update_column_value"
)

type Challenge struct {
	Challenge string `mapstructure:"challenge"`
}

type EventMetadata struct {
	Board     int       `mapstructure:"boardId"`
	Type      string    `mapstructure:"type"`
	PulseId   uint64    `mapstructure:"pulseId"`
	PulseName *string   `mapstructure:"pulseName,omitempty"`
	UserId    int       `mapstructure:"userId"`
	Timestamp time.Time `mapstructure:"triggerTime"`
}

type Event struct {
	EventMetadata EventMetadata `mapstructure:",squash"`

	UpdateColumnValue        *UpdateColumnValue        `mapstructure:",squash"`
	UpdateColumnValueDate    *UpdateColumnValueDate    `mapstructure:",squash"`
	UpdateColumnValueNumeric *UpdateColumnValueNumeric `mapstructure:",squash"`
	UpdateColumnValueStatus  *UpdateColumnValueStatus  `mapstructure:",squash"`
	UpdateColumnValueText    *UpdateColumnValueText    `mapstructure:",squash"`

	UpdateName *UpdateName `mapstructure:",squash"`
	DeleteItem *DeleteItem `mapstructure:",squash"`
	CreateItem *CreateItem `mapstructure:",squash"`
}

type Payload struct {
	Challenge Challenge `mapstructure:",squash"`
	Event     Event     `mapstructure:"event"`
}

type UpdateColumnValue struct {
	FieldId   string `mapstructure:"columnId"`
	Field     string `mapstructure:"columnTitle"`
	ValueType string `mapstructure:"columnType"`
}

type UpdateColumnValueDate struct {
	Value entities.ColumnDate
}

type UpdateColumnValueText struct {
	Value         entities.ColumnString
	PreviousValue entities.ColumnString
}

type UpdateColumnValueNumeric struct {
	Value entities.ColumnNumeric
}

type UpdateColumnValueStatus struct {
	Value entities.ColumnStatus
}

type DeleteItem struct {
	ItemId   uint64 `mapstructure:"itemId"`
	ItemName string `mapstructure:"itemName"`
}

type CreateItem struct {
	PulseName    string           `mapstructure:"pulseName"`
	ColumnValues entities.Columns `mapstructure:"columnValues"`
}

type UpdateName struct {
	PreviousValue entities.ColumnName `mapstructure:",omitempty"`
	Value         entities.ColumnName `mapstructure:",omitempty"`
}
