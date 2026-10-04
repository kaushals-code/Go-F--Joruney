open System
open System.Net.Http

// let! response = Http.GetAsync
let client = Http.Client()

let! response = client.GetAsync("https://api.github.com/users/kaushals-code").GetAwaiter().GetResult()

printfn "%A" response