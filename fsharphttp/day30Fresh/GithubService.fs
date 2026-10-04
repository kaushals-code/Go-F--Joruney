// namespace ApiClientDemo
    
// open System.Net.Http
// open System.Text.Json

// type GithubUser =
//     {
//         Login: string
//         Name: string option
//         Followers: int
//         PublicRepos: int
//     }

// type GithubService(httpClient: HttpClient) =

//     member _.GetUser(username: string) =
//         task {
//             let! response =
//                 httpClient.GetAsync($"users/{username}")

//             response.EnsureSuccessStatusCode()

//             let! json =
//                 response.Content.ReadAsStringAsync()

//             let options =
//                 JsonSerializerOptions(
//                     PropertyNameCaseInsensitive = true
//                 )

//             let user =
//                 JsonSerializer.Deserialize<GithubUser>(json, options)

//             return user
//         }