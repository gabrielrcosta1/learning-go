package models

import (
	"github.com/gabrielrcosta1/learning-go/api/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	slq := `INSERT INTO todos (title,description,done) VALUES ($1,$2,$3) RETURNING id`
	err = conn.QueryRow(slq, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
