package gurdurr

import (
	"fmt"
	"strings"

	"github.com/Aranyak-Ghosh/golist"
)

type ComparatorType int

const (
	GreaterThanOrEqual ComparatorType = iota
	GreaterThan
	Equals
	LessThan
	LessThanOrEqual
	NotEquals
	Like
	Not
)

func (ct ComparatorType) String() string {
	switch ct {
	case GreaterThanOrEqual:
		return ">="
	case GreaterThan:
		return ">"
	case Equals:
		return "=="
	case LessThan:
		return "<"
	case LessThanOrEqual:
		return "<="
	case NotEquals:
		return "!="
	case Like:
		return "LIKE"
	case Not:
		return "!"
	default:
		return ""
	}
}

type Clause int

const (
	Select Clause = iota
	Where
	OrderBy
	Top
	Limit
	Join
)

type QueryConnector = int

const (
	AND QueryConnector = iota
	OR
)

/***
SELECT [] <- COLUMNS
FROM [] <- TABLE_NAME
JOIN [] <- TABLE_NAME ON ([] <- JOIN_PROPERTY = [] <- JOIN_PROPERTY ) <- COMPARE CLAUSE
WHERE [] <- COMPARE CLAUSE
ORDER BY <-
*/

type queryPart struct {
	clause     Clause
	columnName []string
	comparator []ComparatorType
	tableName  string
}

type whereQueryPart struct {
	columnName string
	comparator ComparatorType
	value      any
	connector  QueryConnector
}

type whereQuery struct {
	queryPart []whereQueryPart
}

type queryObject struct {
	queryString string
	selector    queryPart
	whereFilter golist.Queue[whereQueryPart]
	tableName   string
}

func (q *queryObject) Select(columns []string) *queryObject {
	var part queryPart

	part.clause = Select
	copy(part.columnName, columns)

	selectedColumns := strings.Join(columns, ",")

	selectStatement := fmt.Sprintf("SELECT %s", selectedColumns)
	q.selector = part
	q.queryString = fmt.Sprintf("%s %s", selectStatement, q.queryString)

	return q
}

func (q *queryObject) Where(columnName string, coparator ComparatorType, value any) *queryObject {
	queryString := fmt.Sprintf("%s %s ?", columnName, coparator.String())

	return q
}

func NewQuery(tableName string) *queryObject {
	return &queryObject{
		tableName:   tableName,
		queryString: fmt.Sprintf("FROM %s", tableName),
	}
}
