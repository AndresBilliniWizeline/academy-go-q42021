package examples

import "fmt"

func Superman() {
	fmt.Println("I'm superman")
}

func Adder(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	length := len(values)
	name := "Just for fun"
	return sum, length, name
}

func Multiply(v1, v2 int) int {
	return v1 * v2
}
