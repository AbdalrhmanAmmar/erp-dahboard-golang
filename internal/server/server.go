package server

import (
	"github.com/gin-gonic/gin"
)

// New يجهّز الـ router ويركّب الـ middlewares والروابط.
func New() *gin.Engine {
	r := gin.Default() // بيجي معاه Logger و Recovery تلقائياً

	// نقطة فحص بسيطة نتأكد بيها إن السيرفر عايش
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// نسلّم الـ router لملف الـ routes عشان يركّب الموديولات
	registerRoutes(r)

	return r
}