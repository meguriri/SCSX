package data

import "fmt"

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Fid  int    `json:"fid"`
}

func GetAllCategory() []Category {
	list := make([]Category, 0)
	rows, _ := MysqlDb.Query("SELECT * FROM category")
	for rows.Next() {
		var s Category
		rows.Scan(&s.Id, &s.Name, &s.Fid)
		fmt.Println(s)
		list = append(list, s)
	}
	return list
}
