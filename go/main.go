package main

// =================================== DAY 11 ======================================================

// type Shape interface {
// 	Area() float64
// }

// type Circle struct {
// 	Radius int
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * float64(c.Radius*c.Radius)
// }

// func main() {
// 	fmt.Println("The Weeknd is the GOAT bro")

// 	var s Shape
// 	s = Circle{
// 		Radius: 5,
// 	}

// 	// fmt.Println(c.Area()) // works perfectly fine

// 	switch s.(type) {
// 	case Circle:
// 		fmt.Println(s.Area())
// 	default:
// 		fmt.Println("No such shape")
// 	}

// 	var v any = 10

// 	switch a := v.(type) {
// 	case int:
// 		fmt.Println(a + 10)
// 	case string:
// 		fmt.Println(a + " Hello")
// 	default:
// 		fmt.Println("The type is not recognized")
// 	}
// }

// =================================== DAY 10 ======================================================

// interfact
// type Car interface {
// 	Revv() string
// 	Race() string
// }

// type Nexon struct {
// 	Name string
// }

// func (n Nexon) Revv() string {
// 	return "Nexon is Revving"
// }

// func (n Nexon) Race() string {
// 	return "Let's race bro"
// }

// type Shape interface {
// 	Area() float64
// }

// type Circle struct {
// 	Radius int
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * float64(c.Radius) * float64(c.Radius)
// }

// type Rectangle struct {
// 	Height int
// 	Width  int
// }

// func (r Rectangle) Area() float64 {
// 	return float64(r.Height) * float64(r.Width)
// }

// type Square struct {
// 	Side int
// }

// // accept interfact and return member returned value
// func giveArea(s Shape) float64 {
// 	return s.Area()
// }

// func giveCircle(radius int) Circle {
// 	return Circle{
// 		Radius: radius,
// 	}
// }

// // type switch
// func typeswitch(x any) {
// 	switch v := x.(type) {
// 	case int:
// 		fmt.Println("integer")

// 	case string:
// 		fmt.Printf("string %s", v)

// 	default:
// 		fmt.Println("Unknown type")
// 	}
// }

// func getShape(s any) {
// 	switch v := s.(type) {
// 	case Circle:
// 		fmt.Println(v.Radius)
// 	case Rectangle:
// 		fmt.Println(v.Height, v.Width)
// 	}
// }

// type Speaker interface {
// 	Speak() string
// }

// type Dog struct {
// 	Name string
// }

// func (d *Dog) Speak() string {
// 	d.Name = "Mr. " + d.Name
// 	return "done"
// }

// func main() {
// 	fmt.Println("The Weeknd is the GOAT bro")

// var c Car

// c = Nexon{Name: "Tata Nexon"}

// fmt.Println(c.Revv())

// var s Shape

// s = Circle{
// 	Radius: 5,
// }

// var sr Shape

// sr = Rectangle{
// 	Height: 20,
// 	Width:  15,
// }

// // fmt.Println(s.Area())
// // fmt.Println(sr.Area())
// fmt.Println(giveArea(s))
// fmt.Println(giveArea(sr))

// var s Shape

// s = giveCircle(5)
// fmt.Println(giveArea(s))

// var x any = "kaushal singh thakur"

// value, ok := x.(int)

// fmt.Println(value)
// fmt.Println(ok)

// if v, ok := x.(int); ok {
// 	fmt.Println("int")
// 	fmt.Println(v)
// } else if v, ok := x.(string); ok {
// 	fmt.Println("string")
// 	fmt.Println(v)
// }

// check whethere the struct satisfies the interface
// var _ Shape = Square() // if not then this line produces an error

// 	var s Speaker
// 	s = &Dog{
// 		Name: "Bruno",
// 	}

// 	tw, ok := s.(*Dog)
// 	if ok {
// 		fmt.Println(tw.Name)
// 	}

// 	fmt.Println(s.Speak())

// }

// ===================================== DAY 9 ===================================================

// type User struct {
// 	name string
// 	age  int
// }

// // constructor function for go
// func newUser(name string, age int) User {
// 	return User{
// 		name: name,
// 		age:  age,
// 	}
// }

// func newUserPrt(name string, age int) *User {
// 	return &User{
// 		name: name,
// 		age:  age,
// 	}
// }

// // constructor with validation
// func validNewUser(name string, age int) (*User, error) {
// 	if age < 18 {
// 		return nil, fmt.Errorf("This mf is less that 18 dawg")
// 	}

// 	return &User{
// 		name: name,
// 		age:  age,
// 	}, nil
// }

// func (u *User) setage(n int) {
// 	u.age = n
// }

// type Celcius float64

// func (c Celcius) toCelcius() float64 {
// 	return float64(c)*9/5 + 32
// }

// func main() {
// 	fmt.Println("The Weeknd is the GOAT bro")

// user := User{
// 	name: "theweeknd",
// 	age:  36,
// }

// fmt.Println(user.name)

// user.setage(67)
// fmt.Println(user.age)

// user, err := validNewUser("theweeknd", 36)

// if err == nil {
// 	fmt.Println(user.name)
// } else {
// 	fmt.Println(err.Error())
// }

// 	celcius := Celcius(30)

// 	fmt.Println(celcius.toCelcius())

// }

// ===================================== DAY 8 ===================================================

// type Song struct {
// 	name   string
// 	album  string
// 	artist string
// }

// nested struct example
// type Address struct {
// 	houseno string
// 	area    string
// }

// type User struct {
// 	name    string
// 	email   string
// 	phone   string
// 	address Address
// }

// func altermarks(student *Student) {
// 	student.marks += 50
// }

// type Student struct {
// 	name  string
// 	marks int
// }

// for embedded structs
// type Where struct {
// 	hno string
// }

// type User struct {
// 	name string
// 	Where
// }

// func main() {
// wickedgames := Song{
// 	name:   "Wicked Games",
// 	album:  "Trilogy",
// 	artist: "The Weeknd",
// }

// fmt.Println(wickedgames.name)
// // fmt.Println(wickedgames.album)
// // fmt.Println(wickedgames.artist)

// newuser := User{
// 	name: "The Weeknd"
// 	email: "theweeknd@weeknd.com"
// 	phone: "999999"
// 	address: Address{
// 		houseno: "27-152"
// 		area: "Toronto"
// 	}
// }

// creation of a pointer

// wgptr := &wickedgames
// fmt.Println(wgptr)

// fmt.Println(*&wgptr.artist)
// fmt.Printf("%p is the pointer\n", wgptr)
// fmt.Printf("%s is the pointer\n", wgptr.name)

// wgptr.artist = "theweeknd"
// fmt.Println(wgptr.artist)

// for anonymous struct
// user := struct {
// 	name  string
// 	email string
// }{
// 	name:  "kaushal singh",
// 	email: "kaushal@kaushal.com",
// }

// fmt.Println(user.name)

// for embedded structs

// user := User{
// 	name: "kaushalsingh",
// 	Where: Where{
// 		hno: "222222",
// 	},
// }

// fmt.Println(user.Where.hno)

// xo := Student{
// 	name:  "theweeknd",
// 	marks: 100,
// }

// to := Student{
// 	name:  "theweeknd",
// 	marks: 100,
// }

// fmt.Printf("%p\n%p", &xo, &to)

// altermarks(&xo)

// fmt.Println(xo.marks)

// }

// ===================================== DAY 6 ====================================================

// func main() {
// 	fmt.Println("The Weeknd is the GOAT bro")

// 	// mapp := map[string]int{
// 	// 	"titvik": 48,
// 	// }
// 	// mapp["kaushal"] = 56
// 	// mapp["sushanth"] = 55

// 	// fmt.Println(mapp["titvik"])

// 	// const hello int = 44

// 	// mapp := map[string]int{
// 	// 	"kaushal":  56,
// 	// 	"sushanth": 55,
// 	// 	"ritvik":   48,
// 	// }

// 	// fmt.Println(mapp["kaushal"])

// 	// mapp["ritvik"] = 33
// 	// fmt.Println(mapp["ritvik"]) // 33

// 	// delete(mapp, "ritvik")
// 	// fmt.Println(mapp)

// 	// value, isExist := mapp["sanam"]
// 	// fmt.Printf("%d key %b", value, isExist)
// 	// // fmt.Printf(value + " " + isExist) // produces error

// 	fmt.Println(mapp)

// }

// ===================================== DAY 5 ====================================================

// func main() {
// 	fmt.Println("The Weeknd is the GOAT bro")

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
// }

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
