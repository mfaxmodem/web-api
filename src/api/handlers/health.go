package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfaxmodem/web-api/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("Working!", true, 0))

}
