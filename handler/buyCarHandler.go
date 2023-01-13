package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meguriri/SCSX/data"
	"net/http"
	"strconv"
)

func GetBuyCarHtml() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "buycar.html", nil)
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie("login")
		if err != nil || !data.Exist(sid) {
			fmt.Println("fail or not exit")
			c.JSON(http.StatusOK, gin.H{
				"msg":   "fail",
				"list":  nil,
				"total": -1,
			})
		} else {
			s, _ := data.Get(sid)
			var user data.User
			json.Unmarshal([]byte(s), &user)
			list, total := data.GetAllProducts(user.Id)
			l, _ := json.Marshal(list)
			c.JSON(http.StatusOK, gin.H{
				"msg":   "ok",
				"list":  string(l),
				"total": total,
			})
		}
	}
}

func DelBuyCar() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.PostForm("id")
		i, _ := strconv.Atoi(id)
		if ok := data.DelBuyCar(i); !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": "fail",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		}
	}
}

func DelAllBuyCar() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie("login")
		if err != nil || !data.Exist(sid) {
			fmt.Println("fail or not exit")
			c.JSON(http.StatusOK, gin.H{
				"msg":   "fail",
				"list":  nil,
				"total": -1,
			})
		} else {
			s, _ := data.Get(sid)
			var user data.User
			json.Unmarshal([]byte(s), &user)
			if ok := data.DelAllBuyCar(user.Id); !ok {
				c.JSON(http.StatusOK, gin.H{
					"msg": "fail",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": "ok",
				})
			}
		}
	}
}
