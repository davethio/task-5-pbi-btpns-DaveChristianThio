package helpers


import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, gin.H{
		"response": payload,
	})
}