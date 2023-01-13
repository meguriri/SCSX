package main

import (
	"fmt"
	"github.com/meguriri/SCSX/data"
	"github.com/meguriri/SCSX/redis"
	"github.com/meguriri/SCSX/router"
)

func main() {
	//初始化redis
	if err := redis.InitClient(); err != nil {
		fmt.Println(err)
		return
	}
	//初始化mysql
	data.InitDB()
	//初始化路由器
	r := router.InitRouter()
	//运行
	r.Run(":8080")
}
