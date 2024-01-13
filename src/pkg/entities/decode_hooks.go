package entities

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

const mondayPartialDateLayout = "2006-01-02 15:04"

var BasicDecodeHooks = mapstructure.ComposeDecodeHookFunc(
	mapstructure.StringToTimeDurationHookFunc(),
	mapstructure.StringToSliceHookFunc(","),
	DecodeStringToTimeHook(),
)

func DecodeStringToTimeHook() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String || t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		in := data.(string)

		res, err := DecodeTimeBestEffort(in)
		if err == nil {
			return *res, nil
		}

		return data, err
	}
}

func DecodeTimeBestEffort(in string) (*time.Time, error) {
	res, err := time.Parse(time.RFC3339, in)
	if err == nil {
		return &res, nil
	}

	res, err = time.Parse("2006-01-02", in)
	if err == nil {
		return &res, nil
	}

	res, err = time.Parse(mondayPartialDateLayout, in)
	if err == nil {
		return &res, nil
	}

	in = strings.TrimSuffix(in, " UTC")
	res, err = time.Parse(time.DateTime, in)
	if err == nil {
		return &res, nil
	}

	return nil, fmt.Errorf("unknown date format %v", in)
}
