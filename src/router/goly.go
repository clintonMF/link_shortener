package router

import (
	"Go_shortener/src/controller"
	"Go_shortener/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesGoly(r *gin.RouterGroup) {
	// r.POST("/", controller.NewGoly)
	// r.GET("/:id", controller.GetGoly)
	// r.PUT("/:id", controller.UpdateGoly)
	// r.DELETE("/:id", controller.DeleteGoly)
	r.GET("/", controller.GetPublicGolies)
	r.POST("/", middleware.OptionalAuth, controller.NewGoly)
	r.GET("/:id", middleware.OptionalAuth, controller.GetGoly)
	r.PUT("/:id", middleware.RequireAuth, controller.UpdateGoly)
	r.DELETE("/:id", middleware.RequireAuth, controller.DeleteGoly)
}
