open System
open Microsoft.AspNetCore.Builder
open Microsoft.AspNetCore.Http
open Microsoft.Extensions.Hosting

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    let app = builder.Build()

    app.MapGet("/", Func<string>(fun () -> "Hello World!")) |> ignore

    // using RouteParams
    app.MapGet(
        "/name/{name}", 
        Func<string, string>(
            fun (name: string) -> 
                $"The given name is {name}"
    )) |> ignore

    app.MapGet("/bool", Func<HttpRequest, string>(fun request -> 
        let completed = request.Query["completed"].ToString()
        $"The Val is set as {completed}"
    )) |> ignore

    // using a basic map group
    let api = app.MapGroup("/todos")
    api.MapGet("/all", Func<string>(fun x -> 
        "Here all the todos"
    )) |> ignore
    api.MapGet("/delete", Func<string>(fun x -> 
        "The todo is deleted"
    )) |> ignore

    app.Run()

    0 // Exit code

