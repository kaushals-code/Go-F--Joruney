open System
open Microsoft.AspNetCore.Builder
open Microsoft.AspNetCore.Http
open Microsoft.Extensions.Hosting
open System.Text.Json

type User = 
    {
        Name: string
        age: int
        email: string opiton
    }

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    let app = builder.Build()

    app.MapGet("/", Func<string>(fun () -> "Hello World!")) |> ignore

    app.MapGet("/hello", Func<string>(fun () -> 
        "The Weeknd is the GOAT bro"
        )
    ) |> ignore

    // let weeknd = app.MapGroup("/weeknd")

    // weeknd.MapGet("/hello", Func<string>(fun () -> 
    //     "Hello from Weeknd")) |> ignore

    // weeknd.MapGet("/hello", Func<string, string>(fun (name: string) -> 
    //     $"Hello {name}"
    //     )
    // ) |> ignore

    // weeknd.MapGet("/hello", Func<string, IResult>(fun (name: string) ->
    //     Results.Ok($"Hello bro {name}")
    //     )
    // ) |> ignore

    let user =
        {
            Name= "theweeknd"
            age= 36
            // email= "theweeknd@gmail.com"
        }

    app.MapGet(
        "/todo", Func<IResult>(
        fun () ->
            // Results.NotFound("/users/id created")
            Results.Ok(
                user
            )
    ))
    |> ignore

    app.Run()

    0 // Exit code

