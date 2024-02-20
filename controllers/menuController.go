package controller

import (
	"restaurant-management-backend/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
