package graphql

import (
	"github.com/mitchellh/mapstructure"
)

func Parse[T any](src map[string]interface{}) (*T, error) {
	var t T
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook:       gqlDecodeHooks,
		Result:           &t,
		WeaklyTypedInput: true,
	})

	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(src); err != nil {
		return nil, err
	}
	return &t, nil
}
