// For more information see https://aka.ms/fsharp-console-apps
printfn "The Weeknd is the GOAT"

// open Calculator
// open StudentReport // it will import all the moduels present in that namepace
// open Student // I need to import eveything like this if I don't use a namespace
// open StudentReport.Core

// printfn "%d" (Calculator.add 8 9)

// let stu = StudentReport.Student.create "theweeknd" [100; 100; 90; 90]

// let res = StudentReport.Student.average stu

// printfn "%A" stu
// printfn "%f" res

// let result = Student.average stu
// printfn "%f" result

// let res = StudentReport.Core.Say.hello "kaushal"
// printfn "%s" res

// let result = StudentReport.Core.Sing.SingWeeknd
// printfn "%s" result

// ============================================= DAY 16 ===========================================================

// let first lst = 
//     List.head lst

// let nums = [1; 2; 3; 4; 5]
// let names = ["theweeknd"; "kaushal"]

// printfn "%d" (first nums)
// printfn "%s" (first names)

// explictily writing the generic types
// let identity<'T> (x: 'T): 'T = 
//     x

// let makepair<'T, 'U> (x: 'T) (y: 'U) : 'T * 'U = 
//     (x, y)

// let res1 = identity<string> "kaushal"
// let res2 = identity<int> 67

// printfn "%s %d" res1 res

// let withfun<'T, 'U> (x: 'T) (f: 'T -> 'U): 'U = 
//     f x

// let res = withfun 5 (fun x -> x * x)

// printfn "%d" res

// type Stack<'T> = 
//     {
//         Items: 'T list
//     }

// let stack : Stack<int> = 
//     {
//         Items= []
//     }
//     |> push 10
//     |> push 20
//     |> push 30


