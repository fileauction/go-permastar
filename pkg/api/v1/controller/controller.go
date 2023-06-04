package controller

import (
	"github.com/kenlabs/permastar/pkg/api/core"
	"github.com/kenlabs/permastar/pkg/option"
)

type Controller struct {
	Core    *core.Core
	Options *option.DaemonOptions
}

func New(core *core.Core, opt *option.DaemonOptions) *Controller {
	return &Controller{
		Core:    core,
		Options: opt,
	}
}
