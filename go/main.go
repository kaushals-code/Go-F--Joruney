package main

import (
	"fmt"
)

// ===================================== DAY 5 ====================================================

func main() {
	fmt.Println("The Weeknd is the GOAT bro")

	// numbers := []int{1, 2, 3, 4, 5}
	// var double []int

	// for _, x := range numbers {
	// 	double = append(double, 2*x)
	// }

	// fmt.Println(double)

	// // Actual Filter implementation
	// numbers := []int{1, 2, 3, 4, 5}
	// even := []int{}

	// for _, x := range numbers {
	// 	if x%2 == 0 {
	// 		even = append(even, x)
	// 	}
	// }

	// fmt.Println(even)

	// 	numbers := []int{1, 2, 3, 4, 5}
	// 	// even := []int{}
	// 	var even []int

	// 	for _, x := range numbers {
	// 		if x%2 == 0 {
	// 			even = append(even, x)
	// 		}
	// 	}

	// 	fmt.Println(even)

	// String are immutable in go!
	// str := "theweeknd"
	// for _, x := range str {
	// 	x = 'a'
	// }

	// fmt.Printf(str)

	// nums := []int{4, 2, 6, 3, 1}
	// slices.Sort(nums)

	// fmt.Println(nums)

	// fmt.Println(slices.Contains(nums, 5))

	// students := []Student{
	// 	{"Alice", 90, 85, 88},
	// 	{"Bob", 75, 80, 72},
	// 	{"Charlie", 95, 92, 96},
	// 	{"Diana", 82, 88, 84},
	// 	{"Ethan", 68, 70, 75},
	// 	{"Fiona", 91, 89, 93},
	// 	{"George", 79, 83, 80},
	// 	{"Hannah", 88, 85, 90},
	// 	{"Ian", 70, 65, 72},
	// 	{"Julia", 96, 94, 98},
	// 	{"Kevin", 84, 82, 86},
	// 	{"Laura", 77, 79, 81},
	// 	{"Michael", 89, 91, 88},
	// 	{"Nora", 83, 80, 85},
	// 	{"Oscar", 92, 95, 90},
	// 	{"Paula", 74, 78, 76},
	// 	{"Quinn", 86, 84, 89},
	// 	{"Rachel", 90, 88, 92},
	// 	{"Sam", 78, 75, 80},
	// 	{"Tina", 94, 92, 95},
	// }

	// slices.SortFunc(students, func(a, b Student) int {
	// 	// return b.marks - a.marks
	// 	avg1 := (a.math + a.physics + a.chemistry) / 3
	// 	avg2 := (b.math + b.physics + b.chemistry) / 3
	// 	return avg2 - avg1
	// })

	// fmt.Println(students)
}

// type Student struct {
// 	name      string
// 	math      int
// 	physics   int
// 	chemistry int
// }

// ===================================== DAY 4 ====================================================

// func total(n ...int) int {
// 	ans := 0
// 	for _, value := range n {
// 		ans += value
// 	}
// 	return ans
// }

// func main() {
// 	fmt.Println("The Weeknd is the GOAT")

// array := []int{1, 2, 3}

// newArray := append(array, 4)

// fmt.Println(array)
// fmt.Println(newArray)

// numbers := []int{1, 2, 3}

// fmt.Println(len(numbers))

// a := []int{1, 2}
// b := []int{3, 4}

// c := append(a, b...)

// fmt.Println(c)

// numbers := []int{10, 20, 30, 40}

// part := numbers[1:3]

// fmt.Println(part)

// revise the make() function once again

// org := []int{10, 20, 30}

// cp := make([]int, len(org))

// copy(cp, org)

// cp[0] = 100

// fmt.Println(cp)
// fmt.Println(org)

// s := []int{}
// fmt.Println(s == nil)

// var t []int
// fmt.Println(t == nil)

// mat := [][]int{
// 	{1, 2, 3},
// 	{4, 5, 6},
// 	{7, 8, 9},
// }

// fmt.Println(mat)

// array := []int{10, 20, 30, 40, 50, 60}
// res := total(array...)
// fmt.Println(res)
// }

// ====================================== SOME EXTRA PROBLEMS =====================================

// func multiplicationTable(n int) {
// 	for i := 1; i <= 20; i++ {
// 		fmt.Printf("%d x %d = %d\n", n, i, (17 * i))
// 	}
// }

// func factorial(n int) int {
// 	ans := 1
// 	for i := 2; i <= n; i++ {
// 		ans *= i
// 	}
// 	return ans
// }

// func fibo(n int) {
// 	if n == 1 {
// 		fmt.Println("1")
// 		return
// 	} else if n == 2 {
// 		fmt.Println("1 1")
// 		return
// 	} else {
// 		fmt.Println("1 1 ")
// 		a := 1
// 		b := 1
// 		fmt.Printf("\b")
// 		for i := 0; i < (n - 2); i++ {
// 			c := a + b
// 			fmt.Printf("%d ", c)
// 			a = b
// 			b = c
// 		}
// 	}
// }

// func isPrime(n int) bool {
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func sayHello(name string) {
// 	fmt.Println(name + " is the GOAT from function")
// }

// // func add(a, b int) int {
// // 	fmt.Println(a + b)
// // 	return a + b
// // }

// // not a good practice to use in the modern systems
// func add(a, b int) (result int) {
// 	result = a + b
// 	return
// }

// // multiple return values
// func arithemetic(a, b int) (sum int, diff int, err error) {

// 	if a < b {
// 		return 0, 0, fmt.Errorf("a < b bro")
// 	}

// 	sum = a + b
// 	diff = a - b
// 	err = nil
// 	return

// }

// // return with error
// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("Cannot divide by 0")
// 	}

// 	return a / b, nil
// }

// // take any number of input as parameters
// func giveTotal(numbers ...int) int {
// 	total := 0

// 	for _, num := range numbers {
// 		total += num
// 	}

// 	return total
// }

// // closures examples
// func counter() func() int {
// 	count := 0

// 	return func() int {
// 		count++
// 		return count
// 	}
// }

// func multiplier(x int) func(int) int {
// 	return func(y int) int {
// 		return x * y
// 	}
// }

// // defer example
// func deferex() {

// 	// defer is LIFO principle based STACK
// 	defer fmt.Println("Goodbye")

// 	fmt.Println("Hello")
// }

// func main() {
// 	// ==================================== DAY 3 ============================================
// 	// fmt.Println("The Weeknd is the GOAT")

// 	// sayHello("The Weeknd")
// 	// sayHello("The Weeknd")
// 	// sayHello("The Weeknd")

// 	// add(2, 3)
// 	// add(4, 5)
// 	// add(8, 9)

// 	// ans, err := divide(10, 5)
// 	// fmt.Printf("%f, %s", ans, err)

// 	// add(2, 3)
// 	// add(4, 5)
// 	// add(8, 9)

// 	// a, b, c := arithemetic(100, 20)

// 	// fmt.Println(a)
// 	// fmt.Println(b)
// 	// // fmt.Println(c)

// 	// if c == nil {
// 	// 	fmt.Println("There is no error")
// 	// } else {
// 	// 	fmt.Println("The is a Error mf")
// 	// }

// 	// ans := giveTotal(10, 20, 30, 2, 4, 55, 6, 3, 33, 2, 3, 77, 7, 99)

// 	// fmt.Println(ans)

// 	// numbers := []int{1, 2, 3, 4, 5}

// 	// res := giveTotal(numbers...)
// 	// fmt.Println(res)

// 	// func() {
// 	// 	fmt.Println("Hello")
// 	// }()

// 	// next := counter()()
// 	// fmt.Println(next())
// 	// fmt.Println(next())
// 	// fmt.Println(next())
// 	// fmt.Println(next())

// 	// fmt.Println(next)
// 	// fmt.Println(next)
// 	// fmt.Println(next)
// 	// fmt.Println(next)

// 	// forx := multiplier(5)
// 	// fory := multiplier(6)

// 	// fmt.Println(forx(5))
// 	// fmt.Println(fory(6))

// 	deferex()

// }

// func main() {

// 	// ==================================== DAY 1 ============================================

// 	// var name string
// 	// var age int
// 	// var height float64
// 	// var isArtist bool

// 	// var name string = "The Weeknd"
// 	// age := 50
// 	// height := 6.12
// 	// isArtist := true

// 	// // fmt.Println(name)
// 	// // fmt.Println(age)
// 	// // fmt.Println(height)
// 	// // fmt.Println(isArtist)

// 	// fmt.Printf("The name is %q\n", name)
// 	// fmt.Printf("The number is %d\n", age)
// 	// fmt.Printf("The double is %f\n", height)
// 	// fmt.Printf("The boolean is %t\n", isArtist)

// 	///////////////////// ERROR HERE

// 	// const name = "The Weeknd"
// 	// name = "Kaushal Singh Thakur"
// 	// fmt.Println(name)

// 	//////////////////// IOTA
// 	// const (
// 	// 	Pending = iota
// 	// 	Active  = iota * 10
// 	// 	Suspended
// 	// 	Deleted
// 	// )

// 	// // status := Active

// 	// fmt.Println(Pending)
// 	// fmt.Println(Active)
// 	// fmt.Println(Suspended)
// 	// fmt.Println(Deleted)

// 	// fmt.Printf("======================================\n")
// 	// fmt.Printf("\t\tProfile\n")
// 	// fmt.Printf("======================================\n")
// 	// fmt.Printf("Name: \t\t %s\n", name)
// 	// fmt.Printf("Age: \t\t %d\n", age)
// 	// fmt.Printf("Height: \t %.2f\n", height)
// 	// fmt.Printf("isArtist: \t %t\n", isArtist)
// 	// fmt.Print("======================================\n")

// 	//============================================ DAY 2====================================

// 	// name := "kaushalsingh"

// 	// if name == "theweeknd" {
// 	// 	fmt.Printf("%s is the GOAT bro", name)
// 	// } else if name == "kaushalsingh" {
// 	// 	fmt.Printf("%s is the GOAT fan bro", name)
// 	// } else {
// 	// 	fmt.Printf("%s is the nood bro", name)
// 	// }

// 	// if name := 20; name >= 18 {
// 	// 	fmt.Println("Adult")
// 	// } else {
// 	// 	fmt.Println("Noob")
// 	// }

// 	////////////////// switch statements
// 	// day := 3
// 	// switch day {
// 	// case 1:
// 	// 	fmt.Println("Day one")
// 	// case 2:
// 	// 	fmt.Println("Day two")
// 	// default:
// 	// 	fmt.Println("No Day")
// 	// }

// 	////////////////// for loop
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println(i)
// 	// }

// 	////////////////// while loop
// 	// i := 0
// 	// for i <= 10 {
// 	// 	fmt.Println(i)
// 	// 	i++
// 	// }

// 	////////////////// range

// 	// numbers := []int{10, 20, 30, 40, 50}
// 	// for index, value := range numbers {
// 	// 	fmt.Printf("%d -> %d\n", index, value)
// 	// }

// 	////////////////// lables keyword

// 	// outer:
// 	// 	for i := 0; i < 10; i++ {
// 	// 		for j := 0; j < 10; j++ {
// 	// 			fmt.Println(i*10 + j)
// 	// 			if i*10+j == 60 {
// 	// 				break outer
// 	// 			}
// 	// 		}
// 	// 	}

// 	// }

// 	// prime number program
// 	// 	num := 31
// 	// 	shouldprint := true
// 	// outer:
// 	// 	for i := 2; i < num-1; i++ {
// 	// 		if num%i == 0 {
// 	// 			fmt.Println("The number is not prime")
// 	// 			shouldprint = false
// 	// 			break outer
// 	// 		} else {
// 	// 			shouldprint = true
// 	// 		}
// 	// 	}

// 	// 	if shouldprint == true {
// 	// 		fmt.Println("The given number is prime")
// 	// 	}
// 	// var a int
// 	// var b float64
// 	// var c string
// 	// var d bool

// 	// fmt.Printf("%d %.1f %q %t\n", a, b, c, d)

// 	// var x int = 10
// 	// var y string = strconv.Itoa(x)

// 	// fmt.Println(y)
// 	// fmt.Printf("%T\n", y)

// 	name := "3.14"

// 	ft, err := strconv.ParseFloat(name, 64)

// 	fmt.Println(ft)
// 	fmt.Println(err)

// }
