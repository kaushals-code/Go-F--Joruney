package main

import (
	"fmt"
	"net/http"
)

// func handleReq(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Method: ", r.Method)
// }

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Bro!")
}

// creating my own handler
type HelloWeekndHandler struct{}

// func (h HelloWeekndHandler) ServeHTTP(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {
// 	fmt.Fprintln(w, "Hello The Weeknd")
// }

func (h HelloWeekndHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintln(w, "Hello The Weeknd Bro, big fan bro")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Weeknd from handleHello the new function")
}

func main() {

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
	mux := http.NewServeMux()
	mux.HandleFunc("/helloweeknd", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handleHello(w, r)
		} else {
			fmt.Fprintln(w, "You cannot use the POST method for this endpoint")
		}
	})
	fmt.Println("The Server is running on the port 8080")
	http.ListenAndServe(":8080", mux)

}
