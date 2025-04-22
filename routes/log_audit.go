package routes

import (
	"pelaporan_keuangan/features/log_audit"

	"github.com/gin-gonic/gin"
)

func Log_audit(r *gin.Engine, handler log_audit.Handler) {
	log_audit := r.Group("/log_audit")

	log_audit.GET("", handler.GetLog_audit)
	log_audit.POST("", handler.CreateLog_audit)

	log_audit.GET("/:id", handler.Log_auditDetails)
	log_audit.PUT("/:id", handler.UpdateLog_audit)
	log_audit.DELETE("/:id", handler.DeleteLog_audit)
}
