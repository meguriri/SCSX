package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meguriri/SCSX/data"
	"net/http"
)

func GetConfirmOrderHtml() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "confirmOrder.html", nil)
	}
}

func ConfirmOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie("login")
		if err != nil || !data.Exist(sid) {
			fmt.Println("fail or not exit")
			c.JSON(http.StatusOK, gin.H{
				"msg": "fail",
				"num": -1,
			})
		} else {
			person, tel, address := c.PostForm("person"), c.PostForm("tel"), c.PostForm("address")
			s, _ := data.Get(sid)
			var user data.User
			json.Unmarshal([]byte(s), &user)
			var order data.Order
			ok, err := order.NewOrder(user.Id, person+","+tel+","+address)
			if !ok {
				fmt.Println("order creat err:", err)
				c.JSON(http.StatusOK, gin.H{
					"msg": "fail",
					"num": -1,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": "ok",
					"num": 100,
				})
			}
		}
	}
}
