package router

import "github.com/gin-gonic/gin"

func getOpportunities(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getOpportunities",
	})
}

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunities", getOpportunities)
		// v1.GET("/opportunities/:id", getOpportunity)
		// v1.POST("/opportunities", createOpportunity)
		// v1.PUT("/opportunities/:id", updateOpportunity)
		// v1.DELETE("/opportunities/:id", deleteOpportunity)
	}
}
