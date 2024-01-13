package webhooks

import (
	"fmt"
	"reflect"

	"github.com/bedoron/monday-go-client/src/pkg/entities"
	"github.com/mitchellh/mapstructure"
)

var decodeHooks = mapstructure.ComposeDecodeHookFunc(
	entities.BasicDecodeHooks,
	decodeWebhookHook(),
)

func decodeWebhookHook() mapstructure.DecodeHookFunc {
	return func(f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {

		if f.Kind() != reflect.Map || t != reflect.TypeOf(Event{}) {
			return data, nil
		}

		d := data.(map[string]interface{})

		metadata, err := parse[EventMetadata](d)
		if err != nil {
			return data, err
		}

		e := &Event{
			EventMetadata: *metadata,
		}

		switch metadata.Type {
		case EventDelete:
			if di, err := parse[DeleteItem](d); err != nil {
				return data, err
			} else {
				e.DeleteItem = di
			}
		case EventCreate:
			if ci, err := parse[CreateItem](d); err != nil {
				return data, err
			} else {
				e.CreateItem = ci
			}
		case EventUpdateName:
			if un, err := parse[UpdateName](d); err != nil {
				return data, err
			} else {
				e.UpdateName = un
			}
		case EventUpdateColumn:
			// Get the column type metadata
			if ucv, err := parse[UpdateColumnValue](d); err != nil {
				return data, fmt.Errorf("column information missing: %w", err)
			} else {
				e.UpdateColumnValue = ucv
			}

			switch e.UpdateColumnValue.ValueType {
			case entities.ValueTypeNumeric:
				if ucvn, err := parse[UpdateColumnValueNumeric](d); err != nil {
					return data, err
				} else {
					e.UpdateColumnValueNumeric = ucvn
				}
			case entities.ValueTypeDate:
				if ucvd, err := parse[UpdateColumnValueDate](d); err != nil {
					return data, err
				} else {
					e.UpdateColumnValueDate = ucvd
				}
			case entities.ValueTypeStatus:
				if ucvs, err := parse[UpdateColumnValueStatus](d); err != nil {
					return data, err
				} else {
					e.UpdateColumnValueStatus = ucvs
				}
			case entities.ValueTypeText:
				if ucvt, err := parse[UpdateColumnValueText](d); err != nil {
					return data, err
				} else {
					e.UpdateColumnValueText = ucvt
				}
			}

		}

		return e, nil
	}
}
