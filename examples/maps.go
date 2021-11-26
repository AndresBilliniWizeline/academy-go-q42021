package examples

import "fmt"

func maps() {
	score := make(map[string]int)
	score["andres"] = 88
	score["frodo"] = 35
	score["sam"] = 84
	score["gandalf"] = 8888
	score["aragorn"] = 880

	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}
}
