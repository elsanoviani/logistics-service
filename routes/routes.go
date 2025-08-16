package routes

import (
	"logistics-service/controllers"
	"logistics-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	authGroup := r.Group("/shipments")
	authGroup.Use(middleware.AuthRequired())
	{
		authGroup.POST("/", controllers.CreateShipment)
		authGroup.PUT("/:tracking_number/status", controllers.UpdateShipmentStatus)
		authGroup.GET("/:tracking_number", controllers.TrackShipment)
		authGroup.GET("/", controllers.GetUserShipments)
	}

	return r
}
