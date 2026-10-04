open System
open Microsoft.AspNetCore.Builder
open Microsoft.Extensions.Hosting
open Microsoft.Extensions.Logging
open Microsoft.Extensions.DependencyInjection
open System.Net.Http
open System.Text.Json

open LoggingDemo

type GitHubUser =
    {
        login: string
        name: string
        followers: int
        following: int
        public_repos: int
    }

[<EntryPoint>]
let main args =
    let builder = WebApplication.CreateBuilder(args)
    builder.Services.AddHttpClient<GithubService>() |> ignore
    let app = builder.Build()

    // app.MapGet(
    //     "/", 
    //     Func<UserService, string>(fun service ->
    //         service.CreateUser 67
    //         "User Created"
    //     )
    // ) |> ignore

    // let url = "https://official-joke-api.appspot.com/random_joke"

    let client = new HttpClient()

    // let! response = client.GetAsync(url)
    // let getExample () = 
    //     task {
    //         let! response = client.GetAsync(url)
    //         printfn "The status code is %O" response.StatusCode
    //     }

    // let res = (getExample ()).GetAwaiter().GetResult()
    // printfn "%A is the result" res

    let gitService = GithubService(client)
    let res = gitService.GetGithub("kaushals-code").GetAwaiter().GetResult()

    // printfn "%A is the status code" res

    let body = res.Content.ReadAsStringAsync().GetAwaiter().GetResult()
    let user = JsonSerializer.Deserialize<GitHubUser>(body)

    printfn "%s is the name" user.name
    printfn "%d is the followers" user.followers

    app.Run()

    0 // Exit code

