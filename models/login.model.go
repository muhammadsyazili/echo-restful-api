package models

import (
	"database/sql"
	"fmt"

	"github.com/muhammadsyazili/echo-rest/db"
	"github.com/muhammadsyazili/echo-rest/helpers"
)


type Login struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Student_id string `json:"student_id"`
}

func CheckLogin(username string, password string) (bool, error) {
	var obj Login

	conn := db.CreateConn()
	defer conn.Close()
	
	sqlQuery := "SELECT * FROM accounts WHERE username = ?"

	err := conn.QueryRow(sqlQuery, username).Scan(&obj.Id, &obj.Username, &obj.Password, &obj.Student_id)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found!")
		return false, err
	}

	if err != nil {
		return false, err
	}

	match, err := helpers.CheckHash(password, obj.Password)
	if !match {
		fmt.Println(obj.Password)
		fmt.Println("Password and hash does't match!")
		return false, err
	}

	return true, nil
}