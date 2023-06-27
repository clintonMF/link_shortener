package controller

import (
	"Go_shortener/src/models"
	"Go_shortener/src/setup"
	"Go_shortener/src/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	db    *gorm.DB = setup.GetDB()
	cache          = setup.InitCache()
)

const name string = "http://localhost:8001/r/"

func GetPublicGolies(c *gin.Context) {
	/*
		This function returns all the publicly available
		golies to anyone who goes to the site.
	*/
	pubgolies, err := models.GetPublicGolies()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":              "success",
		"Golies":              pubgolies,
		"number of redirects": len(pubgolies),
	})

	return
}

func GetUserGolies(c *gin.Context) {
	/*
		This function returns all the golies a user has created
		i.e user goly history.
	*/
	user, _ := c.Get("user")

	cur_user := user.(*models.User)
	golies, err := models.GetGoliesByUserID(cur_user.ID)
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
	/*
		This is used to get an individual Goly by its unique ID
	*/
	var goly *models.Goly
	var err error
	id := c.Param("id")

	ID, _ := strconv.ParseUint(id, 10, 64)

	/* name variable ensures that golies are saved with
	the correct ID  in the cache for easy retrieve*/
	name := "Goly with id" + string(ID)

	value, found := cache.Get(name)
	if found {
		goly = value.(*models.Goly)
	} else {
		goly, err = models.GetGolyByID(uint(ID))

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "error",
				"error":  err.Error(),
			})

			return
		}
	}

	cache.Set(name, goly, 1*time.Hour)
	// value, found = cache.Get("Goly_with_id")
	// fmt.Println(found, "here")
	user, _ := c.Get("user")

	if user == nil && !goly.Public {
		// unknwon user can not access private golies  i.e public = false
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "you don't have access to this resource",
		})

		return
	} else if user == nil && goly.Public {
		// unknwon user can access publicly available golies
		// The goly information will not be complete
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"Goly": map[string]string{
				"goly":         goly.Goly,
				"redirect":     goly.Redirect,
				"QR Code link": goly.Goly + "/generateQRCode",
			},
		})

		return
	}

	cur_user := user.(*models.User)

	if cur_user.ID != goly.UserID && !goly.Public {
		// Known user with different ID cannot access private golies i.e public = false
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "you don't have access to this resource",
		})

		return
	} else if cur_user.ID != goly.UserID && goly.Public {
		// known user with different userID can access publicly available golies
		// The goly information will not be complete
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"Goly": map[string]string{
				"goly":         goly.Goly,
				"redirect":     goly.Redirect,
				"QR Code link": goly.Goly + "/generateQRCode",
			},
		})

		return
	}

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

	/* name variable ensures that golies are saved with
	the correct ID  in the cache for easy retrieve*/
	name := "Goly with id" + string(ID)

	_, found := cache.Get(name)
	fmt.Println(found)

	goly, err := models.GetGolyByID(uint(ID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  utils.ErrNotFound("course", int(ID)).Error(),
		})

		return
	}

	var updatedGoly models.Goly

	if err := c.ShouldBindJSON(&updatedGoly); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  utils.ErrInvalidInput("Title and Code").Error(),
		})

		return
	}

	// updating the goly
	goly.Redirect = updatedGoly.Redirect
	goly.Public = updatedGoly.Public

	err = models.UpdateGoly(goly)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// update the cache data
	if found {
		cache.Set(name, goly, 1*time.Hour)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": utils.UpdatedMessage("goly", int(ID)),
		"goly":    goly,
	})
}

func DeleteGoly(c *gin.Context) {
	/*
		This is used to delete  a goly
	*/
	id := c.Param("id")

	ID, _ := strconv.ParseUint(id, 10, 64)

	err := models.DeleteGoly(uint(ID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": utils.DeletedMessage("Goly", int(ID)),
	})
}
