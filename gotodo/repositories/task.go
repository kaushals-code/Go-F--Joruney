package repositories

import (
	"database/sql"
	"gotodo/config"
)

type TaskRepository struct {
	DB *sql.DB
}

func getDBInstance() (TaskRepository, error) {
	inst, err := config.ConnectDB()

	tr := TaskRepository{}

	if err != nil {
		return tr, err
	}

	tr.DB = inst

	return tr, nil
}

///////////////////////////

// type User struct {
// 	name string
// 	age int
// }

// func main() {
// 	user := User{}
// 	user.name = "kaushal"
// 	user.age = 19

// 	fmt.Println(user)
// }
