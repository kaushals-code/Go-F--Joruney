open System
open Microsoft.AspNetCore.Builder
open Microsoft.Extensions.Hosting
open System.Text.Json
open Microsoft.AspNetCore.Http.Json
open System.Collections.Generic
open Microsoft.AspNetCore.Http
open Microsoft.Extensions.DependencyInjection

type Todo = {
    Id: int
    Title: string
    Completed: bool
}

type CreateTodo = {
    Title: string
    Completed: bool
}

type UpdateTodo = {
    Title: string
    Completed: bool
}

let todos = List<Todo>([
    {
        Id = 1;
        Title = "Learn F# JSON";
        Completed = false;
    }
    {
        Id = 2;
        Title = "Build an API";
        Completed = false;
    }
])

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    builder.Services.ConfigureHttpJsonOptions(fun options -> 
        options.SerializerOptions.PropertyNamingPolicy <- 
            JsonNamingPolicy.CamelCase
    ) |> ignore
    let app = builder.Build()

    app.MapGet("/", Func<string>(fun () -> "Hello World!")) |> ignore

    app.MapGet(
        "/todo",
        Func<IResult>(
            fun () -> 
                Results.Ok(todos)
        )
    ) |> ignore

    app.MapGet(
        "/todo/{id}",
        Func<int, IResult>(fun (id: int) -> 
            match todos |> Seq.tryFind (fun todo -> todo.Id = id) with 
            | Some x -> 
                Results.Ok(x)
            | None -> 
                Results.NotFound(
                    {| error= "todo not found" |}
                )
        )
    ) |> ignore

    app.MapPost(
        "/todo",
        Func<CreateTodo, IResult>(fun (todo: CreateTodo) -> 

            let newId = 
                if todos.Count = 0 then
                    0
                else 
                    todos
                    |> Seq.map (fun t -> t.Id)
                    |> Seq.max
            
            let toadd = {
                Id= newId + 1
                Title= todo.Title
                Completed= todo.Completed
            }
            todos.Add(toadd)
            Results.Created(
                $"User created successfully {toadd.Id}", toadd
            )
        )
    ) |> ignore

    app.MapPatch(
        "/todo/{id}",
        Func<int, UpdateTodo, IResult>(
            fun (id: int) (utodo: UpdateTodo) ->  
                let idx = todos.FindIndex(fun x -> x.Id = id)
                if idx = -1 then 
                    Results.NotFound({|error = "Todo not found"|})
                else 
                    let updated = {
                        Id= id
                        Title= utodo.Title
                        Completed= utodo.Completed
                    }
                    todos.[idx] <- updated
                    Results.Ok(updated)
        )
    ) |> ignore

    app.MapDelete(
        "/todo/{id}",
        Func<int, string>(fun (id: int) -> 
            let idx = todos.FindIndex(fun x -> x.Id = id)
            if idx = -1 then
                "The Todo does not exist"
            else 
                todos.RemoveAt(idx)
                "The Todo is removed successfully"
        )
    ) |> ignore

    app.Run()

    0 // Exit code

