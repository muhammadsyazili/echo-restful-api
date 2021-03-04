package models

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"

	"github.com/muhammadsyazili/echo-rest/db"
	"github.com/muhammadsyazili/echo-rest/helpers"
)

type Account struct {
	Id         int    	`json:"id"`
	Username   string 	`json:"username" validate:"required,unique,min=5,max=50"`
	Password   string 	`json:"password" validate:"required,min=5"`
}

func GetAllAccount() (Response, error) {
	var obj Account
	var arrobj []Account
	var res Response
	
	conn := db.CreateConn()

	sqlQuery := "SELECT * FROM accounts"

	rows, err := conn.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Password)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = arrobj

	return res, nil
}

func GetWhereAccount(id int) (Response, error) {
	var obj Account
	var res Response
	
	conn := db.CreateConn()
	defer conn.Close()
	
	sqlQuery := "SELECT * FROM accounts WHERE id = ?"

	err := conn.QueryRow(sqlQuery, id).Scan(&obj.Id, &obj.Username, &obj.Password)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = obj

	return res, nil
}

func StoreAccount(username string, password string) (Response, error) {
	var res Response

	//validation input
	v := validator.New()

	data := Account{
		Username: username,
		Password: password,
	}

	err := v.Struct(data)
	if err != nil {
		return res, err
	}

	//hashing password
	password_hash, err := helpers.Hash(password)
	if err != nil {
		return res, err
	}

	conn := db.CreateConn()

	sqlQuery := "INSERT accounts (username, password) VALUES (?, ?)"

	q, err := conn.Prepare(sqlQuery)
	defer q.Close()
	if err != nil {
		return res, err
	}

	result, err := q.Exec(username, password_hash)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	data.Id = int(lastInsertId)
	data.Username = username
	data.Password = password_hash

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = data

	return res, nil
}

func UpdateAccount(id int, username string, password string) (Response, error) {
	var res Response
	var err error

	//validation input
	v := validator.New()

	data := Account{
		Username: username,
		Password: password,
	}

	err = v.Struct(data)
	if err != nil {
		return res, err
	}

	//hashing password
	password_hash, err := helpers.Hash(password)
	if err != nil {
		return res, err
	}

	conn := db.CreateConn()

	sqlQuery := "UPDATE accounts SET username = ?, password = ? WHERE id = ?"

	q, err := conn.Prepare(sqlQuery)
	defer q.Close()
	if err != nil {
		return res, err
	}

	_, err = q.Exec(username, password_hash, id)
	if err != nil {
		return res, err
	}

	data.Id = id
	data.Username = username
	data.Password = password_hash

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res, nil
}

func DestroyAccount(id int) (Response, error) {
	var res Response

	conn := db.CreateConn()

	sqlQuery := "DELETE FROM accounts WHERE id = ?"

	q, err := conn.Prepare(sqlQuery)
	defer q.Close()
	if err != nil {
		return res, err
	}

	result, err := q.Exec(id)
	if err != nil {
		return res, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "ok"

	return res, nil
}