open System.Text.Json

// type Todo = {
//     Id: int
//     Title: string
//     Completed: bool
// }

// let todo = {
//     Id= 1
//     Title= "Listen to The Weeknd"
//     Completed= true
// }

// let option = JsonSerializerOptions(
//     PropertyNamingPolicy = JsonNamingPolicy.CamelCase
// )
// let res = JsonSerializer.Serialize(todo, option)

// let actual = JsonSerializer.Deserialize<Todo>(res)\

// printfn "%s" res
// printfn "%A" actual

open System
open System.Text.Json

type User = 
    {
        Name: string
        Age: int
        Email: string
    }

let user1 = 
    {
        Name= "kaushal"
        Age= 19
        Email= "example@gmail.com"
    }

let options = JsonSerializerOptions(
    PropertyNamingPolicy= JsonNamingPolicy.CamelCase
)

let ser = JsonSerializer.Serialize(user1, options)

printfn "The Serialized object is -> %A" ser

// I have a serialized object here
// now I need to deserialize it actually

let sers = """{"name":"theweeknd","age":36,"email":"TheWeeknd@example.com"}"""

let real = JsonSerializer.Deserialize<User>(sers, options)

printfn "The name is %s" real.Name
printfn "The email is %s" real.Email
































