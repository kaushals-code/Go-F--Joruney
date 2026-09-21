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

type AddTodo = {
    Title: string
    Completed: bool
}

type UpdateTodo = {
    Title: string
    Completed: bool
}

let todos = List<Todo>([
    {
        Id= 1;
        Title= "Learn Go Backend";
        Completed= true;
    } 
    {
        Id= 2;
        Title= "Learn F# Backend";
        Completed= true;
    }
])

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    // don't forget to add the converter type here
    builder.Services.ConfigureHttpJsonOptions(fun options -> 
        options.SerializerOptions.PropertyNamingPolicy <- 
            JsonNamingPolicy.CamelCase
    ) |> ignore
    let app = builder.Build()

    app.MapGet("/", Func<string>(fun () -> "This is a medium size todo app backend")) |> ignore

    app.MapGet(
        "/todo/all",
        Func<IResult>(fun () -> 
            Results.Ok(todos)
        )
    ) |> ignore

    app.MapPost(
        "/todo",
        Func<AddTodo, IResult>(fun (x: AddTodo) -> 
            let newId = 
                if todos.Count = 0 then
                    0
                else 
                    todos 
                    |> Seq.map (fun x -> x.Id)
                    |> Seq.max
            let toadd = {
                Id= newId + 1
                Title= x.Title
                Completed= x.Completed
            }
            todos.Add(toadd)
            Results.Ok(
                toadd
            )
        )
    ) |> ignore

    app.MapPut(
        "/todo/{id}",
        Func<int, UpdateTodo, IResult>(fun (id: int) (utodo: UpdateTodo) -> 
            let idx = todos.FindIndex (fun x -> x.Id = id)
            if idx = -1 then 
                Results.NotFound("Todo with the given id found")
            else 
                let todo = {
                    Id= id
                    Title= utodo.Title
                    Completed= utodo.Completed
                }
                todos.[idx] <- todo  
                Results.Ok("Todo Updated successfully")
        )
    ) |> ignore

    app.MapDelete(
        "/todo/{id}",
        Func<int, IResult>(fun (id: int) -> 
            let idx = todos.FindIndex (fun x -> x.Id = id)
            if idx = -1 then
                Results.NotFound("The Todo not found")
            else 
                todos.RemoveAt(idx)
                Results.Ok("The Todo Removed successfully")
        )
    ) |> ignore

    app.Run()

    0 // Exit code

