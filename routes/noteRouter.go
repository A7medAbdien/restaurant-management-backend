package routes

import (
	controller "restaurant-management-backend/controllers"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes", controller.GetNotes())
	incomingRoutes.GET("/notes/:note_id", controller.GetNote())
	incomingRoutes.POST("/notes", controller.CreateNote())
	incomingRoutes.PATCH("/notes/:note_id", controller.UpdateNote())
}
