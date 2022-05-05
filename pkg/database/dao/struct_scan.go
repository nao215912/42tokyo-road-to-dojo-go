package dao

import (
	"database/sql"
	"reflect"
)

func structScan(v interface{}, row *sql.Row) error {
	rv := reflect.ValueOf(v).Elem()
	values := make([]interface{}, rv.NumField())
	rv = reflect.Indirect(rv)

	for i := 0; i < rv.NumField(); i++ {
		values[i] = rv.Field(i).Addr().Interface()
	}
	return row.Scan(values...)
}
