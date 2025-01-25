package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
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
		LeaveType string `form:"leaveType" binding:"required"`
		StartDate string `form:"startDate" binding:"required"`
		EndDate   string `form:"endDate" binding:"required"`
		Reason    string `form:"reason" binding:"required"`
	}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File Upload Failed"})
		return
	}
	fileName := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
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
		UserId:    userID.(uint),
		LeaveType: request.LeaveType,
		StartDate: startDate,
		EndDate:   endDate,
		Reason:    request.Reason,
		FilePath:  fileName,
	}
	if result := config.DB.Create(&leaveRecord); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create leave record"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Leave application submitted",
	})
}

func RemainingLeave(c *gin.Context) {
	UserId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var leaveRecords []models.LeaveRecord
	if err := config.DB.Where("user_id = ?", UserId).Find(&leaveRecords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Leave Record "})
		return
	}
	usedLeaves := make(map[string]int)
	for _, record := range leaveRecords {
		duration := int(record.EndDate.Sub(record.StartDate).Hours()/24) + 1
		var leaveType models.LeaveType
		if err := config.DB.Where("type = ?", record.LeaveType).First(&leaveType).Error; err != nil {
			fmt.Println("Error finding LeaveType:", err)
			continue
		}
		usedLeaves[leaveType.Type] += duration
	}
	var leaveType []models.LeaveType
	config.DB.Find(&leaveType)
	balance := make(map[string]int)
	for _, leaveType := range leaveType {
		balance[leaveType.Type] = leaveType.MaxDays - usedLeaves[leaveType.Type]
	}
	c.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"leave_balance": balance,
	})
}

func ViewLeaveApplication(c *gin.Context) {
	UserId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorized"})
		return
	}
	var leaveRecord []models.LeaveRecord
	if err := config.DB.Where("user_id = ?", UserId).Find(&leaveRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errror": "Fetching data from Leave Records"})
		return
	}

	type LeaveResponse struct {
		ID        uint      `json:"id"`
		LeaveType string    `json:"leave_type"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Status    string    `json:"status"`
		Reason    string    `json:"reason"`
	}
	var leaveApplication []LeaveResponse
	for _, record := range leaveRecord {
		leaveApplication = append(leaveApplication, LeaveResponse{
			ID:        record.ID,
			LeaveType: record.LeaveType,
			StartDate: record.StartDate,
			EndDate:   record.EndDate,
			Reason:    record.Reason,
			Status:    record.Status,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":            "success",
		"leave_application": leaveApplication,
	})
}
