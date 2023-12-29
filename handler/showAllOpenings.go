package handler

import (
	"net/http"

	"github.com/devgugga/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowAllOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error while listing openings")
		return
	}

	sendSuccess(ctx, "show-all-openings", openings)
}
