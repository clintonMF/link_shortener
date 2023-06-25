package router

import (
	"Go_shortener/src/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesRedirect(r *gin.RouterGroup) {
	r.GET("/:redirect", controller.Redirect)
	r.GET("/:redirect/generateQRCode", controller.GenerateQRCode)
}
