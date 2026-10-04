namespace LoggingDemo

open Microsoft.Extensions.Logging
open System.Threading.Tasks
open System.Net.Http

type UserService(logger : ILogger<UserService>) =
    member _.CreateUser (id: int) =   
        logger.LogInformation("Creating User {id}", id)
        logger.LogWarning("The User is having empty credentials")
        logger.LogError("Something failed")
        logger.LogDebug("The Weeknd is the GOAT")

type GithubService (client: HttpClient) =   
    member _.GetGithub (username: string) = 
        task {
            let url = $"https://api.github.com/users/{username}"
            let! respons = client.GetAsync(url)
            return respons
        }
