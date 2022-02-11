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

		// 文章模块的路由接口

	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
