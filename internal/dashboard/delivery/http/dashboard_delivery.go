package dashboard

import (
	"net/http"

	useCase "edtech-app/internal/dashboard/usecase"
	"edtech-app/internal/middleware"
	"edtech-app/pkg/utils"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	useCase useCase.DashboardUseCase
}

func NewDashboardHandler(useCase useCase.DashboardUseCase) *DashboardHandler {
	return &DashboardHandler{useCase}
}

func (handler *DashboardHandler) Route(r *gin.RouterGroup) {
	dashboardHandler := r.Group("/api/v1")

	dashboardHandler.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		dashboardHandler.GET("/dashboards", handler.GetDataDashboard)
	}
}

func (handler *DashboardHandler) GetDataDashboard(ctx *gin.Context) {
	data := handler.useCase.GetDataDashboard()

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}
