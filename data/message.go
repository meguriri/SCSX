package data

import (
	"time"
)

type Message struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
	Mid     int       `json:"mid"`
}

func (m *Message) AddMessage() (bool, error) {
	res, err := MysqlDb.Exec("INSERT INTO shop.message (title,content,date,mid) VALUES (?,?,?,?)", m.Title, m.Content, m.Date, m.Mid)
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
