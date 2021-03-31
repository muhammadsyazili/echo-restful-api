package models

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/muhammadsyazili/echo-rest/db"
	"github.com/muhammadsyazili/echo-rest/helpers"
	"github.com/muhammadsyazili/echo-rest/template"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name" validate:"required,max=255"`
	Username string `json:"username" validate:"required,max=255"`
	Email string `json:"email" validate:"required,max=255"`
	Password *string `json:"password" validate:"required,max=255"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func GetAllUser() (template.Response, error) {
	var obj User
	var arrobj []User
	var res template.Response
	
	conn := db.OpenConn()

	sqlQuery := "SELECT id, name, username, email, created_at, updated_at FROM users"

	rows, err := conn.Query(sqlQuery)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Username, &obj.Email, &obj.Created_at, &obj.Updated_at)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Ok"
	res.Data = arrobj

	return res, nil
}

func GetWhereUser(Id int) (template.Response, error) {
	var obj User
	var res template.Response
	
	conn := db.OpenConn()
	
	sqlQuery := "SELECT id, name, username, email, created_at, updated_at FROM users WHERE id = ?"

	q, err := conn.Prepare(sqlQuery)
    if err != nil {
		return res, err
    }
	defer q.Close()

	q.QueryRow(Id).Scan(&obj.Id, &obj.Name, &obj.Username, &obj.Email, &obj.Created_at, &obj.Updated_at)

	// err := conn.QueryRow(sqlQuery, Id).Scan(&obj.Id, &obj.Name, &obj.Username, &obj.Email, &obj.Created_at, &obj.Updated_at)
	// if err != nil {
	// 	return res, err
	// }

	res.Status = http.StatusOK
	res.Message = "Ok"
	res.Data = obj

	return res, nil
}

func StoreUser(Name string, Username string, Email string, Password string) (template.Response, error) {
	var obj User
	var res template.Response

	//validation input
	v := validator.New()

	data := User{
		Name: Name,
		Username: Username,
		Email: Email,
		Password: &Password,
	}

	err := v.Struct(data)
	if err != nil {
		return res, err
	}

	//hashing password
	Password_hash, err := helpers.Hash(Password)
	if err != nil {
		return res, err
	}

	conn := db.OpenConn()

	sqlQuery := "INSERT users (name, username, email, email_verified_at, password, remember_token, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	q, err := conn.Prepare(sqlQuery)
	if err != nil {
		return res, err
	}
	defer q.Close()

	result, err := q.Exec(Name, Username, Email, nil, Password_hash, nil, template.Timestamp, template.Timestamp)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	//--------------------------------------------------------------------------------------
	
	sqlQuery = "SELECT id, name, username, email, created_at, updated_at FROM users WHERE id = ?"

	q, err = conn.Prepare(sqlQuery)
    if err != nil {
		return res, err
    }
	defer q.Close()

	q.QueryRow(int(lastInsertId)).Scan(&obj.Id, &obj.Name, &obj.Username, &obj.Email, &obj.Created_at, &obj.Updated_at)

	res.Status = http.StatusOK
	res.Message = "Created"
	res.Data = obj

	return res, nil
}

func UpdateUser(Id int, Name string, Username string, Email string, Password string) (template.Response, error) {
	var obj User
	var res template.Response
	var err error

	//validation input
	v := validator.New()

	data := User{
		Name: Name,
		Username: Username,
		Email: Email,
		Password: &Password,
	}

	err = v.Struct(data)
	if err != nil {
		return res, err
	}

	//hashing password
	Password_hash, err := helpers.Hash(Password)
	if err != nil {
		return res, err
	}

	conn := db.OpenConn()

	sqlQuery := "UPDATE users SET name = ?, username = ?, email = ?, password = ?, updated_at = ? WHERE id = ?"

	q, err := conn.Prepare(sqlQuery)
	if err != nil {
		return res, err
	}
	defer q.Close()

	_, err = q.Exec(Name, Username, Email, Password_hash, template.Timestamp, Id)
	if err != nil {
		return res, err
	}

	//--------------------------------------------------------------------------------------
	
	sqlQuery = "SELECT id, name, username, email, created_at, updated_at FROM users WHERE id = ?"

	q, err = conn.Prepare(sqlQuery)
    if err != nil {
		return res, err
    }
	defer q.Close()

	q.QueryRow(Id).Scan(&obj.Id, &obj.Name, &obj.Username, &obj.Email, &obj.Created_at, &obj.Updated_at)

	res.Status = http.StatusOK
	res.Message = "Updated"
	res.Data = obj

	return res, nil
}

func DestroyUser(Id int) (template.Response, error) {
	var res template.Response

	conn := db.OpenConn()

	sqlQuery := "DELETE FROM users WHERE id = ?"

	q, err := conn.Prepare(sqlQuery)
	if err != nil {
		return res, err
	}
	defer q.Close()

	result, err := q.Exec(Id)
	if err != nil {
		return res, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Deleted"

	return res, nil
}