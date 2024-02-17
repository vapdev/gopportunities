package handler

import "github.com/gin-gonic/gin"

func UpdateOpportunityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "updateOpportunity",
	})
}
