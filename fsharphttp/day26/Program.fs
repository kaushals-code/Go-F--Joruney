open System
open Microsoft.AspNetCore.Builder
open Microsoft.AspNetCore.Http
open Microsoft.Extensions.Hosting

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    let app = builder.Build()

    // app.Use(
    //     Func<HttpContext, RequestDelegate, Threading.Tasks.Task>(
    //         fun context next ->
    //             task {
    //                 printfn "Request received"

    //                 do! next.Invoke(context)

    //                 printfn "Request completed"
    //             }
    //     )
    // ) |> ignore

    // app.Use(
    //     Func<HttpContext, RequestDelegate, Threading.Tasks.Task>(
    //         fun context next -> 
    //             task {
    //                 printfn "Request 1 received"
    //                 do! next.Invoke(context)
    //                 printfn "Request 1 completed"
    //             }
    //     )
    // ) |> ignore

    // app.Use(
    //     Func<HttpContext, RequestDelegate, Threading.Tasks.Task>(
    //         fun context next -> 
    //             task {
    //                 printfn "Request 2 received"
    //                 do! next.Invoke(context)
    //                 printfn "Request 2 completed"
    //             }
    //     )
    // ) |> ignore

    // middleware to stop the http request
    // app.Use(
    //     Func<HttpContext, RequestDelegate, Threading.Tasks.Task>(fun context next -> 
    //         task {
    //             context.Response.StatusCode <- 403
    //             do! context.Response.WriteAsync("Blocked")

    //             return null
    //         }
    //     )
    // ) |> ignore

    // use Items of the HttpContext to store some information
    // app.Use(
    //     Func<HttpContext, RequestDelegate, Threading.Tasks.Task>(
    //         fun context next -> 
    //             task {
    //                 context.Items["RequestID"] <- "abc123"
    //                 do! next.Invoke(context)
    //             }
    //     )
    // ) |> ignore

    // app.Use(
    //     Func<HttpContext, RequestDelegate, Threading.Tasks.Task>(
    //         fun context next -> 
    //             task{
    //                 // let request = context.Items["RequestID"] :?> string
    //                 // printfn "%s is the RequestID" request
    //                 // do! next.Invoke(context)

    //                 let request = 
    //                     match context.Items.TryGetValue("RequestID") with 
    //                     | true, value -> 
    //                         value :?> string
    //                     | false, _ -> 
    //                         ""
    //                 printfn "%s is the RequestID" request
    //                 do! next.Invoke(context)
    //             }
    //     )
    // ) |> ignore

    // use of app.Map()
    app.Map(
        "/admin",
        fun adminApp -> 
            adminApp.Use(
                // the adminApp middlware here
            )
    )

    app.MapGet(
        "/",
        Func<string>(fun () -> 
            printfn "Endpoint executing"
            "Hello World"
        )
    ) |> ignore

    app.Run()
    0