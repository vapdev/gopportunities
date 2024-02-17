package handler

import "github.com/gin-gonic/gin"

func GetOpportunitiesHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getOpportunities",
	})
}

func GetOpportunityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getOpportunity",
	})
}

func CreateOpportunityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "createOpportunity",
	})
}

func UpdateOpportunityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "updateOpportunity",
	})
}

func DeleteOpportunityHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "deleteOpportunity",
	})
}
