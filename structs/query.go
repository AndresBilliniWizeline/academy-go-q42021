package structs

import (
	"net/url"
	"strconv"
)

type Query struct {
	Type           string `json:"type"`
	Items          int    `json:"items"`
	ItemsPerWorker int    `json:"items_per_worker"`
}

func (q *Query) ValidateType() bool {
	return q.Type == "odd" || q.Type == "even"
}

func (q *Query) SetValues(rawQuery url.Values) {
	q.Type = rawQuery.Get("type")
	q.Items, _ = strconv.Atoi(rawQuery.Get("items"))
	q.ItemsPerWorker, _ = strconv.Atoi(rawQuery.Get("items_per_worker"))
}
