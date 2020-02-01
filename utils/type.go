package utils

import (
	"reflect"
	"strconv"
)

func StructToMap(structPtr interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	elem := reflect.ValueOf(structPtr).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		res[relType.Field(i).Tag.Get("json")] = elem.Field(i).Interface()
	}
	return res
}

func ValidateInt(str string, maxLen int) int {
	if str == "" || len(str) > maxLen {
		return 0
	}
	res, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return res
}