package handlers

import (
	"fmt"
	"net/http"
	"smarthome/models"
	"smarthome/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SensorHandler handles sensor-related requests
type SensorHandler struct {
	SensorAPI          *services.DevicesAPIClient
	TemperatureService *services.TemperatureService
}

// NewSensorHandler creates a new SensorHandler
func NewSensorHandler(devicesService *services.DevicesAPIClient, temperatureService *services.TemperatureService) *SensorHandler {
	return &SensorHandler{
		SensorAPI:          devicesService,
		TemperatureService: temperatureService,
	}
}

// RegisterRoutes registers the sensor routes
func (h *SensorHandler) RegisterRoutes(router *gin.RouterGroup) {
	sensors := router.Group("/sensors")
	{
		sensors.GET("", h.GetSensors)
		sensors.GET("/:id", h.GetSensorByID)
		sensors.POST("", h.CreateSensor)
		sensors.PUT("/:id", h.UpdateSensor)
		sensors.DELETE("/:id", h.DeleteSensor)
		sensors.PATCH("/:id/value", h.UpdateSensorValue)
		sensors.GET("/temperature/:location", h.GetTemperatureByLocation)
	}
}

// GetSensors handles GET /api/v1/sensors
func (h *SensorHandler) GetSensors(c *gin.Context) {
	ctx := c.Request.Context()

	sensors, err := h.SensorAPI.GetDevices(ctx)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapSensors(sensors))
}

// GetSensorByID handles GET /api/v1/sensors/:id
func (h *SensorHandler) GetSensorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	ctx := c.Request.Context()

	device, err := h.SensorAPI.GetSensorByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sensor not found"})
		return
	}

	c.JSON(http.StatusOK, mapSensorFromDTO(device))
}

// GetTemperatureByLocation handles GET /api/v1/sensors/temperature/:location
func (h *SensorHandler) GetTemperatureByLocation(c *gin.Context) {
	location := c.Param("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location is required"})
		return
	}

	// Fetch temperature data from the external API
	tempData, err := h.TemperatureService.GetTemperature(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch temperature data: %v", err),
		})
		return
	}

	// Return the temperature data
	c.JSON(http.StatusOK, gin.H{
		"location":    tempData.Location,
		"value":       tempData.Value,
		"unit":        tempData.Unit,
		"status":      tempData.Status,
		"timestamp":   tempData.Timestamp,
		"description": tempData.Description,
	})
}

func (h *SensorHandler) CreateSensor(c *gin.Context) {
	var req models.SensorCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sensor, err := h.SensorAPI.CreateSensor(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sensor)
}

// UpdateSensor handles PUT /api/v1/sensors/:id
func (h *SensorHandler) UpdateSensor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	var req models.SensorUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	sensor, err := h.SensorAPI.UpdateSensor(ctx, id, req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("Failed to update sensor: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, sensor)
}

// DeleteSensor handles DELETE /api/v1/sensors/:id
func (h *SensorHandler) DeleteSensor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	ctx := c.Request.Context()

	if err := h.SensorAPI.DeleteSensor(ctx, id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("Failed to delete sensor: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sensor deleted successfully",
	})
}

// UpdateSensorValue handles PATCH /api/v1/sensors/:id/value
func (h *SensorHandler) UpdateSensorValue(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	var req struct {
		Value  float64 `json:"value" binding:"required"`
		Status string  `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	if err := h.SensorAPI.UpdateSensorValue(ctx, id, req.Value, req.Status); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("Failed to update sensor value: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sensor value updated successfully",
	})
}

func mapSensors(dtos []models.Device) []models.Sensor {
	result := make([]models.Sensor, 0, len(dtos))

	for _, dto := range dtos {
		result = append(result, mapSensorFromDTO(&dto))
	}

	return result
}

func mapSensorFromDTO(dto *models.Device) models.Sensor {
	return models.Sensor{
		ID:          dto.ID,
		Name:        dto.Name,
		Type:        dto.Type,
		Location:    dto.Location,
		Value:       dto.Value,
		Unit:        dto.Unit,
		Status:      dto.Status,
		LastUpdated: dto.LastUpdated,
		CreatedAt:   dto.CreatedAt,
	}
}
