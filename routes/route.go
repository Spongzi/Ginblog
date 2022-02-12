package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	routerv1 := r.Group("api/v1")
	{
		// 用户模块的路由接口
		routerv1.POST("user/add", v1.AddUser)
		routerv1.GET("users", v1.GetUsers)
		routerv1.PUT("user/:id", v1.EditUser)
		routerv1.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		routerv1.POST("category/add", v1.AddCategory)
		routerv1.GET("category", v1.GetCategory)
		routerv1.PUT("category/:id", v1.EditCategory)
		routerv1.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		routerv1.POST("article/add", v1.AddArticle)
		routerv1.GET("article", v1.GetArticle)
		routerv1.PUT("article/:id", v1.EditArticle)
		routerv1.DELETE("article/:id", v1.DeleteArticle)
	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
