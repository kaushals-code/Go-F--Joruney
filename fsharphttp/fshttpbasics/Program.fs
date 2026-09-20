open System
open Microsoft.AspNetCore.Builder
open Microsoft.AspNetCore.Http
open Microsoft.Extensions.Hosting

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    let app = builder.Build()

    // app.MapGet("/", Func<string>(fun () -> "Hello World!")) |> ignore

    // app.MapGet("/hello", Func<string>(fun () -> "Hello from the weeknd fan!")) 
    //     |> ignore

    // app.MapGet("/", Func<IResult>(fun () -> 
    //     Results.Ok("The Weeknd is the goat bro") 
    //     // Results.NotFound()
    // )) |> ignore

    // app.MapGet("/hello", Func<string, string>(fun (name: string) -> 
    //     $"Hello bro {name}"
    // ))

    let helloHandler(name: string) (req: HttpRequest) = 
        $"Hello from the kaushal to {name} and the method is {req.Method}"

    app.MapGet("/hellohandler", Func<string, HttpRequest, string>(
        fun (name: string) (req: HttpRequest) -> helloHandler name req
    )) |> ignore

    app.Run()

    0 // Exit code

