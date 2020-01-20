package main

import (
	"errors"
	"fmt"
	"math"
)
//fmt is related to IO operations.
//No commas in list

func main() {
	//This isn't guaranteed in future releases - println("")

	//Declaring variables - less neat
	/*var x int = 34

	var y string = "Hello from cabin 3"

	var z int = 2

	var sum = x + z

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("The sentence is concatenated with value", sum, "as the sum")*/

	//Neater version of that up there
	x := 34
	y := "Hello from cabin 3! "
	z := 2
	add := x + z

	fmt.Println(y)
	fmt.Println("The sum of ", x , " and ", z, " is ", add )

	//If statement - no parentheses around condition
	a := 7

	if a > 234 {
	} else if a != 7 {		fmt.Println(a, " is more than value provided")

		fmt.Println(a, " is equal")
	} else {
		fmt.Println(a, " is less than value provided")
	}

	//Arrays - Same type elements, fixed no of elems(length)
	var array2 [4]int
	array2[0] = 1
	array2[2] = 3

	fmt.Println("Array 2:", array2, " Array length is:", len(array2))

	//Better way to initialize Array values
	array1 := [4]string{ "John", "Doe", "Jane", "Dale"}
	fmt.Println("Array 1", array1)

	//You can't add to a fixed bound array size - Use slices to overcome this
	//When using a slice a new variable is created

	sliceArray := []int{5,6,7,8}
	sliceArray = append(sliceArray, 1,2,3,4)

	fmt.Println("Slice array did not modify the contents of the array, it only appended to it", sliceArray)
	fmt.Println(sliceArray[7])

	//Maps (Like Dictionaries in Python or Associative arrays in PHP)
	vertices := make(map[string]int)

	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["dodecagon"] = 12

	delete(vertices, "triangle")
	fmt.Println(vertices["triangle"], vertices["dodecagon"], vertices["square"])

	//Loops - For is the only type of loop in Go
	for w := 1; w <= 10; w++ {
		fmt.Println("Pure for loop: ", w)
	}

	//For loop also doubles as a while loop as follows
	r := 0
	for r < 10 {
		fmt.Println("Testing while loop with for:", r)
		r++
	}


	//Iterate over each element in an array or slice by using range
	languages := []string{"Go", "Scala", "Rust", "Python"}

	for index, value := range languages {
		fmt.Println("index:", index, "value:", value)
	}

	//Similar thing for map
	studentCourses := map[string]string { //This is map literal
		"alex" : "CS",
		"dave" : "CT",
		"ed" : "EEE",
		"pete": "HR",
	}
	for index, value := range studentCourses {
		fmt.Println("key: ",index,"value: ",value)
	}
	fmt.Println(studentCourses["alex"])

	//Call functions
	result := sum (3,5)
	println("Sum is: ",result)

	result2 := concat("dav", "is")
	println("Concatenating two strings: ",result2)

	sqrtResult, err := sqrt(121)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrtResult)
	}
	//fmt.Println(sqrt(144))

	//create STRUCT of a type
	p := person{
		name:"Daniel",
		age:2,
	}
	fmt.Println(p)

	//Pointers
	i := 7
	inc(&i)
	fmt.Println(i)

	strnll := "Hilove"
	fmt.Println(strnll)
	strn(&strnll)

	strnlll := "DSD"
	fmt.Println(strn2(strnlll))

}
//Pointers
func inc (x *int) {
	*x++ //The asterisk here derefences the value at the memory address and increments it
		 //Without it you're incrementing the memory address
}
func strn (t *string) {
	println(*t)
}

func strn2 (dr string) string {
	return dr
}

//Creating functions: aside from main
func sum(a int, b int) int {
	return a + b
}
func concat(a string,  b string) string {
	return a + b
}


//Multiple return values in Go
//Very useful when throwing errors because Go doesn't have Exceptions like in Java
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

//Struct - collection of fields; so you can group things together to create a more logical type
//Created outside a function
//Similar to case classes in Scala
type person struct {
	name string
	age int
}

//Pointers

