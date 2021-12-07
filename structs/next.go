package structs

import (
	"strconv"
	"strings"
)

type Next struct {
	Offset int
	Limit  int
}

func (n *Next) SetNext(response string) {
	split := strings.Split(response, "?")
	params := strings.Split(split[1], "&")
	offsetString := strings.Split(params[0], "offset=")[1]
	limitString := strings.Split(params[1], "limit=")[1]
	offset, _ := strconv.Atoi(offsetString)
	limit, _ := strconv.Atoi(limitString)
	n.Offset = offset
	n.Limit = limit
}

func (n *Next) GetNextUrl(url string) string {
	offset := strconv.Itoa(n.Offset)
	limit := strconv.Itoa(n.Limit)
	newUrl := url + "?offset=" + offset + "&limit=" + limit
	return newUrl
}
