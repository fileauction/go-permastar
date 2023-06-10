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
	a.router.DELETE("/user", a.userDelete)
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

func (a *API) userDelete(ctx *gin.Context) {
	accountAddr, err := getAccountAddress(ctx)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	if err = a.controller.Core.IPFSNodeAPI.DeleteDir(ctx, "", accountAddr); err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(200, types.NewOKResponse("OK", nil))
}
