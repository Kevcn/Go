package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	sayHello()
	IfStatement()
	arrayDemo()
	mapDemo()
	loops()
	result := funcWithReturnType(9, 10)
	fmt.Println(result)

	// Error handling -> pass in a number greater than 10 to perform the function
	output, err := returnMultipleType(6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}

	p := person{name: "Jake", age: 26}
	fmt.Println("POCO demo ->", p)

	pointer()

}

func sayHello() {
	fmt.Println("Say Hello!")
}

func IfStatement() {
	fmt.Println("------ If Demo ------")

	x := 5

	if x > 3 {
		fmt.Println("Condition met")
	}
}

func arrayDemo() {
	fmt.Println("------ Array Demo ------")

	a := [5]int{5, 2, 3, 2, 1}
	a[2] = 777

	fmt.Println(a)

	// example of a slice, no fixed length like a List in c#
	b := []int{2, 6, 5}
	b = append(b, 888)

	fmt.Println(b)
}

func mapDemo() {
	fmt.Println("------ Map Demo ------")

	// equivalent to directionary, key value pair collection
	aMap := make(map[string]int)

	aMap["OJ"] = 2
	aMap["Sharp"] = 3
	aMap["CCS"] = 4

	delete(aMap, "Sharp")
	fmt.Println(aMap["OJ"])
}

func loops() {
	fmt.Println("------ For Loop ------")

	// The only type of loop is the for loop in GO!
	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}
	fmt.Println("------ While Loop ------")

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	fmt.Println("------ For Range ------")

	// loop through an array
	arr := []string{"abc", "def"}
	// _ if we dont care about the indedx
	for _, v := range arr {
		fmt.Println("Value:", v)
	}

	for index, value := range arr {
		fmt.Println("Index", index, "Value:", value)
	}

	// loop through a map, we get key-value instead of index-value
	aMap := make(map[string]int)

	aMap["One"] = 1
	aMap["Two"] = 2

	for key, value := range aMap {
		fmt.Println("Key", key, "Value:", value)
	}
}

func funcWithReturnType(a int, b int) int {
	return a + b
}

func returnMultipleType(x float64) (float64, error) {

	// Go dont have exception, this is how errors are handled
	if x > 10 {
		return math.Sqrt(x), nil
	}

	return 0, errors.New("condition not met, throwing an error")
}

// Basically a POCO
type person struct {
	name string
	age  int
}

func pointer() {
	i := 7
	// get the pointer

	passByReference(&i)

	fmt.Println(i)

}

func passByReference(x *int) {
	*x++
}
