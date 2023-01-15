package data

import (
	"fmt"
	"time"
)

type Order struct {
	Id      int       `json:"id"`
	Time    time.Time `json:"time"`
	State   int       `json:"state"` //1已付款，2已发货，3已收货
	Mid     int       `json:"mid"`
	Address string    `json:"address"`
}

func (o *Order) NewOrder(mid int, address string) (int, bool, error) {
	t := time.Now()
	fmt.Println(t, mid, address)
	res, err := MysqlDb.Exec("INSERT INTO shop.order (time,state,mid,address) VALUES (?,1,?,?)", t, mid, address)
	if err != nil {
		return -1, false, err
	} else {
		i, _ := res.LastInsertId()
		ok := PayAllBuyCar(mid, int(i))
		fmt.Println("i:", i)
		if i == 0 || !ok {
			return -1, false, nil
		}
		return int(i), true, nil
	}
}

func PayAllBuyCar(mid, oid int) bool {
	list := GetAllProduct(mid)
	fmt.Println(list)
	ok := InsertOrderProduct(list, oid)
	if !ok {
		return false
	}
	_, err := MysqlDb.Exec("delete from buycar where mid=?", mid)
	if err != nil {
		fmt.Println("delete failed", err)
		return false
	} else {
		fmt.Println("delete success", err)
		return true
	}
}

func InsertOrderProduct(list []BuyCar, oid int) bool {
	for _, v := range list {
		res, err := MysqlDb.Exec("INSERT INTO shop.orderdetail (pid,oid,num) VALUES (?,?,?)", v.Pid, oid, v.Num)
		if err != nil {
			fmt.Println(err)
			return false
		} else {
			resS, err := MysqlDb.Exec("UPDATE product set number=number-? where id=?", v.Num, v.Pid)
			if err != nil {
				fmt.Println(err)
			} else {
				id, _ := res.LastInsertId()
				eff, _ := resS.RowsAffected()
				fmt.Println(id, eff)
			}
		}
	}
	return true
}

func GetMyOrderList(mid int) ([]Order, bool) {
	orders := make([]Order, 0)
	rows, err := MysqlDb.Query("select id,time,state,mid,address from shop.order where mid =?;", mid)
	if err != nil {
		fmt.Println("get all my order list err:", err)
		return orders, false
	}
	var order Order
	for rows.Next() {
		rows.Scan(&order.Id, &order.Time, &order.State, &order.Mid, &order.Address)
		str := ""
		if order.State == 1 {
			str = "已付款，未发货，未收货"
		} else if order.State == 2 {
			str = "已付款，已发货，未收货"
		} else if order.State == 3 {
			str = "已付款，已发货，已收货"
		}
		order.Address = str
		orders = append(orders, order)
	}
	fmt.Println(orders)
	return orders, true
}
