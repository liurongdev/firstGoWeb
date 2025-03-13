package user

import (
	"awesomeProject/handle"
	"github.com/gin-gonic/gin"
)

func Registry(r *gin.Engine) {
	userGroup := r.Group("/user").Use(handle.AuthCheck())
	{
		userGroup.POST("/query", handle.QuerySystemInfo)
		userGroup.POST("/insert", handle.InsertSystemUserInfo)
		userGroup.POST("/update", handle.UpdateSystemUserInfo)
		userGroup.POST("/delete", handle.DeleteSystemUserById)
	}

}
