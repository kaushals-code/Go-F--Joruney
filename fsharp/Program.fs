// // printfn "Hello from F#"

// let name = "The Weeknd"
// let age = 30
// let height = 1.75
// let isArtist = true


// // printfn "%s" name
// // printfn "%d" age
// // printfn "%.2f" height
// // printfn "%b" isArtist


// printfn "========================================"
// printfn "                  MY PROFILE            "
// printfn "========================================"

// printfn "Name:\t\t %s" name
// printfn "Age:\t\t %d" age
// printfn "Height:\t\t %.2f" height
// printfn "isArtist:\t %b" isArtist

// printfn "========================================"



// =============================== DAY 2 ==============================


// let marks = 85

// let grade = 
//     if marks > 90 then
//         "A"
//     elif marks >= 75 then
//         "B"
//     elif marks >= 60 then
//         "C"
//     else
//         "F"

// printfn "Grade: %s" grade


// let number = 5

// match number with 
// | 1 -> printfn "One"
// | 2 -> printfn "Two"
// | 3 -> printfn "Three"
// | _ -> printfn "Something Else"

// let number = 5

// let ans = 
//     if number = 4 then
//         "Four"
//     elif number = 5 then
//         "Five"
//     else 
//         "Something Else"

// printfn "%s is the answer" ans

// let name = "theweekd"

// let result =
//     match name with 
//     | "theweeknd" -> "The Weeknd is the Goat bro"
//     | "kaushalsingh" -> "Kaushal Singh this the Goat bro"
//     | _ -> "Duck off bro"

// printfn "%s is the result" result



// this is a complete different language bro, I am liking it

/////////////////////////// simple for loop
// for i = 1 to 100 do 
//     printfn "%d" i


// simple array initialize
// let numbers = [10; 20; 30; 40]

// for number in numbers do 
//     printfn "%d" number

// indentation matters bro

// for i in [1..2..8] do 
//     printfn "%d" i
//     // printfn "%d" i


// let mutable i = 1

// while i <= 5 do
//     printfn "%d" i
//     i <- i + 1

// let mutable name = "theweeknd"

// name <- "kaushalsingh"

// printfn "%s" name


for i in [1..100] do
    let result = 
        if i % 15 = 0 then
            "FizzBuzz"
        elif i % 5 = 0 then
            "Buzz"
        elif i % 3 = 0 then
            "Fizz"
        else
            string i
    printfn "%s" result

