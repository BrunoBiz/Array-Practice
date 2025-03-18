package main

import (
	"fmt"

	product "example.com/array-practice/Product"
)

func main() {
	// 1
	fmt.Println("1")
	hobbies := [3]string{"Reading", "Training", "Playing"}
	fmt.Println(hobbies)

	// 2
	fmt.Println("2")
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3
	fmt.Println("3")
	hobbiesSliced := hobbies[0:2]
	//hobbiesSliced := hobbies[:2]
	fmt.Println(hobbiesSliced)

	// 4
	fmt.Println("4")
	hobbiesSliced = hobbies[1:]
	fmt.Println(hobbiesSliced)

	// 5
	fmt.Println("5")
	courseGoals := []string{"Learn", "New Job!?"}
	fmt.Println(courseGoals)

	// 6
	fmt.Println("6")
	courseGoals[1] = "Update my stack"
	fmt.Println(courseGoals)
	courseGoals = append(courseGoals, "Study")
	fmt.Println(courseGoals)

	// 7
	fmt.Println("7")
	product1, err := product.New("Apple", 1, 0.99)
	product2, err := product.New("Orange", 2, 1.99)
	product3, err := product.New("Banana", 3, 1.50)

	if err != nil {
		fmt.Println(err)
	} else {
		productList := []product.Product{product1, product2}
		fmt.Println(productList)

		productList = append(productList, product3)
		fmt.Println(productList)
	}
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
