package models

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/muhammadsyazili/echo-rest/db"
	"github.com/muhammadsyazili/echo-rest/template"
)

type Transaction struct {
	Id int `json:"id"`
	User_id int `json:"user_id" validate:"required,number,max=20"`
	Title string `json:"title" validate:"required,max=255"`
	Amount float64 `json:"amount" validate:"required,numeric"`
	Time string `json:"time" validate:"required"`
	Type string `json:"type" validate:"required,oneof=expense revenue"`
	Created_at string `json:"created_at" validate:"required"`
	Updated_at string `json:"updated_at" validate:"required"`
}

func GetAllTransaction() (template.Response, error) {
	var obj Transaction
	var arrobj []Transaction
	var res template.Response
	
	conn := db.CreateConn()

	sqlQuery := "SELECT * FROM transactions"

	rows, err := conn.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.User_id, &obj.Title, &obj.Amount, &obj.Time, &obj.Type, &obj.Created_at, &obj.Updated_at)
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

func GetWhereTransaction(Id int) (template.Response, error) {
	var obj Transaction
	var res template.Response
	
	conn := db.CreateConn()
	defer conn.Close()
	
	sqlQuery := "SELECT * FROM transactions WHERE id = ?"

	err := conn.QueryRow(sqlQuery, Id).Scan(&obj.Id, &obj.User_id, &obj.Title, &obj.Amount, &obj.Time, &obj.Type, &obj.Created_at, &obj.Updated_at)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Ok"
	res.Data = obj

	return res, nil
}

func StoreTransaction(User_id int, Title string, Amount float64, Time string, Type string) (template.Response, error) {
	var obj Transaction
	var res template.Response

	//validation input
	v := validator.New()

	data := Transaction{
		User_id: User_id,
		Title: Title,
		Amount: Amount,
		Time: Time,
		Type: Type,
	}

	err := v.Struct(data)
	if err != nil {
		return res, err
	}

	conn := db.CreateConn()

	sqlQuery := "INSERT transactions (user_id, title, amount, time, type, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	q, err := conn.Prepare(sqlQuery)
	defer q.Close()
	if err != nil {
		return res, err
	}

	result, err := q.Exec(User_id, Title, Amount, Time, Type, template.Timestamp, template.Timestamp)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	//--------------------------------------------------------------------------------------

	conn = db.CreateConn()
	defer conn.Close()
	
	sqlQuery = "SELECT * FROM transactions WHERE id = ?"

	err = conn.QueryRow(sqlQuery, int(lastInsertId)).Scan(&obj.Id, &obj.User_id, &obj.Title, &obj.Amount, &obj.Time, &obj.Type, &obj.Created_at, &obj.Updated_at)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Created"
	res.Data = obj

	return res, nil
}

func UpdateTransaction(Id int, User_id int, Title string, Amount float64, Time string, Type string) (template.Response, error) {
	var obj Transaction
	var res template.Response

	//validation input
	v := validator.New()

	data := Transaction{
		User_id: User_id,
		Title: Title,
		Amount: Amount,
		Time: Time,
		Type: Type,
	}

	err := v.Struct(data)
	if err != nil {
		return res, err
	}

	conn := db.CreateConn()

	sqlQuery := "UPDATE transactions SET user_id = ?, title = ?, amount = ?, time = ?, type = ?, updated_at = ? WHERE id = ?"

	q, err := conn.Prepare(sqlQuery)
	defer q.Close()
	if err != nil {
		return res, err
	}

	result, err := q.Exec(User_id, Title, Amount, Time, Type, template.Timestamp, Id)
	if err != nil {
		return res, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	//--------------------------------------------------------------------------------------

	conn = db.CreateConn()
	defer conn.Close()
	
	sqlQuery = "SELECT * FROM transactions WHERE id = ?"

	err = conn.QueryRow(sqlQuery, Id).Scan(&obj.Id, &obj.User_id, &obj.Title, &obj.Amount, &obj.Time, &obj.Type, &obj.Created_at, &obj.Updated_at)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Updated"
	res.Data = obj

	return res, nil
}

func DestroyTransaction(Id int) (template.Response, error) {
	var res template.Response

	conn := db.CreateConn()

	sqlQuery := "DELETE FROM transactions WHERE id = ?"

	q, err := conn.Prepare(sqlQuery)
	defer q.Close()
	if err != nil {
		return res, err
	}

	result, err := q.Exec(Id)
	if err != nil {
		return res, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}