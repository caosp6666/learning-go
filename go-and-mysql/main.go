package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Person model
type Person struct {
	UserID   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

// Db by sqlx
var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "turing:turing0803@tcp(39.108.55.149:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	fmt.Println("connect to database success!")
	Db = database
	// defer Db.Close() // 注意这行代码要写在上面err判断的下面
}

func main() {
	var people []Person
	err := Db.Select(&people, "select * from people")
	if err != nil {
		fmt.Println("select failed，", err)
	} else {
		fmt.Println("select success:", people)
	}

	r, err := Db.Exec("insert into people(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert success:", id)
}
