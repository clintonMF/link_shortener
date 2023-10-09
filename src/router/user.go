package router

import (
	"Go_shortener/src/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesUser(r *gin.RouterGroup) {
	r.POST("/signin", controller.SignIn)
	r.POST("/signup", controller.SignUp)
	r.GET("/signout", controller.SignOut)
}
