package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropmartiniano/mvc-golang/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	userRouter := r.Group("/user")
	userRouter.GET("/id/:id", controller.FindUserById)
	userRouter.GET("/email/:email" , controller.FindUserByEmail)
	userRouter.POST("/" , controller.CreateUser)
	userRouter.PUT("/:id", controller.UpdateUser)
	userRouter.DELETE("/", controller.DeleteUser)
}
