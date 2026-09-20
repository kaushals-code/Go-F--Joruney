package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// =============================================== DAY 23 ======================================================

// the todo project

func main() {
	c := chi.NewRouter()

	// the business logic here

	http.ListenAndServe(":8080", c)
}

// func fun() {
// 	fmt.Println("This is executed")
// }

// standard library doest not provide all the subroutes and gruops aslike the chi for go
// so its better to use chi-go for the backend-dev

// this is for the chi (github.com open-source router for go)
// func main() {
// 	c := chi.NewRouter()

// c.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "The Weeknd is the GOAT bro")
// })

// c.Get("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
// 	name := chi.URLParam(r, "name")
// 	fmt.Fprintln(w, "The name is given as", name)
// })

// wildcard in chi
// c.Get("/files/*", func(w http.ResponseWriter, r *http.Request) {
// 	path := chi.URLParam(r, "*")
// 	fmt.Fprintln(w, "The path is given as", path)
// })

// the main use of route
// c.Route("/users", func(r chi.Router) {
// 	r.Get("/name/{name}", func(w http.ResponseWriter, r *http.Request) {
// 		name := chi.URLParam(r, "name")
// 		fmt.Fprintln(w, name)
// 	})
// 	r.Get("/age/{age}", func(w http.ResponseWriter, r *http.Request) {
// 		age := chi.URLParam(r, "age")
// 		fmt.Fprintln(w, age)
// 	})
// })

// using the group also
// c.Route("/version", func(r chi.Router) {
// 	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {})
// 	r.Group(func(r chi.Router) {
// 		r.Use() // ex: authMiddleWare
// r.Get() // to get any information related to the user
// 		})
// 	})

// 	fmt.Println("Its working good")
// 	http.ListenAndServe(":8080", c)
// }

// func main() {
// 	mux := http.NewServeMux()

// 	// this feature in go is very useful
// 	// GET /todo
// 	// POST /todo/{id}
// 	// DELETE /todo/{id}
// 	// here id is called as a path parameter

// 	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
// 		name := r.PathValue("name")
// 		fmt.Fprintln(w, "This is name", name)
// 	})

// 	// go supports wildcards paths also
// 	mux.HandleFunc("GET /files/{path...}", func(w http.ResponseWriter, r *http.Request) {
// 		value := r.PathValue("path")
// 		fmt.Fprintln(w, "The pats provided is", value)
// 	})

// 	fmt.Println("Server running on 8080")
// 	http.ListenAndServe(":8080", mux)
// }

// =============================================== DAY 22 ======================================================

// func handleReq(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Method: ", r.Method)
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello Bro!")
// }

// creating my own handler
// type HelloWeekndHandler struct{}

// func (h HelloWeekndHandler) ServeHTTP(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {
// 	fmt.Fprintln(w, "Hello The Weeknd")
// }

// func (h HelloWeekndHandler) ServeHTTP(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {
// 	fmt.Fprintln(w, "Hello The Weeknd Bro, big fan bro")
// }

// func handleHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello Weeknd from handleHello the new function")
// }

// func main() {

// the most basic and my first server with go
// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "The Weeknd is the GOAT!")
// })

// fmt.Println("Server is running on http://localhost:8080")
// http.ListenAndServe(":8080", nil)
// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// fmt.Fprintln(w, "The Weeknd is the GOAT bro")
// here we can print the attribues of the r like
// r.Method, r.URL, r.Header, r.Body
// 	fmt.Fprintln(w, "Method: ", r.Method)
// })
// fmt.Println("The Port is running on the port 8080")

// http.ListenAndServe(":8080", nil)
// the arguments are the portnumber and the handler
// mux := http.NewServeMux()
// mux.HandleFunc("/hello", helloHandler)
// fmt.Println("The Servers is runnin on the port 8080")
// http.ListenAndServe(":8080", mux)
// mux := http.NewServeMux()

// mux.Handle("/weeknd", HelloWeekndHandler{})
// Handle handles the custom handlers -> also it need to be called as Handler{}
// whereas HandleFunc handles all the function type handlers

// http.HandleFunc("/weeknd", helloHandler)
// fmt.Println("The Servers is runnin on the port 8080")
// http.ListenAndServe(":8080", mux)

// mux := http.NewServeMux()

// mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusNotFound)
// 	// some more status codes built-in are
// 	// StatusOK
// 	// StatusCreated
// 	// StatusBadRequest
// 	// StatusForbidden
// 	fmt.Fprintln(w, "There is something wrong with your url, please check")
// })

// fmt.Println("The server is runnin on the port 8080")
// http.ListenAndServe(":8080", mux)

// mux := http.NewServeMux()

// mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 	value := r.Header.Get("authorizatoin")
// 	fmt.Fprintln(w, "The Authorization is ", value)

// 	// for query parameters, it is like
// 	// value := r.URL.Query().Get("name")
// })

// mux.Handle("/helloweeknd", HelloWeekndHandler{})
// mux := http.NewServeMux()
// mux.HandleFunc("/helloweeknd", func(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		handleHello(w, r)
// 	} else {
// 		fmt.Fprintln(w, "You cannot use the POST method for this endpoint")
// 	}
// })
// fmt.Println("The Server is running on the port 8080")
// http.ListenAndServe(":8080", mux)

// }
