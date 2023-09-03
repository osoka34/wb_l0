package utils

import (
	"fmt"
	"reflect"
)

func StructToSqlArray(data interface{}) string {

	// Используем отражение для получения значений полей структуры
	v := reflect.ValueOf(data)
	//t := v.Type()
	numFields := v.NumField()

	// Создаем слайс для хранения значений полей
	values := make([]string, 0)
	for i := 0; i < numFields; i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Struct || field.Kind() == reflect.Slice || field.Kind() == reflect.Pointer {
			continue
		}
		if field.Kind() == reflect.String {
			values = append(values, fmt.Sprintf("'%s'", field.Interface()))
			//values[i] = fmt.Sprintf("'%s'", field.Interface())
		} else {
			values = append(values, fmt.Sprintf("%v", field.Interface()))
			//values[i] = fmt.Sprintf("%v", field.Interface())
		}
	}

	// Формируем строку SQL запроса
	sqlValues := "(" + fmt.Sprintf("%s", values[0])
	for i := 1; i < len(values); i++ {
		sqlValues += ", " + fmt.Sprintf("%s", values[i])
	}
	sqlValues += ")"

	return sqlValues
}
