// For more information see https://aka.ms/fsharp-console-apps
printfn "The Weeknd is the GOAT"

open Calculator
open StudentReport // it will import all the moduels present in that namepace
open Student // I need to import eveything like this if I don't use a namespace
open StudentReport.Core

// printfn "%d" (Calculator.add 8 9)

// let stu = StudentReport.Student.create "theweeknd" [100; 100; 90; 90]

// let res = StudentReport.Student.average stu

// printfn "%A" stu
// printfn "%f" res

// let result = Student.average stu
// printfn "%f" result

let res = StudentReport.Core.Say.hello "kaushal"

printfn "%s" res

let result = StudentReport.Core.Sing.SingWeeknd

printfn "%s" result