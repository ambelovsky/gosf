package gosf

import (
	"github.com/ambelovsky/go-structs"
	"github.com/mitchellh/mapstructure"
)

// StructToMap converts the given structure into a map[string]interface{}
func StructToMap(input interface{}) map[string]interface{} {
	return structs.Map(input)
}

// MapToStruct converts the given map[string]interface{} into a struct
func MapToStruct(input map[string]interface{}, output interface{}) error {
	return mapstructure.Decode(input, output)
}
