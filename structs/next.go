package structs

import (
	"challenge/api/errorsHandlers"
	"strconv"
	"strings"
)

// Structure for handling next pagination pokemon or previous
type Next struct {
	Offset int
	Limit  int
}

// Destructures the url query params to set them to a next structure
func (n *Next) SetInfo(response string) {
	split := strings.Split(response, "?")
	params := strings.Split(split[1], "&")
	offsetString := strings.Split(params[0], "offset=")[1]
	limitString := strings.Split(params[1], "limit=")[1]
	offset, err := strconv.Atoi(offsetString)
	errorsHandlers.CheckNilErr(err)
	limit, err := strconv.Atoi(limitString)
	errorsHandlers.CheckNilErr(err)
	n.Offset = offset
	n.Limit = limit
}

// Combines next attributes to return a new url with query params
func (n *Next) GetUrl(url string) string {
	offset := strconv.Itoa(n.Offset)
	limit := strconv.Itoa(n.Limit)
	newUrl := url + "?offset=" + offset + "&limit=" + limit
	return newUrl
}
