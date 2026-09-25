open System
open System.Threading.Tasks
open Microsoft.AspNetCore.Builder
open Microsoft.AspNetCore.Http
open Microsoft.Extensions.Hosting

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    let app = builder.Build()

    app.Use(fun (context: HttpContext) (next: RequestDelgate) -> 
        task {
            printfn "Middleware 1 Before"
            do! next.Invoke(context)
            printfn "Middleware 1 After"
            result
        }
    ) |> ignore

    app.Use(fun context next -> 
        task{
            printfn "Middleware 2 Before"
            let result = next.Invoke()
            printfn "Mideleware 2 After"
            result
        }
    ) |> ignore

    app.MapGet(
        "/", 
        Func<string>(fun () -> 
            printfn "ENDPOINT" 
            "Hello World!"
        )
    ) |> ignore

    // app.MapPost(
    //     "/search", 
    //     Func<string, IResult>(fun (q: string) -> 
    //         if q = "" || q = null then
    //             Results.BadRequest("q is required")
    //         else
    //             Results.Ok({| TheWeeknd = q |})
    //     )
    // ) |> ignore

    app.Run()
    0