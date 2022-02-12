package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	// 需要中间件的
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
	}
	routerv1 := r.Group("api/v1")
	{
		routerv1.POST("user/add", v1.AddUser)
		routerv1.GET("users", v1.GetUsers)
		routerv1.GET("category", v1.GetCategory)
		routerv1.GET("article", v1.GetArticle)
		routerv1.GET("article/list/:id", v1.GetCatArt)
		routerv1.GET("article/info/:id", v1.GetArticleInfo)
		routerv1.POST("login", v1.Login)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
