package handler

import "github.com/gin-gonic/gin"

func CreateOpportunityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "createOpportunity",
	})
}
