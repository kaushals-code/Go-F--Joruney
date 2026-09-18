// For more information see https://aka.ms/fsharp-console-apps
// printfn "The Weeknd is the GOAT"

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

// let firstAsync =
//     async {
//         printfn "The Weeknd is the GOAT BRO"
//         return 67
//     }

// let res = Async.RunSynchronously firstAsync
// printfn "%d" res

// let printData n = 
//     async {
//         printfn "Go fuck yourself %d" n
//     }

// let func n = 
//     async {
//         // let! x = async {return 10}
//         // do! printData
//         // return x

//         // do! Async.Sleep 1000

//         // return! (printData n)

//         // do! (printData n)

//         printfn "%d is starting" 
//         Async.Sleep 1000
//         printfn "%d is ending"
//     }

// // let res1 = Async.RunSynchronously (func 2)
// // let res = Async.RunSynchronously (func 1)

// // printfn "%s" res
// // printfn "%s" res1

// let jobs = 
//     [1..5]
//     |> List.map func

// let all = 
//     jobs
//     |> Async.Parallel

// let res = 
//     all
//     |> Async.RunSynchronously

// printfn "%A" res

// best example for async operations on F#

// let func n = 
//     async {
//         printfn "%d is starting\n" n 
//         do! Async.Sleep 1000
//         printfn "%d is ending\n" n
//     }

// let jobs = 
//     [1..5]
//     |> List.map func

// let all = 
//     jobs
//     |> Async.Parallel

// let res = 
//     all
//     |> Async.RunSynchronously

// printfn "%A" res

// tasks now

// open System.Threading.Tasks

// let work = 
//     task {
//         return 67
//     }

// let res = Async.RunSynchronously work
// let res = work

// printfn "%d" res

// let work id = 
//     task {
//         do! Task.Delay 1000

//         return id
//     }

// let tasks = 
//     [1..5]
//     |> List.map work

// let all = 
//     Task.WhenAll tasks

// let result = all    

// printfn "%A" result

// open System.Net.Http

// let client = new HttpClient()

// let api (clt: HttpClient) (url: string) = 
//     async {
//         let! resp = 
//             clt.GetAsync(url)
//             |> Async.AwaitTask
//         return resp
//     }

// let res = Async.RunSynchronously (api client "https://v2.jokeapi.dev/joke/Any?format=json")
// printfn "%A" res

// open System.Net.Http

// let client = new HttpClient()

// let resp (cli: HttpClient) (url: string) = 
//     async {
//         let! res = 
//             cli.GetAsync(url)
//             |> Async.AwaitTask
        
//         return res
//     }

// let rest = 
//     Async.RunSynchronously (resp client "https://v2.jokeapi.dev/joke/Any?format=json")

// printfn "%A" rest
