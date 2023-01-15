package data

import "fmt"

type Product struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Color   string  `json:"color"`
	Price   float64 `json:"price"`
	State   int     `json:"state"`
	Number  int     `json:"number"`
	Img     string  `json:"img"`
	ImgPath string  `json:"imgPath"`
	Cid     int     `json:"cid"`
}

func GetProductsByCid(cid int) []Product {
	list := make([]Product, 0)
	rows, _ := MysqlDb.Query("SELECT id,name,type,color,price,state,number,img,imgPath,cid FROM product WHERE cid=?", cid)
	for rows.Next() {
		var s Product
		rows.Scan(&s.Id, &s.Name, &s.Type, &s.Color, &s.Price, &s.State, &s.Number, &s.Img, &s.ImgPath, &s.Cid)
		fmt.Println("product:", s)
		list = append(list, s)
	}
	return list
}

func ProductDetail(id int) (s Product) {
	row := MysqlDb.QueryRow("SELECT id,name,type,color,price,state,number,img,imgPath,cid FROM product WHERE id=?", id)
	row.Scan(&s.Id, &s.Name, &s.Type, &s.Color, &s.Price, &s.State, &s.Number, &s.Img, &s.ImgPath, &s.Cid)
	fmt.Println("s:", s)
	return
}

func GetProductsByName(name string) []Product {
	list := make([]Product, 0)
	fmt.Println("sql name:", name)
	rows, err := MysqlDb.Query("SELECT id,name,type,color,price,state,number,img,imgPath,cid FROM product WHERE name like ?", "%"+name+"%")
	if err != nil {
		fmt.Println("by name err:", err)
		return list
	}
	for rows.Next() {
		var s Product
		rows.Scan(&s.Id, &s.Name, &s.Type, &s.Color, &s.Price, &s.State, &s.Number, &s.Img, &s.ImgPath, &s.Cid)
		fmt.Println("product:", s)
		list = append(list, s)
	}
	return list
}
