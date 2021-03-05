package models

import "github.com/muhammadsyazili/echo-rest/db"

type Login struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckLogin(username string, password string) (Login, error) {
	var obj Login

	conn := db.CreateConn()
	defer conn.Close()
	
	sqlQuery := "SELECT * FROM accounts WHERE username = ?"

	err := conn.QueryRow(sqlQuery, username).Scan(&obj.Id, &obj.Username, &obj.Password)

	if err != nil {
		return obj, err
	}

	return obj, nil
}