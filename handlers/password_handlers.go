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

	if num_of_steps, result := validatePassword(pass.Init_password); num_of_steps != 0 {
		c.JSON(400, gin.H{"num_of_steps": num_of_steps, "result": result})
		return
	}

	database.DB.Db.Create(&pass)
	c.JSON(200, pass)

}

func validatePassword(password string) (int, string) {
	var (
		upp, low, num, sym bool
		tot                uint8
	)
	fmt.Printf("Lenght Input Password: %v\n", len(password))
	count := 0
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return count, "Invalid character found"
		}
	}

	fmt.Printf("tot: %v\n", tot)
	//### check Upper,Lower,Emoji ########################################################################
	if !upp {
		return count, "Must include at least one uppercase letter"
	}
	if !low {
		return count, "Must include at least one lowercase letter"
	}
	if !num {
		return count, "Must include at least one number"
	}
	if !sym {
		return count, "Must include at least one special character"
	}
	if tot < 6 || tot >= 20 {
		return count, "Password length should be between 6 and 19 characters"
	}

	return count, ""
}
