package handler

import "github.com/gin-gonic/gin"

func GetOpportunitiesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getOpportunities",
	})
}
