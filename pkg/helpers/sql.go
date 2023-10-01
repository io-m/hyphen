package helpers

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// ConstructInsertStatement is function that generates INSERT SQL statement taking into account optional fields for Golang struct
func ConstructInsertStatement(baseTable string, requiredFields map[string]any, optionalFields map[string]any) (string, []any, error) {
	fieldNames := ""
	placeholders := ""
	values := []any{}
	if len(requiredFields) == 0 || len(optionalFields) == 0 {
		return "", nil, errors.New("not any field passed")
	}
	index := 1
	for rf, rv := range requiredFields {
		if fieldNames != "" {
			fieldNames += ", "
			placeholders += ", "
		}
		fieldNames += rf
		placeholders += "$" + strconv.Itoa(index)
		values = append(values, rv)
		index++
	}
	for field, value := range optionalFields {
		if !reflect.ValueOf(value).IsNil() {
			if fieldNames != "" {
				fieldNames += ", "
				placeholders += ", "
			}
			fieldNames += field
			placeholders += "$" + strconv.Itoa(index)
			values = append(values, value)
			index++
		}
	}

	finalStatement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id;", baseTable, fieldNames, placeholders)
	return finalStatement, values, nil
}

// ConstructUpdateStatement is function that generates UPDATE SQL statement since every field in Update structures can be nil
func ConstructUpdateStatement(baseTable string, updateData map[string]any, condition string) (string, []any, error) {
	setClause := ""
	values := []any{}
	if len(updateData) == 0 {
		return "", nil, errors.New("not any field passed")
	}

	index := 1
	for field, value := range updateData {
		if !reflect.ValueOf(value).IsNil() {
			if setClause != "" {
				setClause += ", "
			}
			setClause += fmt.Sprintf("%s = $%d", field, index)
			values = append(values, value)
			index++
		}
	}

	finalStatement := fmt.Sprintf("UPDATE %s SET %s WHERE %s;", baseTable, setClause, condition)
	return finalStatement, values, nil
}
