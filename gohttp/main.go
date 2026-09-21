package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// =============================================== DAY 24 ======================================================

// this the normal to struct
//
//	type Todo struct {
//		Id        int
//		Title     string
//		Completed bool
//	}
//
// but for the api json responses, we use this
type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty`
}

type Person struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email,omitempty"`
	Gender byte   `json:"-"`
}

// good practice for encoding
// func writeJSON(w http.ResponseWriter, status int, data any) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)

// 	err := json.NewEncoder(w).Encode(data)

// 	if err != nil {
// 		return
// 	}
// }

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}

// in the func main() writeJSON(w, http.BadRequest, ErrorResponse {Error: "Get out bro"})

func main() {

	// the main methods given by encoding/json are
	// Marshall(), Unmarshall, NewEncoder(), NewDecoder()

	// example todo
	todo := Todo{
		Id:        1,
		Title:     "Listen to The Weeknd",
		Completed: true,
	}

	// 1. Marshall()
	data, err := json.Marshal(todo) // returns []byte
	if err != nil {
		fmt.Println("Error while encoding todo")
		return
	}
	fmt.Println(string(data))
	fmt.Println(data) // byte data as a list

	// 2. Unmarshall()
	var deTodo Todo
	er := json.Unmarshal(data, &deTodo)
	if er != nil {
		fmt.Println("Error while decoding")
		return
	}
	fmt.Println(deTodo)
	fmt.Println("Id: ", deTodo.Id)
	fmt.Println("Title: ", deTodo.Title)
	fmt.Println("Completed: ", deTodo.Completed)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		decoder := json.NewDecoder(r.Body) // this is the intialization of the decoder
		decoder.DisallowUnknownFields()    // will produce an error when unknown fields are sent
		err := decoder.Decode(&todo)       // this is the process of decoding
		if err != nil {
			fmt.Println()
		}
	})
	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		todo := Todo{
			Id:        2,
			Title:     "Grind The Weeknd",
			Completed: true,
		}
		encoder.Encode(todo)

	})

	// post endpoint
	mux.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		// json to struct (object)
		var todo Todo
		dc := json.NewDecoder(r.Body)
		dc.DisallowUnknownFields()
		err := dc.Decode(&todo)
		if err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		fmt.Println("Success Decoding")
	})

	fmt.Println("The Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)

	// fmt.Println("Hello World")

	// mux := http.NewServeMux()

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello The Weeknd")
	// })

	// fmt.Println("The server is running on the porg 8080")
	// http.ListenAndServe(":8080", mux)

	// c := chi.NewRouter()

	// c.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello The Weeknd")
	// })

	// fmt.Println("The server is running on the porg 8080")
	// http.ListenAndServe(":8080", c)
}

// =============================================== DAY 23 ======================================================

// the todo project

// type Todo struct {
// 	todo string
// }

// var TodoList = []Todo{}

// func main() {
// 	c := chi.NewRouter()

// the business logic here
// add
// get all

// 	c.Post("/add", func(w http.ResponseWriter, r *http.Request) {
// 		todo := r.URL.Query().Get("todo")
// 		newtodo := Todo{
// 			todo: todo,
// 		}
// 		TodoList = append(TodoList, newtodo)
// 		fmt.Fprintln(w, "Todo added successfully")
// 	})

// 	c.Get("/alltodo", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, TodoList)
// 	})

// 	fmt.Println("The server is running on the port 8080")
// 	http.ListenAndServe(":8080", c)
// }

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
