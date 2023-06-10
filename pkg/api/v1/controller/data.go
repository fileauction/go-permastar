package controller

import (
	"github.com/gin-gonic/gin"
	ipfs_api "github.com/ipfs/go-ipfs-api"
)

func (c *Controller) ListDir(ctx *gin.Context, path string, accountAddr string) ([]*ipfs_api.MfsLsEntry, error) {
	return c.Core.IPFSNodeAPI.ListFilesInPath(ctx, path, accountAddr)
}

//func (c *Controller) CreateDir(ctx *gin.Context, path string, accountAddr string) {
//
//}
//
//func (c *Controller) DeleteDir(ctx *gin.Context) {
//
//}
//
//func (c *Controller) ()  {
//
//}
