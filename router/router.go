package router

import (
	"github.com/gin-gonic/gin"
	h "github.com/meguriri/SCSX/handler"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static/")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", h.GetIndexHtml())
	user := r.Group("/user")
	{
		user.GET("/register", h.GetRegisterHtml())
		user.POST("/register", h.UserRegister())
		user.GET("/login", h.GetLoginHtml())
		user.POST("/login", h.UserLogin())
	}
	home := r.Group("/home")
	{
		home.GET("/userShow", h.UserSHow())
		home.GET("/exit", h.Exit())
		home.GET("/allCategory", h.GetAllCategory())
	}
	product := r.Group("/product")
	{
		product.GET("/", h.GetProductsHtml())
		product.GET("/selByCid", h.GetProductByCid())
		product.GET("/info", h.GetProductHtml())
		product.GET("/productDetail", h.GetProductDetail())
		product.POST("/addCart", h.AddCart())
	}
	buyCar := r.Group("/buyCar")
	{
		buyCar.GET("/", h.GetBuyCarHtml())
		buyCar.GET("/allProducts", h.GetAllProducts())
		buyCar.POST("/delBuyCar", h.DelBuyCar())
		buyCar.POST("/delAllBuyCar", h.DelAllBuyCar())
	}
	order := r.Group("/order")
	{
		order.GET("/", h.GetConfirmOrderHtml())
		order.GET("/allProducts", h.GetAllProducts())
		order.POST("/confirmOrder", h.ConfirmOrder())
		order.GET("/okOrder", h.GetOkOrderHtml())
		order.GET("/myOrder", h.GetMyOrderHtml())
		order.GET("/myOrderList", h.GetMyOrderListHtml())
	}
	message := r.Group("/message")
	{
		message.GET("/", h.GetMessageHtml())
		message.POST("/send", h.GetMessage())
	}
	search := r.Group("/search")
	{
		search.GET("/", h.GetSearchHtml())
		search.GET("/selByName", h.GetProductByName())
	}
	return r
}
