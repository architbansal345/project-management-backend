package controllers

import (
	"net/http"
	"project-management-backend/config"
	"project-management-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func ApplyLeave(c *gin.Context) {
	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request struct {
		LeaveTypeId uint   `json:"leaveid"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
		Reason      string `json:"reason"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start Date format"})
		return
	}
	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start Date format"})
		return
	}

	leaveRecord := models.LeaveRecord{
		UserId:      userID.(uint),
		LeaveTypeId: request.LeaveTypeId,
		StartDate:   startDate,
		EndDate:     endDate,
		Reason:      request.Reason,
	}
	if result := config.DB.Create(&leaveRecord); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create leave record"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Leave application submitted"})
}
