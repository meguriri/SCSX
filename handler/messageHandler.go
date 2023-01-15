package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meguriri/SCSX/data"
	"net/http"
	"time"
)

func GetMessageHtml() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "message.html", nil)
	}
}

func GetMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie("login")
		if err != nil || !data.Exist(sid) {
			fmt.Println("fail or not exit")
			c.JSON(http.StatusOK, gin.H{
				"msg": "fail",
			})
		} else {
			s, _ := data.Get(sid)
			var user data.User
			json.Unmarshal([]byte(s), &user)
			msg := data.Message{
				Title:   c.PostForm("title"),
				Content: c.PostForm("content"),
				Date:    time.Now(),
				Mid:     user.Id,
			}
			ok, err := msg.AddMessage()
			if !ok {
				fmt.Println(err)
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
