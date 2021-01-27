package models

import (
	"strconv"
	"strings"
)

const (
	querySep = "__"
)

type Operation struct {
	Field string
	Op    string
}

// postgresql operators.
var postgresOperators = map[string]string{
	"exact":       "= '?'",
	"iexact":      "ILIKE '?'",
	"contains":    "LIKE '%?%'",
	"icontains":   "ILIKE '%?%'",
	"gt":          "> ?",
	"gte":         ">= ?",
	"lt":          "< ?",
	"lte":         "<= ?",
	"eq":          "= ?",
	"ne":          "!= ?",
	"startswith":  "LIKE '?%'",
	"endswith":    "LIKE '%?'",
	"istartswith": "ILIKE ?%'",
	"iendswith":   "ILIKE '%?'",
	"isnull":      "IS NULL",
}

func filter(expr string, values ...interface{}) (out Operation) {
	qr := strings.Split(expr, querySep)
	field := strings.Join(qr[:len(qr)-1], ".")
	lastElem := qr[len(qr)-1]
	if sql, ok := postgresOperators[lastElem]; ok {
		value := ""
		switch v := values[0].(type) {
		case string:
			value = v
		case int:
			value = strconv.Itoa(v)
		case bool:
			if sql != postgresOperators["isnull"] {
				return
			}
		}
		sqlrem := strings.Replace(sql, "?", value, 1)
		out = Operation{Field: field, Op: sqlrem}
	}
	return
}
