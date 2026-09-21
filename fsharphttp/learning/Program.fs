open System.Text.Json

type Todo = {
    Id: int
    Title: string
    Completed: bool
}

let todo = {
    Id= 1
    Title= "Listen to The Weeknd"
    Completed= true
}

let option = JsonSerializerOptions(
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
)
let res = JsonSerializer.Serialize(todo, option)

let actual = JsonSerializer.Deserialize<Todo>(res)

printfn "%s" res
printfn "%A" actual