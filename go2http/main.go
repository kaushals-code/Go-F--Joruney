package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("The Weeknd is the GOAT")

	res, err := http.Get("http://github.com/kaushals-code")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Status)
}

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
