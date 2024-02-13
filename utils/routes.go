package utils

import (
	"cc_terminal/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartRouter(db *gorm.DB) {
	router := gin.Default()

	// Register the route handlers
	router.POST("/register", func(c *gin.Context) {
		controllers.Register(c, db)
	})

	// Start the router
	router.Run(":8080")
}
