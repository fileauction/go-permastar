package permastar

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kenlabs/permastar/pkg/api/core"
	"github.com/kenlabs/permastar/pkg/api/types"
	v1 "github.com/kenlabs/permastar/pkg/api/v1"
	"github.com/kenlabs/permastar/pkg/api/v1/controller"
	"github.com/kenlabs/permastar/pkg/option"
	"github.com/kenlabs/permastar/pkg/util/log"
	"net/http"
)

var logger = log.NewSubsystemLogger()

type API struct {
	router     *gin.Engine
	controller *controller.Controller
}

func NewV1HttpAPI(router *gin.Engine, core *core.Core, opt *option.DaemonOptions) *API {
	return &API{
		router:     router,
		controller: controller.New(core, opt),
	}
}

func (a *API) RegisterAPIs() {
	a.registerUserAPI()
	a.registerDataAPI()
}

func HandleError(ctx *gin.Context, err error) {
	var apiErr *v1.Error
	var code = http.StatusBadRequest

	logger.Error(err)
	if errors.As(err, &apiErr) {
		ctx.AbortWithStatusJSON(apiErr.Status(), types.NewErrorResponse(apiErr.Status(), apiErr.Error()))
		return
	}

	ctx.AbortWithStatusJSON(code, types.NewErrorResponse(code, err.Error()))
}
