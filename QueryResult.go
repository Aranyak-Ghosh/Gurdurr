package gurdurr

import (
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
)

type queryResult struct {
	res *sqlx.Rows
}

type QueryResult interface {
	Result(any) error
}

func (q *queryResult) Result(data any) error {
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
		out := make([]interface{}, 0)
		if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
			for q.res.Next() {

				var val = reflect.New(v.Type().Elem()).Elem().Interface()

				if err := q.res.StructScan(&val); err != nil {
					return err
				}

				out = append(out, val)
			}
			data = &out
		} else {
			q.res.Next()
			if err := q.res.StructScan(&data); err != nil {
				return err
			}
		}
	}
	return fmt.Errorf("Value must be of type pointer")
}

// func Result[T any](q *queryResult) T {
// 	var out T

// 	for q.res.Next() {

// 	}
// }
