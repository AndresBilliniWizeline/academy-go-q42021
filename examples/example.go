package examples

import "fmt"

func main() {
	var init float64 = 5
	init += 23.5
	grades := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(
		"Init es igual a:", init,
	)
	fmt.Println(len(grades), grades)
	a := grades[3:]
	slice := make([]int, 3, 5)
	slice = append(slice, 1, 2, 10)

	fmt.Println(a, slice)
	fmt.Println(
		"Length: ", len(slice),
		"Capacity: ", cap(slice),
	)
	for key, value := range slice {
		fmt.Println(key, value)
		key++
	}
	// user
	andres := User{"Andrés", "andresbillini@gmail.com", 33}
	fmt.Println(andres)

	sam := new(User)
	sam.Name = "Sam"
	sam.Email = "sam@lco.dev"
	sam.Age = 22
	fmt.Printf("%v\n", sam.Email)

	tobby := &User{"Tobby", "tobby@lco.dev", 29}
	fmt.Printf("%v\n", tobby.Age)

	// Functions
	Superman()
	mymultiply := Multiply(6, 4)
	myresult, mylength, myname := Adder(3, 6, 4, 2, 4, mymultiply)
	fmt.Println(myresult, mylength, myname)

	// Handling Error
	total, err := SumOf(2, 10)
	if err != nil {
		fmt.Println("Hubo un error:", err)
	} else {
		fmt.Println("El total de la suma es:", total)
	}

	// Maps
	maps()
}
