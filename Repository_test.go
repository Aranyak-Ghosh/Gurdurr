package gurdurr

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Name  string `db:"name"`
	Email string `db:"email"`
}

func TestRepository(t *testing.T) {
	dbConn, err := sql.Open(MySQL, "root:mysqlpassword@tcp(127.0.0.1:3306)/GurdurrTest")

	if err != nil {
		t.Errorf("Failed to connect to db with error %v", err)
	}
	repo := NewRepository(dbConn, MySQL)

	q := NewQuery("user")

	q.Select([]string{"name", "email"})

	res, err := repo.Exec(q)

	if err != nil {
		t.Errorf("Failed to execute query: %v", err)
	}

	var result = make([]User, 0)

	res.Result(&result)

	fmt.Printf("%v", result)

}

func TestSqlxRepository(t *testing.T) {
	dbConn, err := sql.Open(MySQL, "root:mysqlpassword@tcp(127.0.0.1:3306)/GurdurrTest")

	if err != nil {
		t.Errorf("Failed to connect to db with error %v", err)
	}

	db := sqlx.NewDb(dbConn, MySQL)

	var result = make([]User, 0)

	rows, err := db.Queryx("SELECT name,email FROM user")

	if err != nil {
		t.Errorf("Failed to execute query: %v", err)
	}

	for rows.Next() {
		var val User

		if err := rows.StructScan(&val); err != nil {
			t.Errorf("Error Parsing data %v", err)
		}

		result = append(result, val)
	}

	// res.Result(result)

	fmt.Printf("%v", result)

}
