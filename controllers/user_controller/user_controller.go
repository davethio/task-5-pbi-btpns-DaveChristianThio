package usercontroller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/database"
	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/helpers"
	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func IndexUser(c *gin.Context) {
	var users []models.User

	// Fetch users from the database
	if err := database.DB.Preload("Photos").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func ShowUser(c *gin.Context) {
	var users []models.User
	id := c.Param("id")

	if err := database.DB.First(&users, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": users})
}



// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	// Implement user registration logic here
	// Validate input, create user, save to database, and respond

	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JSONResponse(c, http.StatusBadRequest, response)
		return
	}

	// Hash the password with bcrypt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JSONResponse(c, http.StatusInternalServerError, response)
		return
	}
	userInput.Password = string(hashPassword)

	// Insert into the database
	if err := database.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JSONResponse(c, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helpers.JSONResponse(c, http.StatusOK, response)
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	// Implement user login logic here
	// Validate input, verify credentials, generate JWT token, and respond

	var userInput models.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JSONResponse(c, http.StatusBadRequest, response)
		return
	}

	// Take user data based on username
	var user models.User
	if err := database.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username/password is not correct"}
			helpers.JSONResponse(c, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helpers.JSONResponse(c, http.StatusInternalServerError, response)
			return
		}
	}

	// password validation
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username/password is not correct"}
		helpers.JSONResponse(c, http.StatusUnauthorized, response)
		return
	}

	// generate JWT token (after password is validated)
	expTime := time.Now().Add(time.Minute * 30)
	claims := &helpers.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "task-5-pbi-btpns-DaveChristianThio",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// Sign-in algorithm declaration
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgo.SignedString(helpers.JWT_KEY)

	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JSONResponse(c, http.StatusInternalServerError, response)
		return
	}

	// set token to cookie
	c.SetCookie("token", token, 0, "/", "", false, true)

	response := map[string]string{"message": "Login successful"}
	helpers.JSONResponse(c, http.StatusOK, response)

}

func Logout(c *gin.Context){
	c.SetCookie("token", "", -1, "/", "", false, true)

	response := map[string]string{"message": "Logout successful"}
	helpers.JSONResponse(c, http.StatusOK, response)
}

// UpdateUser allows users to update their profile
func UpdateUser(c *gin.Context) {
	// Get the user ID from the request parameters
	id := c.Param("id")

	// Check if the user exists
	var users models.User
	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&users).Where("id = ?", id).Updates(&users).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Cannot update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data successfully updated"})
}

// DeleteUser allows users to delete their account
func DeleteUser(c *gin.Context) {
	var user models.User

	// Parse the userId from the request URL parameters
	userId := c.Param("id")

	// Check if the userId is a valid integer
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid userId"})
		return
	}

	// Check if the user exists before attempting to delete
	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Delete the user from the database
	if rowsAffected := database.DB.Delete(&user).RowsAffected; rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
}

