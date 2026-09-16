package day15

import (
	"fmt"
	"theweeknd/day15/student"
)

func main() {
	fmt.Println("The Weeknd is the GOAT")

	stu := student.Student{
		Name:  "The Weeknd",
		Marks: 100,
	}
	fmt.Println(stu.Name)
}
