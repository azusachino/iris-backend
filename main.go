package main

import (
	"iris/pkg/orm"
	"iris/pkg/setting"
	"iris/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	// TODO docker打包后找不到配置文件
	setting.SetUp()
	// CGO 与 SQLite, 考虑使用Mysql
	// TODO 2021/06/01 15:21:14 Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
	orm.SetUp()
}

func main() {
	app := gin.Default()

	cors := cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true允许所有
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 1 * time.Hour,
	})

	app.Static("/assets", "./assets")

	app.Use(cors)
	// app.Use(middleware.Logger())

	router.ApplyRouter(app)

	app.Run(":" + setting.AppConfig.Port)
}
