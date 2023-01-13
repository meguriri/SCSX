package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meguriri/SCSX/data"
	"net/http"
	"strconv"
)

func GetAllCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryList := data.GetAllCategory()
		fmt.Println(categoryList)
		c.JSON(http.StatusOK, gin.H{
			"list": categoryList,
		})
	}
}
func GetProductsHtml() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "selById.html", nil)
	}
}
func GetProductByCid() gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := c.Query("cid")
		cid, _ := strconv.Atoi(cc)
		fmt.Println("cid:", cid)
		list := data.GetProductsByCid(cid)
		str, err := json.Marshal(&list)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "fail",
				"list": nil,
			})
		} else {
			fmt.Println(string(str))
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"list": string(str),
			})
		}
	}
}

func GetProductHtml() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "productInfo.html", nil)
	}
}
func GetProductDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		fmt.Println("id:", id)
		product := data.ProductDetail(id)
		str, err := json.Marshal(&product)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "fail",
				"list": nil,
			})
		} else {
			fmt.Println(string(str))
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"list": string(str),
			})
		}
	}
}

func AddCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, err := c.Cookie("login")
		if err != nil && data.Exist(sid) {
			fmt.Println("fail or not exit")
			c.JSON(http.StatusOK, gin.H{
				"msg":   "fail",
				"type":  0,
				"total": 0,
				"price": 0,
			})
		} else {
			s, _ := data.Get(sid)
			var user data.User
			json.Unmarshal([]byte(s), &user)
			sPid, sNum := c.PostForm("pid"), c.PostForm("num")
			pid, _ := strconv.Atoi(sPid)
			num, _ := strconv.Atoi(sNum)
			bCar := data.BuyCar{
				Pid: pid,
				Num: num,
				Mid: user.Id,
			}
			ok, _ := bCar.Exist()
			if ok {
				ok, err := bCar.Update()
				if !ok {
					fmt.Println("add cart fail ", err)
					c.JSON(http.StatusOK, gin.H{
						"msg":   "fail",
						"type":  0,
						"total": 0,
						"price": 0,
					})
				} else {
					tType, total, price := bCar.GetInfo()
					c.JSON(http.StatusOK, gin.H{
						"msg":   "ok",
						"type":  tType,
						"total": total,
						"price": price,
					})
				}
			} else {
				ok, err := bCar.Add()
				if !ok {
					fmt.Println("add cart fail ", err)
					c.JSON(http.StatusOK, gin.H{
						"msg":   "fail",
						"type":  0,
						"total": 0,
						"price": 0,
					})
				} else {
					tType, total, price := bCar.GetInfo()
					c.JSON(http.StatusOK, gin.H{
						"msg":   "ok",
						"type":  tType,
						"total": total,
						"price": price,
					})
				}
			}
		}
	}
}
