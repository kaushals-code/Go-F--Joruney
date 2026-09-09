package main

import (
	"fmt"
)

func main() {

	// ==================================== DAY 1 ============================================

	// var name string
	// var age int
	// var height float64
	// var isArtist bool

	// var name string = "The Weeknd"
	// age := 50
	// height := 6.12
	// isArtist := true

	// // fmt.Println(name)
	// // fmt.Println(age)
	// // fmt.Println(height)
	// // fmt.Println(isArtist)

	// fmt.Printf("The name is %q\n", name)
	// fmt.Printf("The number is %d\n", age)
	// fmt.Printf("The double is %f\n", height)
	// fmt.Printf("The boolean is %t\n", isArtist)

	///////////////////// ERROR HERE

	// const name = "The Weeknd"
	// name = "Kaushal Singh Thakur"
	// fmt.Println(name)

	//////////////////// IOTA
	// const (
	// 	Pending = iota
	// 	Active  = iota * 10
	// 	Suspended
	// 	Deleted
	// )

	// // status := Active

	// fmt.Println(Pending)
	// fmt.Println(Active)
	// fmt.Println(Suspended)
	// fmt.Println(Deleted)

	// fmt.Printf("======================================\n")
	// fmt.Printf("\t\tProfile\n")
	// fmt.Printf("======================================\n")
	// fmt.Printf("Name: \t\t %s\n", name)
	// fmt.Printf("Age: \t\t %d\n", age)
	// fmt.Printf("Height: \t %.2f\n", height)
	// fmt.Printf("isArtist: \t %t\n", isArtist)
	// fmt.Print("======================================\n")

	//============================================ DAY 2====================================

	// name := "kaushalsingh"

	// if name == "theweeknd" {
	// 	fmt.Printf("%s is the GOAT bro", name)
	// } else if name == "kaushalsingh" {
	// 	fmt.Printf("%s is the GOAT fan bro", name)
	// } else {
	// 	fmt.Printf("%s is the nood bro", name)
	// }

	// if name := 20; name >= 18 {
	// 	fmt.Println("Adult")
	// } else {
	// 	fmt.Println("Noob")
	// }

	////////////////// switch statements
	// day := 3
	// switch day {
	// case 1:
	// 	fmt.Println("Day one")
	// case 2:
	// 	fmt.Println("Day two")
	// default:
	// 	fmt.Println("No Day")
	// }

	////////////////// for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	////////////////// while loop
	// i := 0
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	////////////////// range

	// numbers := []int{10, 20, 30, 40, 50}
	// for index, value := range numbers {
	// 	fmt.Printf("%d -> %d\n", index, value)
	// }

	////////////////// lables keyword

outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i*10 + j)
			if i*10+j == 60 {
				break outer
			}
		}
	}

}
