package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfaxmodem/web-api/src/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
// @security AuthBearer
func (h *HealthHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse("Working!", true, 0))

}
