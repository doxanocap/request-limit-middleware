package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Welcome(ctx *gin.Context) {
	data := map[string]interface{}{}

	if err := ctx.BindJSON(&data); err != nil {
		log.Println("Error in the parsing", err)
		ctx.JSON(400, "Parsing error")
		return
	}

	// ... some very expensive database query
	ctx.JSON(200, gin.H{"message": "Your request were handled", "status": 200})
}
