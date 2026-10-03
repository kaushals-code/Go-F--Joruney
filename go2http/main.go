package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GitHubUser struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	Followers int    `json:"followers"`
	Repos     int    `json:"public_repos"`
}

func main() {
	url := "http://api.github.com/users/kaushals-code"

	// fmt.Println("The Weeknd is the GOAT")

	// client := &http.Client{
	// 	Timeout: 5 * time.Second,
	// }

	// res, err := client.Get(url)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer res.Body.Close()

	// var user GitHubUser
	// er := json.NewDecoder(res.Body).Decode(&user)

	// if er != nil {
	// 	fmt.Errorf("There is some error occuring, please try again later")
	// 	return
	// }

	// fmt.Println("Login:", user.Login)
	// fmt.Println("Name:", user.Name)
	// fmt.Println("Followers:", user.Followers)
	// fmt.Println("Repositories:", user.Repos)

	// till here, it is working good and the response is perfectly getting decoded

	// this timeout can also be done at the request level also

	// client := &http.Client{
	// 	Timeout: 3 * time.Second,
	// }

	// ctx, cancel := context.WithTimeout(
	// 	context.Background(),
	// 	3*time.Second,
	// )
	// defer cancel()

	// url := "http://api.github.com/users/kaushals-code"

	// req, err := http.NewRequestWithContext(
	// 	ctx,
	// 	http.MethodGet,
	// 	url,
	// 	nil,
	// )

	// if err != nil {
	// 	fmt.Println("Error creating request:", err)
	// 	return
	// }

	// req.Header.Set("Content-type", "application/json")

	// resp, err := client.Do(req)

	// if err != nil {
	// 	fmt.Println("Error creating request:", err)
	// 	return
	// }

	// defer resp.Body.Close()

	// var user GitHubUser
	// er := json.NewDecoder(resp.Body).Decode(&user)

	// if er != nil {
	// 	fmt.Errorf("There is some error occuring, please try again later")
	// 	return
	// }

	// fmt.Println("Login:", user.Login)
	// fmt.Println("Name:", user.Name)
	// fmt.Println("Followers:", user.Followers)
	// fmt.Println("Repositories:", user.Repos)

	// u, err := url.Parse(url)

	// if err != nil {
	// 	fmt.Println("There is an error")
	// 	return
	// }

	// query := u.Query()
	// u.Set("city", "Hyderabad")

	// u.RawQuery = query.Encode()

	// use of http.Transport in a previous example

	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
	}

	res, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	var user GitHubUser
	er := json.NewDecoder(res.Body).Decode(&user)

	if er != nil {
		fmt.Errorf("There is some error occuring, please try again later")
		return
	}

	fmt.Println("Login:", user.Login)
	fmt.Println("Name:", user.Name)
	fmt.Println("Followers:", user.Followers)
	fmt.Println("Repositories:", user.Repos)

	// its better to retry an api service not at the fixed interval times but in the exponential increasing wating
	// times like after the first failure 100ms, and then 200ms and then 300ms.... so on

	
}

// ======================================== DAY 30 ==================================================

// type UserService struct {
// 	logger *slog.Logger
// }

// func (u *UserService) HandleLogging(msg string) {
// 	slog.Info(msg)
// }

// func loggingMiddleware(logger UserService, next http.Handler) http.Handler {
// 	return http.HandlerFunc((func(w http.ResponseWriter, r *http.Request) {
// 		// slog.Info("Middleware started serving")
// 		logger.HandleLogging("Middleware started serving")
// 		next.ServeHTTP(w, r)
// 		// slog.Info("Middleware ended serving")
// 		logger.HandleLogging("Middleware ended serving")
// 	}))
// }

// func actualFunc(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "The Weeknd is the GOAT BRO")
// }

// func main() {

// 	// initializing the logger
// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelInfo,
// 	}
// 	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

// 	userService := UserService{
// 		logger: logger,
// 	}

// 	r := chi.NewRouter()

// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(w, "Hello World")
// 	})

// 	r.HandleFunc("/hello", actualFunc)

// 	handler := loggingMiddleware(userService, r)

// 	fmt.Println("Server running on port 8080")
// 	http.ListenAndServe(":8080", handler)
// }

// type UserService struct {
// 	logger *slog.Logger
// }

// func NewUserService(logger *slog.Logger) *UserService {
// 	return &UserService{
// 		logger: logger,
// 	}
// }

// func (s *UserService) CreateUser(str string) {
// 	s.logger.Info(str)
// }

// func main() {
// 	// fmt.Println("Hello world")
// 	// log.Println("This is the problem")
// 	// slog.Info("This is slog problem")

// 	// slog.Debug(
// 	// 	"loking up user",
// 	// 	"user_id", "100",
// 	// )

// 	// slog.Warn(
// 	// 	"looking up user",
// 	// 	"user_id", "100",
// 	// )

// 	// slog.Error(
// 	// 	"loking up user",
// 	// 	"user_id", "100",
// 	// )

// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	// handler := slog.NewJSONHandler(os.Stdout, opts)

// 	// logger := slog.New(handler)

// 	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
// 	logs := logger.With("GOAT", "The Weeknd")

// 	// logger.Debug("debug message")
// 	// logger.Info("info message")
// 	// logger.Warn("warning message")
// 	// logger.Error("error message")

// 	// slog.Error(
// 	// 	"The main error is",
// 	// 	"port", 6767,
// 	// )

// 	userService := UserService{
// 		logger: logs,
// 	}

// 	userService.CreateUser("The Weeknd is the GOAT")
// }
