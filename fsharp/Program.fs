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


// for i in [1..100] do
//     let result = 
//         if i % 15 = 0 then
//             "FizzBuzz"
//         elif i % 5 = 0 then
//             "Buzz"
//         elif i % 3 = 0 then
//             "Fizz"
//         else
//             string i
//     printfn "%s" result


// for i = 1 to 10 do
//     printfn "%d" i


// let number = 31
// let mutable prime = true
// for i = 2 to (number - 1) do
//     match (number % i) with 
//     | 0 -> prime <- false
//     | _ -> ()
// let result = 
//     if prime = true then
//         "The given number is prime"
//     else 
//         "The given number is NOT prime"
// printfn "%s" result

// let x = 10
// let y = 20

// printfn "%d and %d" x y

// ====================================================== DAY 3 ================================================

// let add a b = 
//     a + b

// // let result = add 3 4

// // printfn "%d" result

// let first = add 3

// let second = first 4

// printfn "%d" second

///////////////// not working
// let add a b = 
//     a + b

// let arithemetic func c d
//     func c d

// let result = arithemetic add 8 9
// printfn "%d" result


/////////////////////////////// PIPE function usage
// let double = 
//     fun x -> 2 * x

// // let d = double 5

// let d = 5 |> double

// printfn "%d" d

// let add a b =
//     a + b

// let r = 9 |> add
// let res = 1 |> r 

// printfn "%d" res

// let add a = 
//     fun y -> a + y

// let applyOp func a b = 
//     func a b

// printfn "%d" (applyOp 3 4)


// prime number checker function
// let isPrime n = 
//     if n < 2 then
//         false 
//     else
//         let rec check div = 
//             if div * div > n then
//                 true
//             elif n % div = 0 then
//                 false
//             else 
//                 check (n + 1)
//         check 2
    

// let number = 80 

// if isPrime number then
//     printfn "The number is prime"
// else 
//     printfn "The number is NOT a prime"


// multiplication table generator
// let multiplicatointable n = 
//     for i = 1 to 10 do
//         printf "%d x %d = %d\n" n i (n * i)

// let number = 8

// multiplicatointable number



// factorial

// let mutable ans = 1
// let multiply num = 
//     ans <- ans * num

// let factorial n = 
//     for i = 1 to n do 
//         multiply i
//     ans

// printfn "%d" (factorial 9)



// fibonacci

// let mutable a = 1
// let mutable b = 1 

// printfn "1 \n1 "

// let fibo n = 
//     for i = 1 to (n - 1) do 
//         let tot = a + b
//         printfn "%d " tot

//         a <- b
//         b <- tot

// fibo 10


// Palindrome

// let checkpalindrome (str: string) = 
//     let len = str.Length
    
//     let rec check i = 
//         if i >= (len / 2) then
//             true
//         elif str[i] <> str[len - 1 - i] then
//             false
//         else 
//             check (i + 1)
//     check 0

// let s = "theweeknd"
// let res = checkpalindrome s

// printfn "%b" res



// String Reverse

// Wrong
// let reversestring (str: string) =
//     let mutable input = str
//     let len = str.Length
//     for i = 0 to int ((len / 2) - 1) do
//         let store <- input[i]
//         input[i] <- input[len - 1 - i]
//         input[len - i - 1] <- store
//     input

// let inp = "kaushal"
// let res = reversestring inp

// printfn "%s" res

// ================================================== DAY 4 ==========================================

// a normal list which is immutble by default
// let list = [1; 2; 3; 4; 5]

// for i = 0 to (list.Length - 1) do
//     printfn "%d\n" list[i]

// let a = [1; 2; 3]
// let b = [4; 5; 6]

// let add0toa = 0 :: a
// let final = -1 :: add0toa

// for i = 0 to (final.Length - 1) do
//     printfn "%d" final[i]

// let c = b @ a
// for i = 0 to (c.Length - 1) do
//     printfn "%d" c[i]

// // arrays are mutable by default
// let arr = [| 1; 2; 3; 4; 5 |]

// arr[0] <- 100


// for i = 0 to (arr.Length - 1) do
//     printfn "%d" arr[i]


// a sequence is said to be lazy
// let numbers : seq<int> = seq { 1; 2; 3; 4; 5; 6 }
// it is very usefull because of its lazy value initialization

// let squares = 
//     [for x in 1 .. 10 do
//         if x % 2 = 0 then
//             yield x * x]

// for i = 0 to (squares.Length - 1) do 
//     printfn "%d" squares[i]

let numbers = 
    Array.init 5 (fun i -> i * 10)

for i = 0 to (numbers.Length - 1) do
    printfn "%d" numbers[i]
