package main

import "fmt"
// Comments
/*
	To run in Terminal:
	go run <filename>
	go run herewego.go
	
	To lookup documentation on a function
	godoc <package name> <Func name>
	godoc fmt Println

	variables in go are statically typed - values cannot change once defined
*/

func main() {
	fmt.Println("Hello World")

	var age int = 40

	var favNum float64 = 1.6180339


	fmt.Println(age, favNum) // will print out with space between

	var numOne = 1.000
	var num99 = .9999

	fmt.Println(numOne - num99)

	// + - * / % arithmetic works

	const pi float64 = 3.14159265

	var myName string = "Justice"

	fmt.Println(len(myName))
	fmt.Println(myName + " is a robot")
	fmt.Println("I like \n \nNewlines")

	// var isOver40 bool = false

	// var (
	// 	varA = 2,
	// 	varB = 3
	// )

	fmt.Printf("%f \n", pi) // print out float
	fmt.Printf("%.3f \n", pi) // print out float with certain precision
	fmt.Printf("%T \n", pi) // print out variable type
	fmt.Printf("%d \n", 100) // print out integer
	fmt.Printf("%b \n", 100) // print out binary
	fmt.Printf("%c \n", 100) // print out char code
	fmt.Printf("%x \n", 100) // print out hex
	fmt.Printf("%e \n", pi) // print out sci notation

	// &&, ||, ! logical operators

	// == != < > <= >= relational operators
	
	/*
	i := 1
	for i <= 3 {
		fmt.Println(i)

		i++

	}
	*/

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else if yourAge >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can have fun")	
	}

	switch yourAge {
		case 16 : fmt.Println("Go drive")
		case 18 : fmt.Println("Go vote")
		default : fmt.Println("Go have fun")
	}

	var favNums2[3] float64
	favNums2[0] = 1
	favNums2[1] = 2
	favNums2[2] = 5

	fmt.Println(favNums2[2])

	favNums3 := [5]float64 {1,2,4,6,7}

	for _, value := range favNums3 {
		fmt.Println(value)
	}
	// for i, value := range favNums3 {
	// 	fmt.Println(value, i)
	// }

	numSlice := []int {5,4,3,2,1}

	numSlice2 := numSlice[3:5] // get 3 and 4 and ignore 5

	fmt.Println("numSlice2[0] =", numSlice2[0])

	fmt.Println("numSlice[:2] =",  numSlice[:2]) //start from the beginning stop before 2
	fmt.Println("numSlice[2:] =",  numSlice[2:]) //start from the beginning stop before 2

	numSlice3 := make([]int, 5, 10) // type, default zero for the first 5, max size of 10

	fmt.Println("numSlice3[:] =",  numSlice3[:]) //start from the beginning stop before 2
	
	numSlice3 = append(numSlice3, 0 , -1)

	fmt.Println("numSlice3[:] =",  numSlice3[:]) //start from the beginning stop before 2

	// Maps

	presAge := make(map[string] int)
	presAge["TheodoreRoosevelt"] = 42
	presAge["John F. Kennedy"] = 43

	fmt.Println(len(presAge))
	
	delete(presAge, "John F. Kennedy")
	fmt.Println(len(presAge))
	
	listNums := []float64{1,2,3,4,5}
	
	fmt.Println("Sum :", addThemUp(listNums))

	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)
}

func addThemUp(numbers []float64) float64 { // (attributes) return type
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func next2Values(number int) (int, int) {
	return number+1, number+2
}
