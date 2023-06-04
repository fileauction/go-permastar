package permastar

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kenlabs/permastar/pkg/api/types"
	v1 "github.com/kenlabs/permastar/pkg/api/v1"
	"io"
	"net/http"
)

func (a *API) registerUserAPI() {
	a.router.POST("/user", a.userRegister)
}

func (a *API) userRegister(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		HandleError(ctx, v1.NewError(v1.InternalServerError, http.StatusInternalServerError))
		return
	}
	if err = a.controller.RegisterUser(context.Background(), bodyBytes); err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, types.NewOKResponse("OK", nil))
}
