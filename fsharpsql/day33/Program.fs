open System
open Microsoft.AspNetCore.Builder
open Microsoft.Extensions.Hosting
open Npgsql

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    let app = builder.Build()

    app.MapGet("/", Func<string>(fun () -> "Hello World!")) |> ignore

    let connection = new NpgsqlConnection(builder.Configuration.["DB_STRING"])

    app.MapGet(
        "/dbcheck",
        Func<string>(fun () -> 
            connection.Open()
            "Connected to " + connection.Database
        )
    ) |> ignore

    app.Run()

    0 // Exit code

