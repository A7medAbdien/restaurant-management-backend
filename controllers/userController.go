package controller

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// context with timeout
		// retrieve with pagination
		// decode
		// response
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// context with timeout
		// retrieve by Id and decode
		// response
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// context with timeout
		// bind and decode
		// validate
		// add extra fields (created_at, updated_at, ID)
		// check if already exist
		// hash password
		// generate token
		// response
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// context with timeout
		// bind and decode
		// validate
		// find user
		// verify password
		// refresh tokens
		// response
	}
}

func HashPassword(password string) string {}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {}
