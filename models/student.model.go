package models

import (
	"fmt"
	"net/http"

	validator "github.com/go-playground/validator/v10"

	"github.com/muhammadsyazili/echo-rest/db"
)

type Student struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama" validate:"required"`
	NIM     int `json:"nim" validate:"required,numeric,len=14"`
	Jurusan string `json:"jurusan" validate:"required"`
}

func GetAllStudent() (Response, error) {
	fmt.Println("sampe model")
	var obj Student
	var arrobj []Student
	var res Response
	
	conn := db.CreateConn()

	sqlQuery := "SELECT * FROM students"

	rows, err := conn.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.NIM, &obj.Jurusan)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func GetWhereStudent(id int) (Response, error) {
	var obj Student
	var res Response
	
	conn := db.CreateConn()
	defer conn.Close()
	
	sqlQuery := "SELECT * FROM students WHERE id = ?"

	err := conn.QueryRow(sqlQuery, id).Scan(&obj.Id, &obj.Nama, &obj.NIM, &obj.Jurusan)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func StoreStudent(nama string, nim int, jurusan string) (Response, error) {
	var res Response

	//validation input
	v := validator.New()

	data := Student{
		Nama: nama,
		NIM: nim,
		Jurusan: jurusan,
	}

	err := v.Struct(data)
	if err != nil {
		return res, err
	}

	conn := db.CreateConn()

	sqlQuery := "INSERT students (nama, nim, jurusan) VALUES (?, ?, ?)"

	q, err := conn.Prepare(sqlQuery)
	defer q.Close()
	if err != nil {
		return res, err
	}

	result, err := q.Exec(nama, nim, jurusan)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_insert_id": lastInsertId,
	}

	return res, nil
}

func UpdateStudent(id int, nama string, nim int, jurusan string) (Response, error) {
	var res Response

	//validation input
	v := validator.New()

	data := Student{
		Nama: nama,
		NIM: nim,
		Jurusan: jurusan,
	}

	err := v.Struct(data)
	if err != nil {
		return res, err
	}

	conn := db.CreateConn()

	sqlQuery := "UPDATE students SET nama = ?, nim = ?, jurusan = ? WHERE id = ?"

	q, err := conn.Prepare(sqlQuery)
	defer q.Close()
	if err != nil {
		return res, err
	}

	result, err := q.Exec(nama, nim, jurusan, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DestroyStudent(id int) (Response, error) {
	var res Response

	conn := db.CreateConn()

	sqlQuery := "DELETE FROM students WHERE id = ?"

	q, err := conn.Prepare(sqlQuery)
	defer q.Close()
	if err != nil {
		return res, err
	}

	result, err := q.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}