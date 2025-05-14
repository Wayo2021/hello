package main

import "fmt"

// // import "rsc.io/quote"

// func main() {
//     fmt.Println("Hello, World!")
//     // fmt.Println(quote.Go())
// }

// // Hello returns a greeting for the named person.
// func Hello(name string) string {
//     // Return a greeting that embeds the name in a message.
//     message := fmt.Sprintf("Hi, %v. Welcome!", name)
//     return message
// }

//Declare Variables-----------------------------------------------------------------------------------------------------

// func main(){
//     var a string
//     var b int
//     var c bool

//     fmt.Println(a)
//     fmt.Println(b)
//     fmt.Println(c)
// }

// func main(){
//     var student1 string
//     student1 = "John"
//     fmt.Println(student1)
// }

// var a int
// var b int = 2
// var c = 3

// func main(){
//     a = 1
//     fmt.Println(a)
//     fmt.Println(b)
//     fmt.Println(c)
// }

//Declare Multiple Variables-----------------------------------------------------------------------------------------------------

// func main(){
//     var a, b, c, d int = 1, 3, 5 ,7

//     fmt.Println(a)
//     fmt.Println(b)
//     fmt.Println(c)
//     fmt.Println(d)
// }

// func main(){
//     var a, b = 6, "Hello"
//     c, d := 7, "World!"

//     fmt.Println(a)
//     fmt.Println(b)
//     fmt.Println(c)
//     fmt.Println(d)
// }

// func main(){
//     var(
//         a int
//         b int = 1
//         c string = "hello"
//     )

//     fmt.Println(a)
//     fmt.Println(b)
//     fmt.Println(c)
// }

//Constant-----------------------------------------------------------------------------------------------------

// const PI = 3.14

// func main(){
//     fmt.Println(PI)
// }

// const A = 3.14

// func main(){
//     fmt.Println(A)
// }

// const(
// 	A int = 1
// 	B = 3.14
// 	C = "Hi!"
// )
// func main(){
// 	fmt.Println(A)
// 	fmt.Println(B)
// 	fmt.Println(C)
// }

// Output Function-----------------------------------------------------------------------------------------------------

// func main(){
// 	var i, j string = "Hello", "World"

// 	fmt.Println(i)
// 	fmt.Println(j)
// }

// func main(){
// 	var i, j string = "Hello", "World"

// 	fmt.Print(i, "\n")
// 	fmt.Print(j, "\n")
// }

// func main(){
// 	var i, j string = "Hello", "World"

// 	fmt.Println(i, "\n",j)
// }

// func main(){
// 	var i, j string = "Hello", "World"

// 	fmt.Println(i, "", j)
// }

// func main(){
// 	var i, j = 10, 20

// 	fmt.Print(i, j)
// }

// func main(){
// 	var i, j string = "Hello", "World"

// 	fmt.Println(i,j)
// }

// func main() {
//   var i string = "Hello"
//   var j int = 15

//   fmt.Printf("i has value: %v and type: %T\n", i, i)
//   fmt.Printf("j has value: %v and type: %T", j, j)
// }

//Formatting Verbs-----------------------------------------------------------------------------------------------------

// func main() {
//   var i = 15.5
//   var txt = "Hello World!"

//   fmt.Printf("%v\n", i)
//   fmt.Printf("%#v\n", i)
//   fmt.Printf("%v%%\n", i)
//   fmt.Printf("%T\n", i)

//   fmt.Printf("%v\n", txt)
//   fmt.Printf("%#v\n", txt)
//   fmt.Printf("%T\n", txt)
// }

// func main() {
//   var i = 15

//   fmt.Printf("%b\n", i)
//   fmt.Printf("%d\n", i)
//   fmt.Printf("%+d\n", i)
//   fmt.Printf("%o\n", i)
//   fmt.Printf("%O\n", i)
//   fmt.Printf("%x\n", i)
//   fmt.Printf("%X\n", i)
//   fmt.Printf("%#x\n", i)
//   fmt.Printf("%4d\n", i)
//   fmt.Printf("%-4d\n", i)
//   fmt.Printf("%04d\n", i)
// }

// func main() {
//   var txt = "Hello"

//   fmt.Printf("%s\n", txt)
//   fmt.Printf("%q\n", txt)
//   fmt.Printf("%8s\n", txt)
//   fmt.Printf("%-8s\n", txt)
//   fmt.Printf("%x\n", txt)
//   fmt.Printf("% x\n", txt)
// }

// func main() {
//   var i = true
//   var j = false

//   fmt.Printf("%t\n", i)
//   fmt.Printf("%t\n", j)
// }

// func main() {
//   var i = 3.141

//   fmt.Printf("%e\n", i)
//   fmt.Printf("%f\n", i)
//   fmt.Printf("%.2f\n", i)
//   fmt.Printf("%6.2f\n", i)
//   fmt.Printf("%g\n", i)
// }

//Basic Data Types-----------------------------------------------------------------------------------------------------

// func main() {
// var a bool = true     // Boolean
// var b int = 5         // Integer
// var c float32 = 3.14  // Floating point number
// var d string = "Hi!"// String

// fmt.Println("Boolean: ", a)
// fmt.Println("Integer: ", b)
// fmt.Println("Float:   ", c)
// fmt.Println("String:  ", d)
// }

//Boolean-----------------------------------------------------------------------------------------------------

// func main() {
// var b1 bool = true // typed declaration with initial value
// var b2 = true // untyped declaration with initial value
// var b3 bool // typed declaration without initial value
// b4 := true // untyped declaration with initial value

// fmt.Println(b1) // Returns true
// fmt.Println(b2) // Returns true
// fmt.Println(b3) // Returns false
// fmt.Println(b4) // Returns true
// }

//Integer-----------------------------------------------------------------------------------------------------

// func main() {
// var x int = 500
// var y int = -4500
// fmt.Printf("Type: %T, value: %v", x, x)
// fmt.Printf("Type: %T, value: %v", y, y)
// }

// func main() {
// 	var x uint = 500
// 	var y uint = 4500
// 	fmt.Printf("Type: %T, value: %v", x, x)
// 	fmt.Printf("Type: %T, value: %v", y, y)
// }

// func main() {
// var x float32 = 123.78
// var y float32 = 3.4e+38
//   fmt.Printf("Type: %T, value: %v\n", x, x)
//   fmt.Printf("Type: %T, value: %v", y, y)
// }

// func main() {
// var x float64 = 1.7e+308
// fmt.Printf("Type: %T, value: %v", x, x)
// }

//String-----------------------------------------------------------------------------------------------------

// func main() {
// 	var txt1 string = "Hello!"
// 	var txt2 string
// 	txt3 := "World 1"

// 	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
// 	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
// 	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
// }

//Arrays-----------------------------------------------------------------------------------------------------

// func main() {
// 	var arr1 = [3]int{1, 2, 3}
// 	arr2 := [5]int{4, 5, 6, 7, 8}

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// }

// func main() {
// var arr1 = [...]int{1,2,3}
// arr2 := [...]int{4,5,6,7,8}

// fmt.Println(arr1)
// fmt.Println(arr2)
// }

// func main() {
//   var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
//   fmt.Print(cars)
// }

// func main() {
// prices := [3]int{10,20,30}

// fmt.Println(prices[0])
// fmt.Println(prices[2])
// }

// func main() {
// prices := [3]int{10,20,30}

// prices[2] = 50
// fmt.Println(prices)
// }

// func main() {
// arr1 := [5]int{} //not initialized
// arr2 := [5]int{1,2} //partially initialized
// arr3 := [5]int{1,2,3,4,5} //fully initialized

// fmt.Println(arr1)
// fmt.Println(arr2)
// fmt.Println(arr3)
// }

// func main() {
// arr1 := [5]int{1:10,2:40}

// fmt.Println(arr1)
// }

// func main() {
// arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
// arr2 := [...]int{1,2,3,4,5,6}

// fmt.Println(len(arr1))
// fmt.Println(len(arr2))
// }

// func main() {
//   myslice1 := []int{}
//   fmt.Println(len(myslice1))
//   fmt.Println(cap(myslice1))
//   fmt.Println(myslice1)

//   myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
//   fmt.Println(len(myslice2))
//   fmt.Println(cap(myslice2))
//   fmt.Println(myslice2)
// }

// func main() {
// arr1 := [6]int{10, 11, 12, 13, 14,15}
// myslice := arr1[2:4]

//   fmt.Printf("myslice = %v\n", myslice)
//   fmt.Printf("length = %d\n", len(myslice))
//   fmt.Printf("capacity = %d\n", cap(myslice))
// }

// func main() {
//   myslice1 := make([]int, 5, 10)
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   // with omitted capacity
//   myslice2 := make([]int, 5)
//   fmt.Printf("myslice2 = %v\n", myslice2)
//   fmt.Printf("length = %d\n", len(myslice2))
//   fmt.Printf("capacity = %d\n", cap(myslice2))
// }

//Modify Slice-----------------------------------------------------------------------------------------------------

// func main() {
// prices := []int{10,20,30}

// fmt.Println(prices[0])
// fmt.Println(prices[2])
// }

// func main() {
//   myslice1 := []int{1, 2, 3, 4, 5, 6}
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   myslice1 = append(myslice1, 20, 21)
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))
// }

// func main() {
// myslice1 := []int{1,2,3}
// myslice2 := []int{4,5,6}
//   myslice3 := append(myslice1, myslice2...)

//   fmt.Printf("myslice3=%v\n", myslice3)
//   fmt.Printf("length=%d\n", len(myslice3))
//   fmt.Printf("capacity=%d\n", cap(myslice3))
// }

// func main() {
//   arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
//   myslice1 := arr1[1:5] // Slice array
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   myslice1 = arr1[1:3] // Change length by re-slicing the array
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))
// }

// func main() {
//   numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
//   // Original slice
//   fmt.Printf("numbers = %v\n", numbers)
//   fmt.Printf("length = %d\n", len(numbers))
//   fmt.Printf("capacity = %d\n", cap(numbers))

//   // Create copy with only needed numbers
//   neededNumbers := numbers[:len(numbers)-10]
//   numbersCopy := make([]int, len(neededNumbers))
//   copy(numbersCopy, neededNumbers)

//   fmt.Printf("numbersCopy = %v\n", numbersCopy)
//   fmt.Printf("length = %d\n", len(numbersCopy))
//   fmt.Printf("capacity = %d\n", cap(numbersCopy))
// }

//Operator-----------------------------------------------------------------------------------------------------

// func main() {
// 	var (
// 		sum1 = 100 + 50    // 150 (100 + 50)
// 		sum2 = sum1 + 250  // 400 (150 + 250)
// 		sum3 = sum2 + sum2 // 800 (400 + 400)
// 	)
// 	fmt.Println(sum3)
// }

//Assignment Operators-----------------------------------------------------------------------------------------------------

// func main() {
// var x = 10
// fmt.Println(x)
// }

// func main() {
// var x = 10
// x +=5
// fmt.Println(x)
// }

//Comparison Operators-----------------------------------------------------------------------------------------------------

// func main() {
// var x = 5
// var y = 3
// fmt.Println(x>y) // returns 1 (true) because 5 is greater than 3
// }

// if statement-----------------------------------------------------------------------------------------------------

// func main(){
// 	if 20 > 18{
// 		fmt.Println("20 is greater than 18")
// 	}
// }

// func main() {
// 	x := 20
// 	y := 18
// 	if x > y {
// 		fmt.Println("x is greater than y")
// 	}
// }
// if else statement
// func main() {
// 	time := 20
// 	if time < 18 {
// 		fmt.Println("Good day.")
// 	} else {
// 		fmt.Println("Good evening.")
// 	}
// }

// func main() {
// 	temperature := 14
// 	if temperature > 15 {
// 		fmt.Println("It is warm out there")
// 	} else {
// 		fmt.Println("It is cold out there")
// 	}
// }

// else if statement-----------------------------------------------------------------------------------------------------

// func main() {
// 	time := 22
// 	if time < 10 {
// 		fmt.Println("Good morning.")
// 	} else if time < 20 {
// 		fmt.Println("Good day.")
// 	} else {
// 		fmt.Println("Good evening.")
// 	}
// }

// func main() {
// 	a := 14
// 	b := 14
// 	if a < b {
// 		fmt.Println("a is less than b.")
// 	} else if a > b {
// 		fmt.Println("a is more than b.")
// 	} else {
// 		fmt.Println("a and b are equal.")
// 	}
// }

// func main() {
// 	x := 30
// 	if x >= 10 {
// 		fmt.Println("x is larger than or equal to 10.")
// 	} else if x > 20 {
// 		fmt.Println("x is larger than 20.")
// 	} else {
// 		fmt.Println("x is less than 10.")
// 	}
// }

//Nested if statement-----------------------------------------------------------------------------------------------------

// func main() {
// 	num := 20
// 	if num >= 10 {
// 		fmt.Println("Num is more than 10.")
// 		if num > 15 {
// 			fmt.Println("Num is also more than 15.")
// 		}
// 	} else {
// 		fmt.Println("Num is less than 10.")
// 	}
// }

// switch statement-----------------------------------------------------------------------------------------------------

// func main() {
// 	day := 4

// 	switch day {
// 	case 1:
// 		fmt.Println("Monday")
// 	case 2:
// 		fmt.Println("Tuesday")
// 	case 3:
// 		fmt.Println("Wednesday")
// 	case 4:
// 		fmt.Println("Thursday")
// 	case 5:
// 		fmt.Println("Friday")
// 	case 6:
// 		fmt.Println("Saturday")
// 	case 7:
// 		fmt.Println("Sunday")
// 	}
// }

//multi-case switch statement-----------------------------------------------------------------------------------------------------

// func main() {
// 	day := 5

// 	switch day {
// 	case 1, 3, 5:
// 		fmt.Println("Odd weekday")
// 	case 2, 4:
// 		fmt.Println("Even weekday")
// 	case 6, 7:
// 		fmt.Println("Weekend")
// 	default:
// 		fmt.Println("Invalid day of day number")
// 	}
// }

//For loops-----------------------------------------------------------------------------------------------------

// func main() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// }

// func main(){
// 	for i:= 0; i <= 100; i+=10{
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		if i == 3 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		if i == 3 {
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	adj := [2]string{"big", "tasty"}
// 	fruits := [3]string{"apple", "orange", "banana"}
// 	for i := 0; i < len(adj); i++ {
// 		for j := 0; j < len(fruits); j++ {
// 			fmt.Println(adj[i], fruits[j])
// 		}
// 	}
// }

// func main() {
// 	fruits := [3]string{"apple", "orange", "banana"}
// 	for idx, val := range fruits {
// 		fmt.Printf("%v\t%v\n", idx, val)
// 	}
// }

// func main() {
// 	fruits := [3]string{"apple", "orange", "banana"}
// 	for _, val := range fruits {
// 		fmt.Printf("%v\n", val)
// 	}
// }

// func main() {
// 	fruits := [3]string{"apple", "orange", "banana"}

// 	for idx, _ := range fruits {
// 		fmt.Printf("%v\n", idx)
// 	}
// }

// Function-----------------------------------------------------------------------------------------------------

// func myMessage() {
// 	fmt.Println("I just got executed!")
// }

// func main() {
// 	myMessage() // call the function
// }

// func myMessage() {
// 	fmt.Println("I just got executed!")
// }

// func main() {
// 	myMessage()
// 	myMessage()
// 	myMessage()
// }

// func familyName(fname string) {
// 	fmt.Println("Hello", fname, "Refsnes")
// }

// func main() {
// 	familyName("Liam")
// 	familyName("Jenny")
// 	familyName("Anja")
// }

// func familyName(fname string, no int) {
// 	fmt.Println("Hello", fname, "Refsnes", "No", no)
// }

// func main() {
// 	familyName("Liam", 1)
// 	familyName("Jenny", 2)
// 	familyName("Anja", 3)
// }

// Function Returns-----------------------------------------------------------------------------------------------------

// func myFunction(x int, y int) int {
// 	return x + y
// }
// func myFunction2(x int, y int) int {
// 	return x % y
// }

// func main() {
// 	fmt.Println(myFunction(2, 10))
// 	fmt.Println(myFunction2(10, 2))
// }

// func myFunction(x int, y int) (result int) {
// 	result = x + y
// 	return
// }

// func main() {
// 	fmt.Println(myFunction(1, 2))
// }

// func myFunction(x int, y int) (result int) {
// 	result = x + y
// 	return result
// }

// func main() {
// 	fmt.Println(myFunction(1, 2))
// }

// func myFunction(x int, y int) (result int) {
// 	result = x + y
// 	return
// }

// func main() {
// 	total := myFunction(1, 2)
// 	fmt.Println(total)
// }

// func myFunction(x int, y string) (result int, txt1 string) {
// 	result = x + x
// 	txt1 = y + " World!"
// 	return
// }

// func main() {
// 	fmt.Println(myFunction(5, "Hello"))
// }

// func myFunction(x int, y string) (result int, txt1 string) {
// 	result = x + x
// 	txt1 = y + " World!"
// 	return
// }

// func main() {
// 	a, b := myFunction(5, "Hello")
// 	fmt.Println(a, b)
// 	fmt.Println(a)
// }

// func myFunction(x int, y string) (result int, txt1 string) {
// 	result = x + x
// 	txt1 = y + " World!"
// 	return
// }

// func main() {
//    _, b := myFunction(5, "Hello")
// 	fmt.Println(b)
// }

//Recursion Functions-----------------------------------------------------------------------------------------------------

// func testcount(x int) int {
//   if x == 11 {
//     return 0
//   }
//   fmt.Println(x)
//   return testcount(x + 1)
// }

// func main(){
//   testcount(1)
// }

// func factorial_recursion(x float64) (y float64) {
// 	if x > 0 {
//      y = x * factorial_recursion(x-1)
// 	} else {
//      y = 1
// 	}
// 	return
// }

// func main() {
// 	fmt.Println(factorial_recursion(4))
// }

//struct-----------------------------------------------------------------------------------------------------

// type Person struct {
//   name string
//   age int
//   job string
//   salary int
// }

// func main() {
//   var pers1 Person
//   var pers2 Person

//   // Pers1 specification
//   pers1.name = "Hege"
//   pers1.age = 45
//   pers1.job = "Teacher"
//   pers1.salary = 6000

//   // Pers2 specification
//   pers2.name = "Cecilie"
//   pers2.age = 24
//   pers2.job = "Marketing"
//   pers2.salary = 4500

//   // Access and print Pers1 info
//   fmt.Println("Name: ", pers1.name)
//   fmt.Println("Age: ", pers1.age)
//   fmt.Println("Job: ", pers1.job)
//   fmt.Println("Salary: ", pers1.salary)

//   // Access and print Pers2 info
//   fmt.Println("Name: ", pers2.name)
//   fmt.Println("Age: ", pers2.age)
//   fmt.Println("Job: ", pers2.job)
//   fmt.Println("Salary: ", pers2.salary)
// }

// type Person struct {
//   name string
//   age int
//   job string
//   salary int
// }

// func main() {
//   var pers1 Person
//   var pers2 Person

//   // Pers1 specification
//   pers1.name = "Hege"
//   pers1.age = 45
//   pers1.job = "Teacher"
//   pers1.salary = 6000

//   // Pers2 specification
//   pers2.name = "Cecilie"
//   pers2.age = 24
//   pers2.job = "Marketing"
//   pers2.salary = 4500

//   // Print Pers1 info by calling a function
//   printPerson(pers1)

//   // Print Pers2 info by calling a function
//   printPerson(pers2)
// }

// func printPerson(pers Person) {
//   fmt.Println("Name: ", pers.name)
//   fmt.Println("Age: ", pers.age)
//   fmt.Println("Job: ", pers.job)
//   fmt.Println("Salary: ", pers.salary)
// }

//Maps-----------------------------------------------------------------------------------------------------

// func main() {
// 	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
// 	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

// 	fmt.Printf("a\t%v\n", a)
// 	fmt.Printf("b\t%v\n", b)
// }

// func main() {
// 	var a = make(map[string]string) // The map is empty now
// 	a["brand"] = "Ford"
// 	a["model"] = "Mustang"
// 	a["year"] = "1964"
// 	// a is no longer empty
// 	b := make(map[string]int)
// 	b["Oslo"] = 1
// 	b["Bergen"] = 2
// 	b["Trondheim"] = 3
// 	b["Stavanger"] = 4

// 	fmt.Printf("a\t%v\n", a)
// 	fmt.Printf("b\t%v\n", b)
// }

// func main() {
// 	var a = make(map[string]string)
// 	var b map[string]string

// 	fmt.Println(a == nil)
// 	fmt.Println(b == nil)
// }

// func main() {
// 	var a = make(map[string]string)
// 	a["brand"] = "Ford"
// 	a["model"] = "Mustang"
// 	a["year"] = "1964"

// 	fmt.Println(a)

// 	a["year"] = "1970" // Updating an element
// 	a["color"] = "red" // Adding an element

// 	fmt.Println(a)
// }

// func main() {
// 	var a = make(map[string]string)
// 	a["brand"] = "Ford"
// 	a["model"] = "Mustang"
// 	a["year"] = "1964"

// 	fmt.Println(a)

// 	delete(a,"year")

// 	fmt.Println(a)
// }

// func main() {
// 	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

// 	val1, ok1 := a["brand"]// Checking for existing key and its value
// 	val2, ok2 := a["color"]// Checking for non-existing key and its value
// 	val3, ok3 := a["day"]	// Checking for existing key and its value
// 	_, ok4 := a["model"]	// Only checking for existing key and not its value

// 	fmt.Println(val1, ok1)
// 	fmt.Println(val2, ok2)
// 	fmt.Println(val3, ok3)
// 	fmt.Println(ok4)
// }

// func main() {
// 	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
// 	b := a

// 	fmt.Println(a)
// 	fmt.Println(b)

// 	b["year"] = "1970"
// 	fmt.Println("After change to b:")

// 	fmt.Println(a)
// 	fmt.Println(b)
// }

func main() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}
}

// func main() {
// 	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

// 	var b = []string	// defining the order
// 	b = append(b, "one", "two", "three", "four")

// 	for k, v := range a {	// loop with no order
// 	fmt.Printf("%v : %v, ", k, v)
// 	}

// 	fmt.Println()

// 	for _, element := range b {	// loop with the defined order
// 	fmt.Printf("%v : %v, ", element, a[element])
// 	}
// }
