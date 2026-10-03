namespace LoggingDemo

open Microsoft.Extensions.Logging

type UserService(logger : ILogger<UserService>) =
    member _.CreateUser (id: int) =   
        logger.LogInformation("Creating User {id}", id)
        logger.LogWarning("The User is having empty credentials")
        logger.LogError("Something failed")
        logger.LogDebug("The Weeknd is the GOAT")