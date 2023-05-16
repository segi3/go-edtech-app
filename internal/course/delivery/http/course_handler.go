package course

import (
	"net/http"

	dto "edtech-app/internal/course/dto"
	usecase "edtech-app/internal/course/usecase"
	"edtech-app/internal/middleware"
	"edtech-app/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	useCase usecase.CourseUseCase
}

func NewCourseHandler(useCase usecase.CourseUseCase) *CourseHandler {
	return &CourseHandler{useCase}
}

func (handler *CourseHandler) Route(r *gin.RouterGroup) {
	courseHandler := r.Group("/api/v1")

	courseHandler.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		courseHandler.POST("/courses", handler.Create)
	}
}
func (handler *CourseHandler) Create(ctx *gin.Context) {
	var input dto.CourseBindingRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)
	input.CreatedBy = user.ID

	data, err := handler.useCase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", data))
}
