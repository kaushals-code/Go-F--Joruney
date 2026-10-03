open System
open Microsoft.AspNetCore.Builder
open Microsoft.Extensions.Hosting
open Microsoft.Extensions.Logging
open LoggingDemo
open Microsoft.Extensions.DependencyInjection   // <-- this one

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    builder.Services.AddSingleton<UserService>() |> ignore
    let app = builder.Build()

    app.MapGet(
        "/", 
        Func<UserService, string>(fun service ->
            service.CreateUser 67
            "User Created"
        )
    ) |> ignore
    app.Run()

    0 // Exit code

