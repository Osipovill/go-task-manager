package exceptions

import "github.com/gin-gonic/gin"

func HandleError(c *gin.Context, err error, message string) {
	c.JSON(500, gin.H{
		"error":   err.Error(),
		"message": message,
	})
}
