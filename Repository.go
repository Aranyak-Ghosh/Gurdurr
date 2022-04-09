package gurdurr

import "github.com/jmoiron/sqlx"

type QueryExecutor interface {
	Exec(*queryObject) *queryResult
}

type queryExecutor struct {
	db *sqlx.DB
}

func (q *queryExecutor) Exec(ob *queryObject) (QueryResult, error) {
	tx := q.db.MustBegin()
	rows, err := tx.Queryx(ob.queryString)

	if err != nil {
		return nil, err
	} else {
		return &queryResult{
			res: rows,
		}, nil
	}
}
