package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	MysqlDb    *sql.DB
	MysqlDbErr error
)

func InitDB() {
	MysqlDb, MysqlDbErr = sql.Open("mysql", "root"+":"+"xyy001019"+"@tcp("+"localhost"+":"+"3306"+")/"+"shop"+"?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	if MysqlDbErr != nil {
		panic(MysqlDbErr.Error())
	} else {
		fmt.Println("connect success!!")
	}
}
