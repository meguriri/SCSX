package data

import "fmt"

type BuyCar struct {
	Id  int `json:"id"`
	Pid int `json:"pid"`
	Num int `json:"num"`
	Mid int `json:"mid"`
}

type AllBuyCar struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Color   string  `json:"color"`
	Number  int     `json:"number"`
	Price   float64 `json:"price"`
	ImgPath string  `json:"imgPath"`
}

func (b *BuyCar) Add() (bool, error) {
	res, err := MysqlDb.Exec("INSERT INTO buycar(pid,number,mid)VALUES(?,?,?)", b.Pid, b.Num, b.Mid)
	if err != nil {
		return false, err
	} else {
		i, err := res.RowsAffected()
		if err != nil {
			return false, err
		}
		fmt.Println(i)
		return true, nil
	}
}

func (b *BuyCar) Exist() (bool, error) {
	row := MysqlDb.QueryRow("SELECT * FROM buycar WHERE pid=? AND mid=?", b.Pid, b.Mid)
	var car BuyCar
	err := row.Scan(&car.Id, &car.Pid, &car.Num, &car.Mid)
	fmt.Println("car:", car)
	if err != nil || car.Id == 0 {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

func (b *BuyCar) Update() (bool, error) {
	_, err := MysqlDb.Exec("update buycar set number=number+? where pid=? AND mid=?", b.Num, b.Pid, b.Mid)
	if err != nil {
		fmt.Println("update failed", err)
		return false, err
	} else {
		fmt.Println("update success", err)
		return true, nil
	}
}

func (b *BuyCar) GetInfo() (tType int, total int, price float64) {
	row := MysqlDb.QueryRow("select count(*),sum(b.number),sum(b.number*p.price) from shop.buycar b,shop.product p where b.pid=p.id and b.mid=?;", b.Mid)
	err := row.Scan(&tType, &total, &price)
	if err != nil {
		fmt.Println("get info error:", err)
		return -1, -1, -1
	}
	fmt.Println(tType, total, price)
	return
}

func GetAllProducts(mid int) ([]AllBuyCar, float64) {
	cars, totalPrice := make([]AllBuyCar, 0), 0.0
	rows, _ := MysqlDb.Query("select b.id,p.name,p.color,b.number,p.price*b.number,p.imgPath,p.img from shop.buycar b ,shop.product p where b.pid = p.id and b.mid =?;", mid)
	var car AllBuyCar
	for rows.Next() {
		img, imgP := "", ""
		rows.Scan(&car.Id, &car.Name, &car.Color, &car.Number, &car.Price, &imgP, &img)
		car.ImgPath = imgP + img
		totalPrice += car.Price
		cars = append(cars, car)
	}
	fmt.Println(cars)
	fmt.Println(totalPrice)
	return cars, totalPrice
}

func GetAllProduct(mid int) []BuyCar {
	cars := make([]BuyCar, 0)
	rows, err := MysqlDb.Query("select * from buycar where mid =?;", mid)
	if err != nil {
		fmt.Println("get all product err", err)
		return cars
	}
	var car BuyCar
	for rows.Next() {
		rows.Scan(&car.Id, &car.Pid, &car.Num, &car.Mid)
		cars = append(cars, car)
	}
	fmt.Println(cars)
	return cars
}

func DelBuyCar(id int) bool {
	_, err := MysqlDb.Exec("delete from buycar where id=?", id)
	if err != nil {
		fmt.Println("delete failed", err)
		return false
	} else {
		fmt.Println("delete success", err)
		return true
	}
}

func DelAllBuyCar(mid int) bool {
	_, err := MysqlDb.Exec("delete from buycar where mid=?", mid)
	if err != nil {
		fmt.Println("delete failed", err)
		return false
	} else {
		fmt.Println("delete success", err)
		return true
	}
}
