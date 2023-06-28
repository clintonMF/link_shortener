package controller

import (
	"Go_shortener/src/models"
	"Go_shortener/src/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func SignUp(c *gin.Context) {
	// swagger:operation POST /user/signup createUser
	//
	// create a new user
	//
	// ---
	// tags:
	//    [user]
	// produces:
	// - application/json
	// parameters:
	// - name: user
	//   in: body
	//   description: course details
	//   required: true
	//   type: object
	//   schema:
	//     "$ref": "#/definitions/User"
	// responses:
	//   '201':
	//     description: course created
	//     schema:
	//       type: object
	//       "$ref": "#/definitions/UserResponse"
	//   '400':
	//     description: bad request
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})

		return
	}

	// confirm the strength and validity of email
	if !utils.PasswordStrength(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Password is too weak",
		})

		return
	}

	if !utils.IsValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Not a valid email",
		})

	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	user.Password = hash

	createdUser, err := user.CreateUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"error":   err.Error(),
			"message": "Enter appropriate details",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": utils.CreatedMessage("goly"),
		"User":    createdUser,
	})
}

func SignIn(c *gin.Context) {
	// swagger:operation POST /user/signin LoginUser
	//
	// user login
	//
	// ---
	// tags:
	//    [user]
	// produces:
	// - application/json
	// parameters:
	// - name: user
	//   in: body
	//   description: login details
	//   required: true
	//   type: object
	//   schema:
	//     "$ref": "#/definitions/UserLogin"
	// responses:
	//   '204':
	//     description: user logged in, no content
	//   '400':
	//     description: bad request
	var body struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})

		return
	}

	user, err := models.GetUserByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	if utils.ComparePasswords(body.Password, user.Password) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid password or email",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("my_secret_key")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "internal error",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})

}
