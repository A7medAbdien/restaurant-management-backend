package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant-management-backend/database"
	"restaurant-management-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		// context and timeout
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// retrieve and decode
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing the menu items"})
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}

		// response
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// context and timeout
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// retrieve and decode
		menuId := c.Param("menu_id")
		var menu models.Menu

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "error occurred while fetching the menu"})
		}

		// response
		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// context and timeout
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// bind and validate
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		// insert
		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil {
			msg := fmt.Sprintf("Menu item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		// response
		c.JSON(http.StatusOK, result)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Bind JSON data from the request into a Menu struct
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Extract the menu ID from the request parameters
		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		// Check if the provided start and end dates are within a valid time span
		if menu.Start_Date != nil && menu.End_Date != nil && !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time span"})
			return
		}

		// Prepare the update object with fields to be updated
		var updateObj primitive.D
		updateObj = append(updateObj, bson.E{"start_date", menu.Start_Date})
		updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{"name", menu.Name})
		}
		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{"category", menu.Category})
		}
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

		// Set up options for the UpdateOne operation, including upsert
		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		// Perform the update operation on the MongoDB collection
		result, err := menuCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		// Handle errors during the update operation
		if err != nil {
			msg := "Menu update failed"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		// Respond with the result of the update operation
		c.JSON(http.StatusOK, result)

	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}
