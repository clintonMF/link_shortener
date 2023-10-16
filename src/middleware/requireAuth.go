package middleware

import (
	"Go_shortener/src/services"
	"Go_shortener/src/utils"
	"net/http"

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

	token, err := utils.ValidateToken(tokenString)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		utils.CheckIfExpired(c, claims)

		userID := claims["sub"].(float64)
		user, err := services.GetUserByID(uint(userID))
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

	token, err := utils.ValidateToken(tokenString)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		utils.CheckIfExpired(c, claims)

		userID := claims["sub"].(float64)
		user, err := services.GetUserByID(uint(userID))
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
