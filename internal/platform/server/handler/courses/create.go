package courses

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "gohex.com/m/internal/platform"
)

type createRequest struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		course := mooc.NewCourse(req.Id, req.Name, req.Duration)
		if err := Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}

func Save(ctx context.Context, course mooc.Course) error {
	return nil
}