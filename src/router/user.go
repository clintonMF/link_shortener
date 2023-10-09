package router

import (
	"Go_shortener/src/controller"
	"Go_shortener/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesUser(r *gin.RouterGroup) {
	r.POST("/signin", controller.SignIn)
	r.POST("/signup", controller.SignUp)
	r.GET("/signout", controller.SignOut)
	r.GET("/:id/history", middleware.RequireAuth, controller.GetUserGolies)
}
