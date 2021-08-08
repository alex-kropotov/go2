package main

import (
	"fmt"
	"reflect"
	"strings"
)

func sqlConvert(sqlText string, params ...interface{}) (sqlTextResult string, paramsResult []interface{}) {

	paramsResult = []interface{}(nil)
	sqlTextParts := strings.Split(sqlText, "?")
	sqlTextResult = sqlTextParts[0]
	for i, val := range params {
		rv := reflect.ValueOf(val)
		if rv.Kind() == reflect.Slice {
			blank := ""
			for j := 0; j < rv.Len(); j++ {
				paramsResult = append(paramsResult, rv.Index(j).Interface())
				if j < rv.Len() -1 {
					blank += "?,"
				} else {
					blank += "?"
				}
			}
			sqlTextResult = sqlTextResult + blank + sqlTextParts[i+1]
		} else {
			paramsResult = append(paramsResult, interface{}(val))
			sqlTextResult = sqlTextResult + "?" + sqlTextParts[i+1]
		}
	}
	return sqlTextResult, paramsResult
}

func main() {
	a, b := sqlConvert("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?",
		false, []int{1, 6, 234}, 555 )
	fmt.Println(a, b)
	//
	// SELECT * FROM table WHERE deleted = ? AND id IN(?,?,?) AND count < ? [false 1 6 234 555]
	//
}