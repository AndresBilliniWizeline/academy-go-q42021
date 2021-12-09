package structs

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"challenge/api/errorsHandlers"
)

type Query struct {
	Type           string `json:"type"`
	Items          int    `json:"items"`
	ItemsPerWorker int    `json:"items_per_worker"`
}

// Validates that the query param type is odd or even
func (q *Query) ValidateType() bool {
	qType := strings.ToLower(q.Type)
	return qType == "odd" || qType == "even"
}

// Validates that the query param Items is a natural number
func (q *Query) ValidateItems() bool {
	return q.Items > 0
}

// Validates that the query param Items Per Worker is a natural number
func (q *Query) ValidateItemsPerWorker() bool {
	return q.ItemsPerWorker > 0
}

// Transforms the query values to the Query structure
func (q *Query) SetValues(rawQuery url.Values) {
	q.Type = rawQuery.Get("type")
	Items, itemsErr := strconv.Atoi(rawQuery.Get("items"))
	errorsHandlers.CheckNilErr(itemsErr)
	q.Items = Items
	ItemsPerWorker, itemsPerWorkerError := strconv.Atoi(rawQuery.Get("items_per_worker"))
	errorsHandlers.CheckNilErr(itemsPerWorkerError)
	q.ItemsPerWorker = ItemsPerWorker
}

// Sends an error message if at least one of the query params is missing
func (q *Query) SendErrorMessage(w http.ResponseWriter, element string, multiple int) {
	var message string
	if multiple > 1 {
		message = "The " + element + " params are not valid"
	} else {
		message = "The " + element + " param is not valid"
	}
	http.Error(w, message, http.StatusBadRequest)
}

// Handles the error message and how many query params are missing
func (q *Query) HandleError() (string, int) {
	var queryError string
	var typeError string
	var itemsError string
	var itemsPerWorkerError string
	var connector1 string
	var connector2 string
	multiple := 0

	if !q.ValidateType() {
		typeError = "type"
		multiple += 1
	}
	if !q.ValidateItems() {
		if len(typeError) > 0 {
			itemsError = "items"
		} else {
			typeError = "items"
		}
		multiple += 1
	}
	if !q.ValidateItemsPerWorker() {
		itemsPerWorkerError = "items per worker"
		multiple += 1
	}
	if multiple == 2 {
		connector1 = " and "
	}
	if multiple == 3 {
		connector1 = ", "
		connector2 = ", and "
	}
	queryError = typeError + connector1 + itemsError + connector2 + itemsPerWorkerError
	return queryError, multiple
}
