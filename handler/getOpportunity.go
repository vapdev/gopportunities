package handler

import "github.com/gin-gonic/gin"

func GetOpportunityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getOpportunity",
	})
}
