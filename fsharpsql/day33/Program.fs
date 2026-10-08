open System
open Microsoft.AspNetCore.Builder
open Microsoft.Extensions.Hosting
open Npgsql

let checkDatabase(dataSrouce : NpgsqlDataSource) = 
    task {
        use command = dataSource.CreateCommand("SELECT 1")

        let! result = command.ExecuteScalarAsync()

        return result
    }

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    let app = builder.Build()

    app.MapGet("/", Func<string>(fun () -> "Hello World!")) |> ignore

    let connection = new NpgsqlConnection(builder.Configuration.["ConnectionString:DB_STRING"])
    printfn "The string is : %A" connection

    app.MapGet(
        "/dbcheck",
        Func<string>(fun () -> 
            connection.Open()
            "Connected to " + connection.Database
        )
    ) |> ignore

    app.Run()

    0 // Exit code