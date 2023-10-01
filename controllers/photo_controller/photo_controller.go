package photocontroller

import (
	"net/http"
	"strconv"

	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/database"
	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/helpers"

	// "github.com/davethio/task-5-pbi-btpns-DaveChristianThio/middlewares"
	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/models"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func GetPhotos(c *gin.Context) {
	// var photo []models.Photo

	// // Fetch users from the database
	// if err := database.DB.Preload("Users").Find(&photo).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
	// 	return
	// }

	// c.JSON(http.StatusOK, photo)

	var photos []models.Photo
	userID := c.MustGet("userID").(int64)

	// if !exists {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "UserID not found in context"})
	// 	return
	// }

	// Fetch photos from the database for the specific UserID
	if err := database.DB.Preload("Users").Where("user_id = ?", userID).Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}
	userIDStr := strconv.FormatInt(userID, 10)
	response := map[string]string{"message": userIDStr}
	helpers.JSONResponse(c, http.StatusOK, response)
	c.JSON(http.StatusOK, photos)
}

func CreatePhoto(c *gin.Context) {
	var photo models.Photo

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := database.DB.Create(&photo).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.JSONResponse(c, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helpers.JSONResponse(c, http.StatusOK, response)
}

// UpdatePhoto allows users to update photo details
func UpdatePhoto(c *gin.Context) {
	// Implement logic to update photo details
	// Check permissions, validate input, update photo data, and respond

	// Get the user ID from the request parameters
	id := c.Param("photoId")
	// Check if the user exists
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&photo).Where("id = ?", id).Updates(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Cannot update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data successfully updated"})
}

// DeletePhoto allows users to delete a photo
func DeletePhoto(c *gin.Context) {
	var photo models.Photo

	// Parse the photo from the request URL parameters
	userPhoto := c.Param("photoId")

	// Check if the photo is a valid integer
	id, err := strconv.ParseInt(userPhoto, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid photo"})
		return
	}

	// Check if the photo exists before attempting to delete
	result := database.DB.First(&photo, id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Photo not found"})
		return
	}

	// Delete the photo from the database
	if rowsAffected := database.DB.Delete(&photo).RowsAffected; rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo successfully deleted"})
}
