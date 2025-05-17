package controllers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func SendEmail(to, subject, body string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("APP_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", from, password, smtpHost)

	message := []byte("Subject: " + subject + "\r\n" + "\r\n" + body)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}

func OnBoarding(c *gin.Context) {
	var json struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	subject := "Welcome to the Onboarding Process"
	body := fmt.Sprintf("Hi %s,\n\nPlease complete your onboarding process here: [link to your self-onboarding form].", json.Name)

	err := SendEmail(json.Email, subject, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Invitation sent successfully"})
}
