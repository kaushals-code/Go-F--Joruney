open System
open Microsoft.AspNetCore.Builder
open Microsoft.AspNetCore.Http
open Microsoft.Extensions.Hosting

type CreateUserRequest = {
    Name: string
    Email: string
    age: int
}

// validus example
type CreateUserRequest = {
    Name: string `validate:"required,min=3"`
    Email: string `validate:"required,email"`
    age: int `validate:"gte=18,lte=100"`
}

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    let app = builder.Build()

    app.MapGet("/", Func<string>(fun () -> "Hello World!")) |> ignore

    app.MapGet(
        "/{id}",
        Func<int,string>(fun (id: int) -> 
            $"The id is {id}"
    )) |> ignore

    app.MapGet(
        "/search",
        Func<string, int, IResult>(fun (q: string) (limit: int) -> 
            Results.Ok({|
                query= q
                limit= int limit
            |})
        )
    ) |> ignore

    app.MapPost(
        "/users",
        Func<CreateUserRequest, IResult>(fun (user: CreateUserRequest) -> 
            Results.Ok(user)
        )
    ) |> ignore

    app.Run()

    0 // Exit code

