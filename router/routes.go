package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vapdev/gopportunities.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunities", handler.GetOpportunitiesHandler)
		v1.GET("/opportunities/:id", handler.GetOpportunityHandler)
		v1.POST("/opportunities", handler.CreateOpportunityHandler)
		v1.PUT("/opportunities/:id", handler.UpdateOpportunityHandler)
		v1.DELETE("/opportunities/:id", handler.DeleteOpportunityHandler)
	}
}
