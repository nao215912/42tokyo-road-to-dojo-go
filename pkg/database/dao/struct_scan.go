package dao

import (
	"database/sql"
	"reflect"
)

func structScan(v interface{}, row *sql.Row) error {
	rv := reflect.ValueOf(v).Elem()
	n := rv.NumField()
	values := make([]interface{}, n)
	rv = reflect.Indirect(rv)

	for i := 0; i < n; i++ {
		values[i] = rv.Field(i).Addr().Interface()
	}
	return row.Scan(values...)
}
