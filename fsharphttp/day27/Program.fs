open System
open Microsoft.AspNetCore.Builder
open Microsoft.Extensions.Hosting
open Microsoft.AspNetCore.Http

type AppError = 
    | BadRequest    of string
    | Unauthorized  of string
    | Forbidden     of string
    | NotFound      of string
    | Conflict      of string
    | Internal

type User =
    {
        Id:         int
        Name:       string
        Email:      string
    }

type ProblemDetailsResponse = 
    {
        Type:       string
        Title:      string
        Status:     int
        Details:    string
        Instance:   string
    }

let findUser (id: int) : Result<User, AppError> =
    if id = 1 then 
        Ok {
            Id = 1
            Name = "Kaushal"
            Email = "kaushal@email.com"
        }
    else 
        Error (NotFound $"User {id} is not found in the db")

let errorToResult (error: AppError) : IResult =
    match error with
    | BadRequest message ->
        Results.BadRequest(message)
    | Unauthorized message ->
        Results.Unauthorized()
    | Forbidden message ->
        Results.StatusCode(StatusCodes.Status403Forbidden)
    | NotFound message ->
        Results.NotFound(message)
    | Conflict message ->
        Results.Conflict(message)
    | InternalError ->
        Results.StatusCode(StatusCodes.Status500InternalServerError)

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    builder.Services.AppProblemDetails() |> ignore
    let app = builder.Build()

    app.MapGet("/", Func<string>(fun () -> "Hello World!")) |> ignore

    // Result is an F# type
    // -> Ok value
    // -> Error error

    // IResult is an Http type
    // Results.Ok(...)
    // Results.NotFound(...)
    // Results.BadRequest(...)

    // we need to do this
    //   Result<'T,'E>
    //        ↓
    //    mapResult
    //        ↓
    //     IResult

    let user = {
        Id = 1
        Name = "Kaushal"
        Email = "kaushal21gs@gmail.com"
    }

    let result = findUser 2

    match result with 
    | Ok x -> 
        printfn "%A" x
    | Error x -> 
        match x with 
        | BadRequest x -> 
            printfn "Bad Request"
        | NotFound x -> 
            printfn "Not Found"

    app.Run()

    0 // Exit code

