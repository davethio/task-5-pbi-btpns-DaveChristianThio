package photocontroller

import (
	"net/http"
	"strconv"

	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/database"
	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/helpers"
	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/models"
	"github.com/gin-gonic/gin"
)

func GetUserIDByUsername(username any) int64 {
	var user models.User

	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return -1
	}

	return user.Id
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	var username = c.MustGet("userID")

	userID := GetUserIDByUsername(username)

	if userID == -1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	// Fetch photos from the database for the specific UserID
	if err := database.DB.Preload("Users").Where("user_id = ?", userID).Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	// Output JSON, which excludes user email, hashed password, and photo to protect user privacy and security
	var length int = len(photos)
	var output = make([]models.Photo, length)

	for i := 0; i < len(photos); i++ {
		output[i].ID = photos[i].ID
		output[i].Title = photos[i].Title
		output[i].Caption = photos[i].Caption
		output[i].PhotoUrl = photos[i].PhotoUrl
		output[i].UserID = photos[i].UserID
		output[i].Users.Username = photos[i].Users.Username
		output[i].Users.CreatedAt = photos[i].Users.CreatedAt
		output[i].Users.UpdatedAt = photos[i].Users.UpdatedAt
	}

	c.JSON(http.StatusOK, output)
}

func CreatePhoto(c *gin.Context) {
	var photo models.Photo

	var username = c.MustGet("userID")
	userID := GetUserIDByUsername(username)
	photo.UserID = userID

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

func UpdatePhoto(c *gin.Context) {
	
	userPhoto := c.Param("photoId")

	// Check if the photo is a valid integer
	id, err := strconv.ParseInt(userPhoto, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid photo"})
		return
	}

	var photosToCheck models.Photo
	var photo models.Photo

	// Check if the user exists
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Where("id = ?", id).Find(&photosToCheck).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	var username = c.MustGet("userID")
	userID := GetUserIDByUsername(username)
	if photosToCheck.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access Denied"})
		return
	}

	if database.DB.Where("id = ?", id).Model(&photo).Updates(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Cannot update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data successfully updated"})
}

func DeletePhoto(c *gin.Context) {
	var photo models.Photo
	userPhoto := c.Param("photoId")

	id, err := strconv.ParseInt(userPhoto, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid photo"})
		return
	}

	result := database.DB.First(&photo, id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Photo not found"})
		return
	}

	var username = c.MustGet("userID")
	userID := GetUserIDByUsername(username)
	if photo.UserID != userID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access Denied"})
		return
	}

	if rowsAffected := database.DB.Delete(&photo).RowsAffected; rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo successfully deleted"})
}
