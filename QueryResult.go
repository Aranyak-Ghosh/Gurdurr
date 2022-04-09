package gurdurr

import (
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
	if reflect.ValueOf(data).Kind() == reflect.Array {
		i := data.([]any)
		for q.res.Next() {
			var val interface{}

			if err := q.res.StructScan(&val); err != nil {
				return err
			}
			i = append(i, val)
		}
		data = i
	} else {
		q.res.Next()
		if err := q.res.StructScan(&data); err != nil {
			return err
		}
	}
	return nil
}
