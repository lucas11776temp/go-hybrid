package builder

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	fields []string
	table  string
	where  map[uint32][]string
	limit  int
}

func (q *QueryBuilder) Select(fields []string) *QueryBuilder {
	return q
}

func (q *QueryBuilder) Table(name string) *QueryBuilder {
	return q
}

func (q *QueryBuilder) Where(clue []string) *QueryBuilder {
	if q.where == nil {
		q.where = make(map[uint32][]string)
	}

	q.where[uint32(len(q.where))] = []string{"age", ">", "10"}
	// or

	var orWhere map[uint32]map[string][]string = make(map[uint32]map[string][]string)

	fmt.Println(orWhere)

	return q
}

func (q *QueryBuilder) Limit(name uint32) *QueryBuilder {
	return q
}

func (q *QueryBuilder) Get() *QueryBuilder {
	return q
}

func (q *QueryBuilder) whereQueryBuilder() string {

	var where map[uint32]map[string][]string = make(map[uint32]map[string][]string)

	where[0] = map[string][]string{}

	where[0][""] = []string{"name", "=", "lucas11776"}
	where[0]["or"] = []string{"age", "<", "12"}
	where[0]["and"] = []string{"email", "LIKE", "lucas11776"}

	query := ""

	for wk := range where {
		ws := len(where[wk])

		if ws > 1 {
			query += "("
		}

		for qk, qv := range where[wk] {
			qs := len(qv)

			if qk != "" {
				query += " " + qk + " "
			}

			if qs > 2 {
				if strings.ToUpper(qv[1]) == "LIKE" {
					query += qv[0] + " LIKE " + "%" + qv[2] + "%"
				} else {
					query += qv[0] + " " + qv[1] + " " + qv[2]
				}
			}

			if qs == 2 {
				query += qv[0] + " = " + qv[2]
			}
		}

		if ws > 1 {
			query += ")"
		}
	}

	return query
}
