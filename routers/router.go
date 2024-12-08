package routers

import (
	"github.com/PratikforCoding/linksnap/handlers"
	"github.com/PratikforCoding/linksnap/middlewares"
	"github.com/gin-gonic/gin"
)

// SetUpRouter sets up the routes for the application
func SetUpRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/urls", middlewares.RateLimitMiddleware(),handlers.CreateShortURL)
		apiGroup.GET("/urls", handlers.ListURLs)
		apiGroup.GET("/urls/:code", handlers.GetURLDetails)
		apiGroup.DELETE("/urls/:code", handlers.DeleteURL)
		apiGroup.GET("/urls/:code/stats", handlers.GetURLStats)
		apiGroup.GET("/urls/:code/qr", handlers.GetQRCode)
	}
	router.GET("/:shortCode", handlers.RedirectURL)
	return router
}