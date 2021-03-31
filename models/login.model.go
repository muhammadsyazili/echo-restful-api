package models

import "github.com/muhammadsyazili/echo-rest/db"

type Login struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckLogin(username string, password string) (Login, error) {
	var obj Login

	conn := db.OpenConn()
	
	sqlQuery := "SELECT * FROM users WHERE username = ?"

	q, err := conn.Prepare(sqlQuery)
    if err != nil {
		return obj, err
    }
	defer q.Close()

	q.QueryRow(username).Scan(&obj.Id, &obj.Username, &obj.Password)

	//err := conn.QueryRow(sqlQuery, username).Scan(&obj.Id, &obj.Username, &obj.Password)

	if err != nil {
		return obj, err
	}

	return obj, nil
}