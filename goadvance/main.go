package main

// ====================================== DAY 17 ====================================================

// func theTruth() {
// 	fmt.Println("The Weeknd Trilogy is the best album")
// }

// func worker(id int) {
// 	fmt.Printf("The worker started %d\n", id)
// 	fmt.Printf("The worker stopped %d\n", id)
// }

// func api(url string) {
// 	resp, err := http.Get(url)

// 	if err != nil {
// 		fmt.Printf("There is some error %v\n", err)
// 	}

// 	fmt.Printf("%s -> %d\n", url, resp.StatusCode)
// 	fmt.Println(resp)
// }

// func main() {
// 	var wg sync.WaitGroup

// wg.Add(1)

// go func() {
// 	defer wg.Done()

// 	theTruth()
// }()

// wg.Wait()

// fmt.Println("Main finished")

// for i := 1; i <= 5; i++ {
// 	wg.Add(1)

// 	go func() {
// 		defer wg.Done()

// 		worker(i)
// 	}()
// }

// wg.Wait()

// fmt.Println("All Workder finished")

// wg.Add(1)

// go func() {
// 	defer wg.Done()
// 	api("https://v2.jokeapi.dev/joke/Any?format=json")
// }()

// wg.Wait()

// fmt.Println("There is nothing you can do bro to wait and see the above result")

// best practice

// 	urls := []string{"www.google.com"}
// 	for _, url := range urls {
// 		wg.Add(1)

// 		go func(u string) {
// 			defer wg.Done()
// 			// call the function with 1
// 		}(url)
// 	}

// 	wg.Wait()

// 	fmt.Println("How much time can you wait bro")

// }

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

// type Stack[T any] struct {
// 	items []T
// }

// func main() {
// 	// fmt.Println(calculator.Add(4, 5))

// 	fmt.Println("The Weeknd is the GOAT bro")

// 	// // nums := []string{"theweeknd", "kaushal"}
// 	// fmt.Println(first("theweeknd", 100))

// 	// fmt.Println(swap(10, 2))

// 	// numbers := []int{1, 2, 3, 4, 5}
// 	// mapp := map[int]int{}

// 	// for _, val := range numbers {
// 	// 	mapp[val] = val * val
// 	// }

// 	// fmt.Println(mapp)

// 	stack := Stack[int]{
// 		items: []int{},
// 	}

// 	stack.items = append(stack.items, 67)
// 	stack.items = append(stack.items, 67)
// 	stack.items = append(stack.items, 67)
// 	stack.items = append(stack.items, 67)
// 	stack.items = append(stack.items, 67)
// 	stack.items = append(stack.items, 67)

// 	fmt.Println(stack.items)

// }
