package graphql

import (
	"reflect"
	"strings"

	"github.com/bedoron/monday-go-client/src/pkg/entities"
)

type fieldsTagMapper map[string]map[string]reflect.Value

const idTag = "mapstructure"

type GraphQLLoader struct {
	tableRow    interface{}
	fieldsByTag fieldsTagMapper
}

func NewLoader[T any](tableRow *T) *GraphQLLoader {
	loader := &GraphQLLoader{
		tableRow:    tableRow,
		fieldsByTag: make(fieldsTagMapper),
	}

	buildFieldsByTagMap(idTag, tableRow, loader.fieldsByTag)
	return loader
}

func (gqll *GraphQLLoader) Load(gqlItem *GQLItem) {
	for _, column := range gqlItem.Columns {
		if column.Empty() {
			// logger.Debugf("column '%v' is empty", column.Id)
			continue
		}

		colVal, ok := gqll.fieldsByTag[idTag][column.Id]
		if !ok {
			logger.Warnf("column '%v' undefined, skipping", column.Id)
			continue
		}

		var refVal reflect.Value
		switch column.Type {
		case entities.ValueTypeDate, ValueTypeDateCreation:
			c := &entities.ColumnDate{
				Date: column.GQLDate.Date,
			}

			refVal = reflect.ValueOf(c)
		case entities.ValueTypeNumeric, ValueTypeNumericGQL:
			c := &entities.ColumnNumeric{
				Value: column.GQLNumeric.Value,
			}

			refVal = reflect.ValueOf(c)
		case ValueTypeStatusGQL:
			c := &entities.ColumnStatus{
				Label: entities.Label{
					Text: column.GQLStatus.Text,
				},
			}

			refVal = reflect.ValueOf(c)
		case ValueTypeText, ValueTypePeopleGQL:
			c := &entities.ColumnString{
				Value: column.GQLText.Text,
			}

			refVal = reflect.ValueOf(c)
		}

		if !refVal.IsValid() {
			logger.Warnf("column '%v' wasn't parsed, type '%v' unsopported", column.Id, column.Type)
			continue
		}

		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Warnf("failed assigning column '%v': %v", column.Id, r)
				}
			}()

			if colVal.Kind() == reflect.Pointer {
				colVal.Set(refVal)
			} else {
				colVal.Set(refVal.Elem())
			}
		}()

	}
}

func buildFieldsByTagMap(key string, tableRow interface{}, fields fieldsTagMapper) {
	rt := reflect.TypeOf(tableRow)
	var value reflect.Value
	if rt.Kind() == reflect.Pointer {
		ptr := reflect.ValueOf(tableRow)
		value = ptr.Elem()
		rt = value.Type()
	}

	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}

	if !value.IsValid() {
		value = reflect.ValueOf(tableRow)
	}

	if fields[key] == nil {
		fields[key] = make(map[string]reflect.Value)
	}

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get(key), ",")[0]
		if v == "" || v == "-" {
			continue
		}

		fields[key][v] = value.Field(i)
	}
}
