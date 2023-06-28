package controller

import (
	"Go_shortener/src/models"
	"Go_shortener/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	// swagger:operation GET /r/{redirect} getRedirect
	//
	// performing the redirect operation
	//
	// ---
	// tags:
	//    [redirect]
	// produces:
	// - application/json
	// parameters:
	// - name: redirect
	//   in: query
	//   description: shortened link string
	//   required: true
	//   type: integer
	//
	// responses:
	//   '301':
	//     description: redirect to the original website/link
	//   '400':
	//     description: bad request
	//   '401':
	//     description: Unauthorized
	//   '404':
	//     description: goly not found

	golyUrl := c.Param("redirect")
	goly, err := models.GetGolyByURL(name + golyUrl)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	goly.Clicked += 1

	err = models.UpdateGoly(goly)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Redirect(http.StatusTemporaryRedirect, goly.Redirect)
}

func GenerateQRCode(c *gin.Context) {
	// swagger:operation GET /r/{redirect}/generateQRcode getQrcode
	//
	// Generating the QR code
	//
	// ---
	// tags:
	//    [redirect]
	// produces:
	// - application/json
	// parameters:
	// - name: redirect
	//   in: query
	//   description: shortened link string
	//   required: true
	//   type: integer
	//
	// responses:
	//   '200':
	//     description: QR code generated
	//     type: image

	golyUrl := c.Param("redirect")
	goly, err := models.GetGolyByURL(name + golyUrl)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	qrCode, err := utils.GenerateQRCode(goly.Goly)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Data(http.StatusOK, "image/png", qrCode)
}
