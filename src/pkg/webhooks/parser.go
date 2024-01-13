package webhooks

import "github.com/bedoron/monday-go-client/src/pkg/entities"

func Parse[T any](src map[string]interface{}) (*T, error) {
	return entities.Parse[T](src, decodeHooks)
}

func parse[T any](src map[string]interface{}) (*T, error) {
	return entities.Parse[T](src, entities.BasicDecodeHooks)
}
