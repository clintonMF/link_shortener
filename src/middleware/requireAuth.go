package middleware

import (
	"Go_shortener/src/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	/*
	   This function authenticates the user and returns
	   unauthorized if the user is unknown or unauthenticated.
	   This ensures that unknown or unauthenticated users do not get access
	   to private resources.
	*/
	tokenString, err := c.Cookie("UserAuth")
	if err != nil {
		// User is unauthenticated, return Unauthorized
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the algorithm is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")

		my_secret_key := os.Getenv("my_secret_key")

		return []byte(my_secret_key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			// Token is valid but has expired, user is unauthenticated
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userID := claims["sub"].(float64)
		user, err := models.GetUserByID(uint(userID))
		if user == nil || err != nil {
			// User is unauthorized or unauthenticated
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)

		c.Next()
	} else {
		// User is unauthenticated or unauthorized
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func OptionalAuth(c *gin.Context) {
	/*
		Unlike the RequireAuth function which returns unauthorized
		if the user is unknown. This function is used to determine
		if the user is known or unknown so that we can modify what
		the user sees.
	*/
	tokenString, err := c.Cookie("UserAuth")
	if err != nil {
		c.Next()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		my_secret_key := os.Getenv("my_secret_key")

		return []byte(my_secret_key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		userID := claims["sub"].(float64)
		user, err := models.GetUserByID(uint(userID))
		if user == nil || err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
