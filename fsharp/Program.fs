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

// let numbers = 
//     Array.init 5 (fun i -> i * 10)

// for i = 0 to (numbers.Length - 1) do
//     printfn "%d" numbers[i]


// ============================================== DAY 4 PRACTICE ====================================================
// List Implementation
// let numbers = [1; 2; 3; 4; 5]

// for i = 0 to (numbers.Length - 1) do 
//     printfn "%d" numbers[i]

// Array 
// let mylist = [| 1; 2; 3; 4; 5 |]
// mylist[0] <- 200

// for i = 0 to (mylist.Length - 1) do 
//     printfn "%d" mylist[i]

// Sequence
// let sequence : seq<int> = seq {1; 2; 3; 3; 5}

// Operations on Array
// let numbers = [1; 2; 3; 4; 5]
// let squares = 
//     [for x in numbers do 
//         yield x * x]

// for i = 0 to (squares.Length - 1) do 
//     printfn "%d" squares[i]

// Array.init

// let nums = 
//     Array.init 5 (fun i -> (i + 1) * (i + 1))

// for i = 0 to (nums.Length - 1) do 
//     printfn "%d" nums[i]

// ============================================== DAY 5 =====================================================

// let numbers = [1; 2; 3; 4; 5; 6; 7; 8; 9; 10]

// let even = 
//     numbers 
//     |> List.filter (fun x -> x % 2 = 0)

// let sum = 
//     numbers
//     |> List.fold (fun total n -> total + n) 0

// let chain = 
//     numbers 
//     |> List.filter (fun x -> x % 2 = 0) 
//     |> List.map (fun x -> x * 2)
//     |> List.fold (fun total x -> total + x) 0

// printfn "%d" chain

// let nums = [4; 2; 6; 8; 5; 1; 0]

// let sorted = 
//     nums 
//     |> List.sort

// printfn "%A" sorted

// type Student =
//     {
//         Name: string
//         Math: int
//         Physics: int
//         Chemistry: int
//         Grade: string
//     }

// let students = 
//     [
//         {
//             Name = "Alice"
//             Math = 85
//             Physics = 43
//             Chemistry = 93
//             Grade = "A"
//         };
//         {
//             Name = "Bob"
//             Math = 84
//             Physics = 99
//             Chemistry = 93
//             Grade = "B"
//         };
//         {
//             Name = "Charlie"
//             Math = 80
//             Physics = 88
//             Chemistry = 66
//             Grade = "A"
//         }
//     ]

// let order = 
//     students
//     |> List.sortByDescending (fun s -> s.Math)

// let order = 
//     students
//     |> List.sumBy (fun x -> x.Math)

// let order = 
//     students
//     |> List.averageBy (fun x -> float x.Math)

// let groupby = 
//     students
//     |> List.groupBy (fun x -> x.Grade)

// printfn "%A" groupBy

// let numbers = [1; 2; 3; 4; 5; 6; 6]
// let result = 
//     numbers 
//     |> List.choose (fun x -> 
//         if x % 2 = 0 then
//             Some (x * 2)
//         else 
//             None)

// printfn "%A" result

// let numbers = [1; 2; 3; 4; 5]
// let result = 
//     numbers
//     |> List.collect (fun x -> [x; x])

// printfn "%A" result

// the above operations with Array also
// let numbers = [|4; 1; 5; 3; 8; 9; 5|]

// let sorted = 
//     numbers 
//     |> 

// printfn "%A" sorted


// ======================================= DAY 6 ===========================================================

// let rollnos =
//     Map.ofList [
//         ("kaushal", 56)
//         ("theweeknd", 1)
//         ("sushanth", 55)
//     ]
// maps are immutable

// let newrolls = 
//     rollnos
//     |> Map.remove "kaushal"

// printfn "%d" rollnos["kaushal"]

// let iskaushalthere = 
//     rollnos
//     |> Map.tryFind "kaushal"


// match rollnos |> Map.tryFind "sushanth" with 
// | Some(x) -> 
//     printfn "Age: %d" x
// | None -> 
//     printfn "sushanth not found"

// rollnos
// |> Map.iter (fun x y -> 
//     printfn "%s -> %d" x y)

// it seems that the counting the frequeincies of the words are so hard here
// although let me try it once

//lets revise Map.fold

// let number = [1; 2; 3; 4; 5]

// let sum = 
//     number
//     |> List.fold (fun total x -> total + x) 0

// printfn "%d is the total" sum

// based on the same concept

// let words = ["go"; "fsharp"; "go"; "rust"; "go"; "fsharp"]

(*
we want 
    go -> 3
    fshart -> 2
    rust -> 1
*)

// let frequencies = 
//     words 
//     |> List.fold (fun counts word ->
//         let current = 
//             counts
//             |> Map.tryFind word
//             |> Option.defaultValue 0
        
//         counts
//         |> Map.add word (current + 1)
//     ) Map.empty

// frequencies
// |> Map.iter (
//     fun x y -> printfn "%s -> %d" x y
// )

// open System.Collections.Generic

// let ages = Dictionary<string, int>()

// ages["kaushal"] <- 19
// ages["theweeknd"] <- 36

// ages["kaushal"] <- 20

// ages["sushanth"] <- 55

// for KeyValue(key, value) in ages do 
//     printfn "%s -> %d" key value

// ages.Remove("sushanth") |> ignore

// for KeyValue(key, value) in ages do 
//     printf "%s -> %d" key, value

// open System.Collections.Generic

// let rollnos = Dictionary<string, int>()

// rollnos["kaushal"] <- 56
// rollnos["sushanth"] <- 55
// rollnos["ritvik"] <- 48

// for KeyValue(key, value) in rollnos do 
//     printfn "%s -> %d" key value

// I didn't remember that the list could be iterated like this
// let rollnos = [1; 2; 3; 4; 5]
// for i in rollnos do 
//     printfn "%d" i

// let data = 
//     [
//         ("kaushal", 56)
//         ("triansh", 60)
//         ("theweeknd", 1)
//     ]

// let mapp = Map.ofList data

// mapp
// |> Map.iter (
//     fun x y -> printfn "%s -> %d" x y
// )

// let res = 
//     mapp
//     |> Map.tryFind "abcd"
//     |> Option.defaultValue 0

// let res = 
//     mapp
//     |> Map.find "kaushal"

// printfn "%d" res

// let newmap = 
//     mapp
//     |> Map.add "triansh" 45

// newmap
// |> Map.iter (fun x y -> printf "%s -> %d\n" x y)

(*
kaushal -> 56
theweeknd -> 1
triansh -> 45
*)

// let remmap = 
//     newmap
//     |> Map.remove "goat"

// remmap
// |> Map.iter (fun x y -> printf "%s -> %d\n" x y)

// let res = 
//     newmap
//     |> Map.containsKey "theweeknd"

// printfn "%b" res

// printfn "%d" newmap.Count

// ============================================= DAY 8 =======================================================

// type Song = {
//     Name: string
//     Album: string
//     Artist: string
// }

// let wickedgames = {
//     Name = "Wicked Games"
//     Album = "Trilogy"
//     Artist = "The Weeknd"
// }

// printfn "%F" wickedgames

// type User = {
//     Name: string
//     Age: int
// }

// let user1 = {
//     Name = "User 1"
//     Age = 18
// }


// can be printed like this using a function
// let printuser user = 
//     printfn "%s" user.Name
//     printfn "%d" user.Age

// printuser user1

// let user2 = {
//     Name = "User 2"
//     Age = 20
// }

// let userdupe = {
//     user1 with 
//         Age = 25
//         Name = "User Dupe"
// }

// printfn "%s" userdupe.Name

// type Album = {
//     Name: string
//     Artist: string
// }

// type Song = {
//     Song_Name: string
//     Album: Album
// }

// let wg = {
//     Song_Name = "Wicked Games"
//     Album = {
//         Name = "Trilogy"
//         Artist = "The Weeknd" 
//     }
// }

// printfn "%s" wg.Album.Artist

// anonymous records
// let point = {| X= 10; Y= 20 |}
// printfn "%d %d" point.X point.Y


// type User = {
//     Name: string
//     Age: int
// }

// let user1 = {
//     Name = "User 1"
//     Age = 17
// }

// let alterage user = 
//     {user with 
//         Age = user.Age + 1}

// let checkage user = 
//     if user.Age >= 18 then 
//         true
//     else 
//         false

// // let newuser = alterage user1
// // printfn "%d" newuser.Age

// let canapplydl = checkage user1
// printfn "%b" canapplydl

// ========================================== DAY 9 ============================================================

// type User = 
//     {
//         name: string
//         email: string
//     }

//     member this.fact = 
//         this.name + " has the email " + this.email

//     member this.greet(name: string) =
//         printfn "Hello %s" name
    
// // let greetperson name = 
// //     "Hello " + name

// // printfn "%s" (greetperson "kaushal")

// let user = {
//     name = "theweeknd"
//     email = "kaushal21gs@gmail.com"
// }

// // printfn "%s" user.fact

// printfn "%s" (user.greet "the weeknd")

// type User = 
//     {
//         name: string
//         age: int
//     }

//     member this.fact =
//         this.name + " age is " + (string this.age)
    
//     static member truth = 
//         "humans are mortal"

// let user = {
//     name = "The Weeknd"
//     age = 36
// }

// printfn "%s" User.truth

// implementing a constructor to the type

// type User = 
//     {
//         name: string
//         age: int
//     }

//     static member create(name: string, age: int) =
//         {
//             name = name
//             age = age
//         }
    
// let newuser = {
//     user with 
//         name = "the weeknd"
//         age = 36
// }

// printfn "%s" newuser.name
// printfn "%d" newuser.age

// type User with 
//     member this.fact =
//         this.name + " age is " + (string this.age)

// let user = User.create ("kaushal", 19)


// printfn "%s" user.fact

// type User = 
//     {
//         FirstName: string
//         LastName: string
//         Age: int
//     }

// // here is the module
// module UserFunctions = 
    
//     let fullName user = 
//         user.FirstName + " " + user.LastName
    
//     let isAdult user = 
//         user.Age >= 18

// let user = {
//     FirstName= "the"
//     LastName = "weeknd"
//     Age= 36
// }

// printfn "%b" (UserFunctions.isAdult user)
// printfn "%s" (UserFunctions.fullName user)

// we can also use pipelines
// let fullname = 
//     user
//     |> UserFunctions.fullName

// printfn "%s" fullname

// type User(name: string, age: int) = 
//     member val Name = name
//     member val Age = age

// let user = User("kaushal", 19)

// printfn "%s" user.Name
// printfn "%d" user.Age

// immutable objects

// type User(name: string, age: int) =
//     member val name = name with get, set
//     member val age = age with get, set

// let user = User("theweeknd", 36)

// user.age <- 136

// printfn "%d" user.age

// type User with 
//     member this.isadult = 
//         this.age >= 18

//     member this.fromgreetings(name: string) = 
//         "Hello " + this.name + " from " + name
    
// printfn "%b" user.isadult
// printfn "%s" (user.fromgreetings("kaushal"))

// type User = 
//     {
//         name: string
//         age: int
//     }

//     member this.create(name: string, age: int) = 
//         User(
//             .WithNname(name)
//             .Withage(age)
//         )
        
// let user = User.create("kaushal", 19)

// printfn "%d" user.age

// =========================================== DAY  10 ===========================================

// just basics
// type User = {
//     name: string
//     age: int
// }

// let user = {
//     name= "kaushal"
//     age= 19
// }

// printfn "%s age is %d" user.name user.age

// type Shape = 
//     | Rectangle of Length : int * Width : int
//     | Circle of Radius : int

// let rect = Rectangle(10, 5)

// let area = 
//     match rect with 
//     | Rectangle (len, wid) -> 
//         printfn "%d is the area" (len * wid)
//     | Circle rad -> 
//         printfn "%f is the area" (3.14 * float (rad * rad))

// type Person =
//     | Student of name: string
//     | Teacher of name: string * salary: int

// let func t = 
//     match t with 
//         | Student n -> 
//             // printfn "%s is the student name" n
//             // n + " is the student name" // not correct
//             $"{n} is the student name"
//         | Teacher (n, s) -> 
//             // printfn "%s is the teacher name and his/her salary is %d" n s
//             // n + " is the teacher name and his/her salary is " + s
//             $"{n} is the teacher name and his/her salary is {s}"

// let t = Teacher("Sanjana", 100000)
// let s = Student("kaushal")

// printfn "%s" (func t)
// printfn "%s" (func s)

// Typical class in F# is (just revision)
// type User =
//     {
//         Name: string
//         Age: int
//     }

//     member this.isAdult = 
//         if this.Age >= 18 then
//             true
//         else 
//             false

//     static member createUser(name: string, age: int) =
//         {
//             Name= name
//             Age= age
//         }

// let user = User.createUser("kaushal", 19)

// printfn "%d" user.Age
// printfn "%b" user.isAdult

// F# objects revision
// type User(name: string, age: int) = 
//     member val Name = name with get, set
//     member val Age = age with get, set

// let user = User("kaushal", 19)

// printfn $"{user.Name} age is {user.Age}"

// This is a interface
// type ISpeaker = 
//     abstract member Speak(): unit -> string

// type Dog(name: string) = 
//     interface ISpeaker with 
//         member _.Speak() = 
//             "Woff! my name is " + name
    
// let dog = Dog("sanam")

// let speaker = dog :> ISpeaker

// printfn "%s" (speaker.Speak())