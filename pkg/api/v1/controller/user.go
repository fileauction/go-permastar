package controller

import (
	"context"
	"github.com/kenlabs/permastar/pkg/api/v1/model"
)

func (c *Controller) RegisterUser(ctx context.Context, data []byte) error {
	registerRequest, err := model.ReadRegisterRequest(data)
	if err != nil {
		return err
	}
	return c.Core.IPFSNodeAPI.NewRootDIRForUser(ctx, registerRequest.AccountAddress)
}
