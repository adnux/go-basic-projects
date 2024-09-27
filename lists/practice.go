package lists

import (
	"fmt"

	"github.com/adnux/go-basic-projects/colors"
)

func printWithColor(label string, data interface{}) {
	fmt.Println(colors.Yellow, label, colors.Reset, data)
}

func StartListsPractice() {
	// 1) Create a new array (!) that contains three hobbies you have
	hobbies := [3]string{"Reading", "Coding", "Gaming"}
	// 		Output (print) that array in the command line.
	printWithColor("Hobbies=>", hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	printWithColor("First Hobby=>", hobbies[0])
	//		- The second and third element combined as a new list
	printWithColor("Second and Third Hobby=>", hobbies[1:3])

	// 3) Create a slice based on the first element that contains
	mainHobbies := hobbies[:2]
	//		the first and second elements.
	printWithColor("Hobby Slice 1=>", mainHobbies)
	mainHobbies = hobbies[0:2]
	//		Create that slice in two different ways (i.e. create two slices in the end)
	printWithColor("Hobby Slice 2=>", mainHobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second
	printWithColor("Capacity=>", cap(mainHobbies))
	mainHobbies = hobbies[1:3]
	//		and last element of the original array.
	printWithColor("Hobby Slice 3=>", mainHobbies)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Learn Go", "Build a project"}
	printWithColor("Course Goals=>", courseGoals)

	// 6) Set the second goal to a different one
	courseGoals[1] = "Build a web app"
	// AND then add a third goal to that existing dynamic array
	courseGoals = append(courseGoals, "Learn Docker")
	printWithColor("Course Goals=>", courseGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	type Product struct {
		ID    int
		Title string
		Price float64
	}
	//		dynamic list of products (at least 2 products).
	products := []Product{
		{1, "Book", 10.99},
		{2, "Laptop", 999.99},
	}
	printWithColor("Products=>", products)
	//		Then add a third product to the existing list of products.
	products = append(products, Product{3, "Headphones", 99.99})
	printWithColor("Products=>", products)
}
