package server

import "github.com/gin-gonic/gin"

// registerRoutes بيجمّع روابط كل الموديولات تحت /api
func registerRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// هنا هتضيف الموديولات واحد واحد لاحقاً، مثلاً:
		// auth.RegisterRoutes(api)
		// users.RegisterRoutes(api)
		_ = api // مؤقتاً عشان الكود يبني وهو لسه فاضي
	}
}