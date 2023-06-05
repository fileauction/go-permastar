package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/kenlabs/permastar/pkg/api/core"
	"github.com/kenlabs/permastar/pkg/api/middleware"
	"github.com/kenlabs/permastar/pkg/api/v1/handler/http/permastar"

	"github.com/kenlabs/permastar/pkg/option"
)

func NewHttpRouter(core *core.Core, opt *option.DaemonOptions) *gin.Engine {
	httpRouter := gin.New()
	httpRouter.Use(middleware.WithLoggerFormatter())

	v1HttpAPI := permastar.NewV1HttpAPI(httpRouter, core, opt)
	v1HttpAPI.RegisterAPIs()

	return httpRouter
}
