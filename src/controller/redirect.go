package controller

import (
	"Go_shortener/src/models"
	"Go_shortener/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func updateClicks(url string) {
	// this function updates the click in the background
	goly, err := models.GetGolyByURL(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	goly.Clicked += 1

	err = models.UpdateGoly(goly)
	if err != nil {
		fmt.Println(err.Error())
	}
	cache.Set(url, goly.Redirect, 0)
	fmt.Println("clicks updated")
}

func Redirect(c *gin.Context) {
	golyUrl := c.Param("redirect")
	value, found := cache.Get(name + golyUrl)

	if found {
		/*
			if the cache has the redirect link
			redirect the user and perform update the click value
			in the background
		*/
		go updateClicks(name + golyUrl)
		c.Redirect(http.StatusTemporaryRedirect, value.(string))
		return
	}

	goly, err := models.GetGolyByURL(name + golyUrl)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	goly.Clicked += 1
	cache.Set(name+golyUrl, goly.Redirect, 0)

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
