package gurdurr

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type QueryExecutor interface {
	Exec(*queryObject) (QueryResult, error)
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

const (
	MSSQL    string = "sqlserver"
	Postgres string = "postgres"
	MySQL    string = "mysql"
	SQLITE   string = "sqlite3"
)

func NewRepository(conn *sql.DB, driver string) QueryExecutor {
	db := sqlx.NewDb(conn, driver)
	db.Ping()

	return &queryExecutor{
		db: db,
	}
}
