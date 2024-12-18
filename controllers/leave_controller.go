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
		LeaveTypeId uint   `form:"leaveid"`
		StartDate   string `form:"startDate"`
		EndDate     string `form:"endDate"`
		Reason      string `form:"reason"`
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
		UserId:      userID.(uint),
		LeaveTypeId: request.LeaveTypeId,
		StartDate:   startDate,
		EndDate:     endDate,
		Reason:      request.Reason,
		FilePath:    fileName,
	}
	if result := config.DB.Create(&leaveRecord); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create leave record"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Leave application submitted"})
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
		duration := int(record.EndDate.Sub(record.StartDate).Hours() / 24)
		var leaveType models.LeaveType
		config.DB.First(&leaveType, record.LeaveTypeId)
		usedLeaves[leaveType.Type] += duration
	}
	var leaveType []models.LeaveType
	config.DB.Find(&leaveType)
	balance := make(map[string]int)
	for _, leaveType := range leaveType {
		balance[leaveType.Type] = leaveType.MaxDays - usedLeaves[leaveType.Type]
	}
	c.JSON(http.StatusOK, gin.H{"leave_balance": balance})
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
		ID          uint      `json:"id"`
		LeaveTypeID uint      `json:"leave_type_id"`
		StartDate   time.Time `json:"start_time"`
		EndDate     time.Time `json:"end_date"`
		Status      string    `json:"status"`
		Reason      string    `json:"reason"`
	}
	var leaveApplication []LeaveResponse
	for _, record := range leaveRecord {
		leaveApplication = append(leaveApplication, LeaveResponse{
			ID:          record.ID,
			LeaveTypeID: record.LeaveTypeId,
			StartDate:   record.StartDate,
			EndDate:     record.EndDate,
			Reason:      record.Reason,
			Status:      record.Status,
		})
	}
	c.JSON(http.StatusOK, gin.H{"leave_application": leaveApplication})
}
