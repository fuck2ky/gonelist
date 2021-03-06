package routers

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gonelist/api"
	"gonelist/conf"
	"gonelist/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.Use(static.Serve("/", static.LocalFile(conf.GetDistPATH(), false)))

	// 测试接口
	r.GET("/testapi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.GET("/login", api.Login)
	r.GET("/loginmg", api.LoginMG)
	r.GET("/auth", api.GetCode)
	//r.GET("/cancelLogin", api.CancelLogin)
	// 直接下载接口
	r.GET("/d/*path", middleware.CheckLogin(), api.Download)
	onedrive := r.Group("/onedrive")
	// 中间件判断是否已经登录
	onedrive.Use(middleware.CheckLogin())
	{
		// 主动获取所有文件，返回整个树的目录
		onedrive.GET("/getallfiles", api.MGGetFileTree)
		// 根据路径获取对应数据
		onedrive.GET("/getpath", api.CacheGetPath)
		// 直接下载文件
	}

	return r
}
