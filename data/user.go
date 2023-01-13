package data

import "fmt"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	State    int    `json:"state"`
}

func (u *User) Exist() bool {
	row := MysqlDb.QueryRow("SELECT username FROM member WHERE username=?", u.Username)
	var name string
	row.Scan(&name)
	if name == "" {
		return false
	}
	return true
}
func (u *User) RegisterUser() (bool, error) {
	if u.Exist() {
		return false, nil
	}
	res, err := MysqlDb.Exec("INSERT INTO member(username,password,email,state)VALUES(?,?,?,0)", u.Username, u.Password, u.Email)
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

func CheckUser(username, password string) (bool, User) {
	row := MysqlDb.QueryRow("SELECT id,username,password,email,state FROM member WHERE username=? AND password=?", username, password)
	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.State)
	if err != nil {
		fmt.Println(err)
		return false, user
	}
	fmt.Println("user: ", user)
	if user.Username == "" {
		return false, user
	}
	return true, user
}
