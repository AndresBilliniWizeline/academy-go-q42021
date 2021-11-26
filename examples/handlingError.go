package examples

import (
	"errors"
)

func SumOf(start, end int) (int, error) {
	if start > end {
		return 0, errors.New("start is greater than end")
	}
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total, nil
}
