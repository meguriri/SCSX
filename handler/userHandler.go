package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meguriri/SCSX/data"
	"net/http"
)

func GetIndexHtml() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取login cookie中的sid
		sid, err := c.Cookie("login")
		//无cookie
		if err != nil {
			fmt.Println("no cookie")
			//登录
			c.Redirect(http.StatusFound, "/user/login")
		} else { //有cookie
			//查询sid
			ok := data.Exist(sid)
			if ok {
				//返回index.html
				c.HTML(http.StatusOK, "index.html", nil)
			} else {
				//登录
				c.Redirect(http.StatusFound, "/user/login")
			}
		}
	}
}
func GetRegisterHtml() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "userRegister.html", nil)
	}
}

func UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := data.User{
			Username: c.PostForm("username"),
			Password: c.PostForm("password"),
			Email:    c.PostForm("email"),
			State:    0,
		}
		fmt.Println(user)
		if ok, err := user.RegisterUser(); !ok {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"msg": "fail",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	}
}

func GetLoginHtml() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "userLogin.html", nil)
	}
}

func UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password := c.PostForm("username"), c.PostForm("password")
		fmt.Println(username, password)
		ok, user := data.CheckUser(username, password)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": "fail",
			})
			return
		}
		s, _ := json.Marshal(user)
		fmt.Println("user json:", string(s))
		session := data.Session{
			Message:     string(s),
			MaxLifetime: 600000000000,
		}
		sid, err := session.Set()
		if err != nil {
			fmt.Println(sid)
			c.JSON(http.StatusOK, gin.H{
				"msg": "fail",
			})
			return
		}
		c.SetCookie("login", sid, int(session.MaxLifetime/1e9), "/", "localhost", false, false)
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	}
}

func UserSHow() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie("login")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":      "fail",
				"username": "",
			})
		} else {
			s, err := data.Get(sid)
			var user data.User
			json.Unmarshal([]byte(s), &user)
			fmt.Println("unmarshal:", user)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"msg":      "fail",
					"username": "",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg":      "success",
					"username": user.Username,
				})
			}
		}
	}
}

func Exit() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie("login")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			err := data.Del(sid)
			c.SetCookie("login", "", -1, "/", "localhost", false, false)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": "success",
				})
			}
		}
	}
}
