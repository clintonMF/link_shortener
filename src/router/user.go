package router

import (
	"Go_shortener/src/controller"
	"Go_shortener/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesUser(r *gin.RouterGroup) {
	// r.GET("/r/:redirect", controller.Redirect)
	r.POST("/signin", controller.SignIn)
	r.POST("/signup", controller.SignUp)
	r.GET("/:id/history", middleware.RequireAuth, controller.GetUserGolies)
	// r.PUT("/:id", controller.UpdateGoly)
	// r.DELETE("/:id", controller.DeleteGoly)
	// r.GET("/:id/students", controller.ListStudentsTakingCourse)
}
