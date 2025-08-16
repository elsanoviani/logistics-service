package controllers

import (
	"logistics-service/database"
	"logistics-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShipment(c *gin.Context) {
	var shipment models.Shipment
	if err := c.ShouldBindJSON(&shipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	shipment.UserID = userID.(string)
	shipment.Status = "Shipped"

	if err := database.DB.Create(&shipment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, shipment)
}

func UpdateShipmentStatus(c *gin.Context) {
	trackingNumber := c.Param("tracking_number")
	var body struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var shipment models.Shipment
	if err := database.DB.Where("tracking_number = ?", trackingNumber).First(&shipment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment not found"})
		return
	}

	shipment.Status = body.Status
	database.DB.Save(&shipment)

	c.JSON(http.StatusOK, shipment)
}

func TrackShipment(c *gin.Context) {
	trackingNumber := c.Param("tracking_number")
	var shipment models.Shipment
	if err := database.DB.Where("tracking_number = ?", trackingNumber).First(&shipment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment not found"})
		return
	}
	c.JSON(http.StatusOK, shipment)
}

func GetUserShipments(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var shipments []models.Shipment
	database.DB.Where("user_id = ?", userID.(string)).Find(&shipments)
	c.JSON(http.StatusOK, shipments)
}
