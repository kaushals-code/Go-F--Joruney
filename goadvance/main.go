package main

import "fmt"

// ====================================== DAY 16 =====================================================

// func first(nums []int) int {
// 	return nums[0]
// }

// here we loose the type safety
// func first(nums []int) any {
// 	return nums[0]
// }

// func first[T any](values []T) T {
// 	return values[0]
// }

// func first[A any, B any](a A, b B) (A, B) {
// 	return a, b
// }

// type Number interface {
// 	int | int64 | float64
// }

// type Number interface {
// 	~int | ~int64 | ~float64
// }

// func swap[T Number](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	items []T
}

func main() {
	// fmt.Println(calculator.Add(4, 5))

	fmt.Println("The Weeknd is the GOAT bro")

	// // nums := []string{"theweeknd", "kaushal"}
	// fmt.Println(first("theweeknd", 100))

	// fmt.Println(swap(10, 2))

	// numbers := []int{1, 2, 3, 4, 5}
	// mapp := map[int]int{}

	// for _, val := range numbers {
	// 	mapp[val] = val * val
	// }

	// fmt.Println(mapp)

	stack := Stack[int]{
		items: []int{},
	}

	stack.items = append(stack.items, 67)
	stack.items = append(stack.items, 67)
	stack.items = append(stack.items, 67)
	stack.items = append(stack.items, 67)
	stack.items = append(stack.items, 67)
	stack.items = append(stack.items, 67)

	fmt.Println(stack.items)

}
