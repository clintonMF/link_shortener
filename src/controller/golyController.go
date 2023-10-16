package controller

import (
	"Go_shortener/src/models"
	"Go_shortener/src/services"
	"Go_shortener/src/setup"
	"Go_shortener/src/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	db    *gorm.DB = setup.GetDB()
	cache          = setup.InitCache()
)

var name string = os.Getenv("URL_REDIRECT_NAME")

func GetUserGolies(c *gin.Context) {
	/*
		This function returns all the golies a user has created
			i.e user goly history.
	*/
	user, _ := c.Get("user")

	cur_user := user.(*models.User)
	golies, err := services.GetGoliesByUserID(cur_user.ID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              "success",
		"Golies":              golies,
		"number of redirects": len(golies),
	})
}

func GetGoly(c *gin.Context) {
	var goly *models.Goly
	var err error
	id := c.Param("id")

	ID, _ := strconv.ParseUint(id, 10, 64)

	goly, err = services.GetGolyByID(uint(ID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})

		return
	}

	user, _ := c.Get("user")

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "you don't have access to this resource",
		})

		return
	}

	cur_user := user.(*models.User)

	if cur_user.ID != goly.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "you don't have access to this resource",
		})

		return
	}

	cache.Set(goly.Goly, goly.Redirect, 0)
	c.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"Goly":         goly,
		"QR Code link": goly.Goly + "/generateQRCode",
	})
}

func NewGoly(c *gin.Context) {
	/*
		This is used to create a new goly
	*/
	var goly models.Goly

	if err := c.ShouldBindJSON(&goly); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})

		return
	}

	if !utils.ValidateURL(goly.Redirect) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid URL",
		})

		return
	}

	if goly.Custom {
		if len(goly.Goly) < 10 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "custom link should be longer than 10 characters",
			})

			return
		}
	} else {
		goly.Goly = utils.GenerateShortURL(8)
	}

	// authentication
	user, _ := c.Get("user")

	if user == nil {
		goly.UserID = 0
	} else {
		cur_user := user.(*models.User)
		goly.UserID = cur_user.ID
	}

	goly.Goly = name + goly.Goly
	createdGoly, err := goly.CreateGoly()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"error":   err.Error(),
			"message": "Enter appropriate details",
		})

		return
	}
	cache.Set(goly.Goly, goly.Redirect, 0)

	c.JSON(http.StatusCreated, gin.H{
		"status":       "success",
		"message":      utils.CreatedMessage("goly"),
		"Goly":         createdGoly,
		"QR Code link": createdGoly.Goly + "/generateQRCode",
	})
}

func UpdateGoly(c *gin.Context) {
	/*
		This is used to modify a  goly
	*/
	id := c.Param("id")

	ID, _ := strconv.ParseUint(id, 10, 64)

	goly, err := services.GetGolyByID(uint(ID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  utils.ErrNotFound("course", int(ID)).Error(),
		})

		return
	}

	_, found := cache.Get(goly.Goly)

	if found {
		cache.Delete(goly.Goly)
	}

	var updatedGoly models.Goly

	if err := c.ShouldBindJSON(&updatedGoly); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  utils.ErrInvalidInput("Title and Code").Error(),
		})

		return
	}

	goly.Redirect = updatedGoly.Redirect
	if updatedGoly.Goly != "" {
		goly.Goly = updatedGoly.Goly
	}

	err = services.UpdateGoly(goly)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// update the cache data
	cache.Set(goly.Goly, goly.Redirect, 0)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": utils.UpdatedMessage("goly", int(ID)),
		"goly":    goly,
	})
}

func DeleteGoly(c *gin.Context) {
	id := c.Param("id")

	ID, _ := strconv.ParseUint(id, 10, 64)

	goly, err := services.GetGolyByID(uint(ID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})

		return
	}
	cache.Delete(goly.Goly)
	_ = services.DeleteGoly(uint(ID))

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": utils.DeletedMessage("Goly", int(ID)),
	})
}
