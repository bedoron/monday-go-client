package graphql

import (
	"reflect"
	"strconv"

	"github.com/bedoron/monday-go-client/src/pkg/entities"
	"github.com/bedoron/monday-go-client/src/pkg/webhooks"
	"github.com/mitchellh/mapstructure"
)

var gqlDecodeHooks = mapstructure.ComposeDecodeHookFunc(
	entities.BasicDecodeHooks,
	decodeGQLColumnHook(),
)

func decodeGQLColumnHook() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.Map || t != reflect.TypeOf(GQLColumn{}) {
			return data, nil
		}

		d := data.(map[string]interface{})

		g, err := webhooks.Parse[GQLColumn](d)
		if err != nil {
			return data, err
		}

		switch g.Type {
		case entities.ValueTypeDate, ValueTypeDateCreation:
			if g.Text == nil || g.Value == nil { // Date depends on both
				return g, nil
			}

			g.GQLDate = &GQLDate{}
			t, err := entities.DecodeTimeBestEffort(*g.Text)
			if err != nil {
				return data, err
			}
			g.GQLDate = &GQLDate{
				Date: *t,
			}
		case entities.ValueTypeNumeric, ValueTypeNumericGQL:
			if g.Text == nil || len(*g.Text) == 0 {
				return g, nil
			}
			g.GQLNumeric = &GQLNumeric{}

			// i, err := strconv.ParseInt(*g.Text, 10, 32)
			f, err := strconv.ParseFloat(*g.Text, 64)
			if err != nil {
				return data, err
			}
			g.GQLNumeric = &GQLNumeric{
				Value: f,
			}
		case entities.ValueTypeText, ValueTypePeopleGQL:
			if g.Text == nil || len(*g.Text) == 0 {
				return g, nil
			}

			g.GQLText = &GQLText{
				Text: *g.Text,
			}
		case entities.ValueTypeStatus, ValueTypeStatusGQL:
			if g.Text == nil || len(*g.Text) == 0 {
				return g, nil
			}

			g.GQLStatus = &GQLStatus{
				Text: *g.Text,
			}
		}

		return g, nil
	}
}
