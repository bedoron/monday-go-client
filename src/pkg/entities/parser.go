package entities

import "github.com/mitchellh/mapstructure"

func Parse[T any](src map[string]interface{}, decoders mapstructure.DecodeHookFunc) (*T, error) {
	var t T
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: decoders,
		Result:     &t,
	})

	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(src); err != nil {
		return nil, err
	}
	return &t, nil
}
