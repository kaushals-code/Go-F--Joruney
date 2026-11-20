package main

import (
	"fmt"
	// "gotodo/config"
	"gotodo/repositories"
	// "log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "THE WEEKND is the GOAT bro")
	})

	// database - connection
	// db, err := config.ConnectDB()

	// if err != nil {
	// 	log.Fatal("Error : ", err)
	// 	db.Close()
	// }

	mux.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		// if db == nil {
		// 	fmt.Fprintln(w, "Error: ", db)
		// 	db.Close()
		// 	return
		// }

		// datab := repositories.TaskRepository{
		// 	DB: db,
		// }

		// fmt.Fprintln(w, "Database connected successfully", datab)

		tr := repositories.TaskRepository{}

		tr.DB, err := repositories.getDBInstance()

		if err != nil {
			Fprintln(w, "Error" , err)
		}

		Fprintln(w, "Connected successfully")

	})

	http.ListenAndServe(":8000", mux)

}
