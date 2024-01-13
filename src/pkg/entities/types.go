package entities

import "time"

type ColumnDate struct {
	ChangedAt *time.Time `mapstructure:"changed_at,omitempty"`
	Date      time.Time  `mapstructure:"date,omitempty"`
}

func (c *ColumnDate) GetChangedAt() *time.Time {
	if c == nil {
		return nil
	}

	return c.ChangedAt
}

func (c *ColumnDate) GetDate() *time.Time {
	if c == nil {
		return nil
	}

	return &c.Date
}

type ColumnNumeric struct {
	Unit  string  `mapstructure:"unit,omitempty"`
	Value float64 `mapstructure:"value,omitempty"`
}

func (c *ColumnNumeric) Get() *float64 {
	if c == nil {
		return nil
	}

	return &c.Value
}

type Label struct {
	Index int    `mapstructure:"index,omitempty"`
	Text  string `mapstructure:"text,omitempty"`
}

type ColumnStatus struct {
	Label Label `mapstructure:"label"`
}

func (c *ColumnStatus) GetText() *string {
	if c == nil {
		return nil
	}

	return &c.Label.Text
}

type ColumnString struct {
	Value string `mapstructure:",omitempty"`
}

func (c *ColumnString) Get() *string {
	if c == nil {
		return nil
	}

	return &c.Value
}

type ColumnName struct {
	Name string `mapstructure:",omitempty"`
}

type Fields map[string]interface{}

type Columns struct {
	Fields Fields `mapstructure:",remain"`
}
