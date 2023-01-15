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

func (o *Order) NewOrder(mid int, address string) (bool, error) {
	t := time.Now()
	fmt.Println(t, mid, address)
	res, err := MysqlDb.Exec("INSERT INTO shop.order (time,state,mid,address) VALUES (?,1,?,?)", t, mid, address)
	if err != nil {
		return false, err
	} else {
		i, _ := res.RowsAffected()
		if i == 0 {
			return false, nil
		}
		return true, nil
	}
}
