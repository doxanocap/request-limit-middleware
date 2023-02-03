package handler

import (
	"github.com/gin-gonic/gin"
)

func Welcome(ctx *gin.Context) {
	// ... some very expensive database query
	ctx.JSON(200, gin.H{"message": "Your request were handled", "status": 200})
}
