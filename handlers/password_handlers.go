package handlers

import (
	"fmt"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/pomntv/nantavit_agnos_backend/database"
	"github.com/pomntv/nantavit_agnos_backend/models"
)

func ListPassword(c *gin.Context) {
	var password []models.Password_recommendation
	database.DB.Db.Find(&password)

	c.JSON(200, password)
}

func CreatePassword(c *gin.Context) {
	var pass models.Password_recommendation
	if err := c.ShouldBindJSON(&pass); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// not pass num_of_steps > 0
	num_of_steps := validatePassword(pass.Init_password)
	if num_of_steps != 0 {
		c.JSON(400, gin.H{"num_of_steps": num_of_steps})
		clientIP := c.ClientIP()
		// Save log to database
		logmessage := fmt.Sprintf("Client IP: %s : created error with num_of_steps %d password %s.", clientIP, num_of_steps, pass.Init_password)
		logModel := models.Log{Message: logmessage}
		if err := database.DB.Db.Create(&logModel).Error; err != nil {
			// Handle error saving log to database
			fmt.Println("Failed to save log to database:", err)
		}
		return
	}
	// pass num_of_steps = 0
	database.DB.Db.Create(&pass)
	c.JSON(200, gin.H{"num_of_steps": num_of_steps})
	clientIP := c.ClientIP()
	// Save log to database
	logmessage := fmt.Sprintf("Client IP: %s : recommendation created successfully with ID %d password %s.", clientIP, pass.ID, pass.Init_password)
	logModel := models.Log{Message: logmessage}
	if err := database.DB.Db.Create(&logModel).Error; err != nil {
		// Handle error saving log to database
		fmt.Println("Failed to save log to database:", err)
	}

}

func validatePassword(password string) int {
	var (
		upp, low, num, repeat bool
	)
	fmt.Printf("Input Password: %v\n", password)
	count := 0
	consecutiveCount := 1 // Initialize consecutive character count

	for i, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
		case unicode.IsLower(char):
			low = true
		case unicode.IsNumber(char):
			num = true
		default:
			// return count, "Invalid character found"
		}
		// Check for repeating characters
		if i > 0 && password[i] == password[i-1] {
			consecutiveCount++
		} else {
			consecutiveCount = 1
		}

		// Check for three repeating characters in a row
		if consecutiveCount >= 3 {
			repeat = true
		}
	}

	//### check Upper,Lower,Emoji ########################################################################
	if !upp {
		fmt.Printf("----> Password Contains at least 1 uppercase letter\n")

		count++
	}
	if !low {
		fmt.Printf("----> Password Contains at least 1 lowercase letter\n")
		count++
	}
	if !num {
		fmt.Printf("----> Password Contains at least 1 digit,\n")
		count++
	}
	if repeat {
		fmt.Printf("----> Password does not contain 3 repeating characters\n")
		count++
	}

	if len(password) < 6 || len(password) >= 20 {
		fmt.Printf("----> Password length should be between 6 and 19 characters\n")
		count++
	}

	fmt.Printf("--------> count: %v\n", count)

	return count
}
