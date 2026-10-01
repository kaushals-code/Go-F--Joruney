    open System
    open Microsoft.AspNetCore.Builder
    open Microsoft.AspNetCore.Http
    open Microsoft.Extensions.Hosting
    open Microsoft.Extensions.Options
    open Microsoft.Extensions.Configuration
    open Microsoft.Extensions.DependencyInjection

    [<CLIMutable>]
    type JwtOptions = 
        {
            Secret: string
        }

    // type DatabaseOptions = 
    //     {
    //         Url: string
    //     }

    [<EntryPoint>]
    let main args =
        let builder = WebApplication.CreateBuilder(args)
        builder.Services.Configure<JwtOptions>(
            builder.Configuration.GetSection("Jwt")
        ) |> ignore
        let app = builder.Build()

        app.MapGet("/", Func<string>(fun () -> "Hello World!")) |> ignore

        // app.MapGet("/config", Func<IConfiguration, IResult>(fun (config : IConfiguration) -> 
        //     Results.Ok(config["SANAM"])
        // )) |> ignore

        // app.MapGet("/config", Func<IOptions<AppOptions>, string>(fun (options: IOptions<AppOptions>) -> 
        //     let config = options.Value
        //     $"Port = {config.PORT}"
        // )) |> ignore

        // app.MapGet("/config", Func<IOptions<AppOptions>, IResult>(fun opitons -> 
        //     let config = options.Value
        //     Results.Ok($"{config.PORT}")
        // )) |> ignore

        // learn IOptions once again

        app.MapGet("/config", Func<IOptions<JwtOptions>, string>(fun options -> 
            options.Value.Secret
        )) |> ignore

        app.Run()

        0 // Exit code

