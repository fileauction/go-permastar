package ipfs_shell

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/go-cid"
	ipfs_api "github.com/ipfs/go-ipfs-api"
	"github.com/kenlabs/permastar/pkg/api/v1/model"
	"github.com/kenlabs/permastar/pkg/util/log"
	"io"
	"path"
)

var logger = log.NewSubsystemLogger()

type API struct {
	Shell *ipfs_api.Shell
}

func NewIPFSNodeShell(gateway string) *API {
	return &API{ipfs_api.NewShell(gateway)}
}

func (a *API) NewRootDIRForUser(ctx context.Context, accountAddr string) error {
	return a.Shell.FilesMkdir(ctx, "/"+accountAddr)
}

func (a *API) CreateDir(ctx context.Context, path string, accountAddr string) (*model.DirInfo, error) {
	if err := a.Shell.FilesMkdir(ctx, "/"+accountAddr+path); err != nil {
		return nil, err
	}
	c, err := a.GetObjectCid(ctx, "/"+accountAddr+path)
	if err != nil {
		return nil, err
	}
	return &model.DirInfo{
		Path: path,
		Cid:  c.String(),
	}, nil
}

func (a *API) DeleteDir(ctx context.Context, path string, accountAddr string) error {
	return a.Shell.FilesRm(ctx, "/"+accountAddr+path, true)
}

func (a *API) DownloadFile(ctx context.Context, dstPath string, accountAddr string) ([]byte, error) {
	reader, err := a.Shell.FilesRead(ctx, path.Join("/", accountAddr, dstPath), ipfs_api.FilesRead.Offset(0))
	if err != nil {
		return nil, err
	}
	respBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}

func (a *API) CreateFile(ctx context.Context, dstPath string, data io.Reader, accountAddr string) error {
	return a.Shell.FilesWrite(ctx, path.Join("/", accountAddr, dstPath), data, ipfs_api.FilesWrite.Create(true))
}

func (a *API) DeleteFile(ctx context.Context, dstPath string, accountAddr string) error {
	filePath := path.Join("/", accountAddr, dstPath)
	fileInfo, err := a.Shell.FilesStat(ctx, filePath)
	if err != nil {
		return err
	}
	if fileInfo.Type == "directory" {
		return errors.New(fmt.Sprintf("object %s is a directory, the api called can only delete file", dstPath))
	}
	return a.Shell.FilesRm(ctx, path.Join("/", accountAddr, dstPath), true)
}

func (a *API) ListFilesInPath(ctx context.Context, path string, accountAddr string) ([]*ipfs_api.MfsLsEntry, error) {
	return a.Shell.FilesLs(ctx, "/"+accountAddr+path, ipfs_api.FilesLs.Stat(true))
}

func (a *API) GetObjectCid(ctx context.Context, path string) (cid.Cid, error) {
	entry, err := a.Shell.FilesStat(ctx, path)
	if err != nil {
		return cid.Undef, err
	}
	cidEntry, err := cid.Parse(entry.Hash)
	if err != nil {
		return cid.Undef, err
	}
	return cidEntry, nil
}

func (a *API) GetFileStat(ctx context.Context, dstPath string, accountAddr string) (*model.FileInfo, error) {
	fileStat, err := a.Shell.FilesStat(ctx, path.Join("/", accountAddr, dstPath), ipfs_api.FilesStat.WithLocal(true))
	if err != nil {
		return nil, err
	}

	return &model.FileInfo{
		Path: dstPath,
		Cid:  fileStat.Hash,
		Size: int(fileStat.Size),
	}, nil
}
