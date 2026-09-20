open System
open Microsoft.AspNetCore.Builder
open Microsoft.Extensions.Hosting
open Microsoft.AspNetCore.Http
open System.Collections.Generic

type Todo = {
    todo: string
}

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    let app = builder.Build()

    app.MapGet("/", Func<string>(fun () -> "Hello World!")) |> ignore

    // let todos = List<int>([1; 2; 3; 4; 5])
    // todos.Add(99)
    // for i in todos do 
    //     printfn "%d" i

    let todos = List<Todo>([])
    
    let api = app.MapGroup("/todo")
    
    api.MapPost("/add", Func<HttpRequest, string>(fun (r: HttpRequest) -> 
        let toadd = {
            todo= r.Query["todo"].ToString()
        }    
        todos.Add(toadd)
        "Todo added successfully"
    )) |> ignore

    api.MapGet("/alltodo", Func<List<Todo>>(fun () -> 
        todos
    )) |> ignore

    app.Run()

    0 // Exit code

