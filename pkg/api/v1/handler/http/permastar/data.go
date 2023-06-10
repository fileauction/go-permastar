package permastar

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kenlabs/permastar/pkg/api/types"
	"github.com/kenlabs/permastar/pkg/api/v1/model"
	"io"
	"path"
)

func (a *API) registerDataAPI() {
	a.router.MaxMultipartMemory = 8 << 20 // 8MB
	dataAPI := a.router.Group("/data")
	{
		dataAPI.GET("/dir", a.listDir)
		dataAPI.POST("/dir", a.createDir)
		dataAPI.DELETE("/dir", a.deleteDir)
		dataAPI.GET("/file", a.downloadFile)
		dataAPI.POST("/file", a.uploadFile)
		dataAPI.DELETE("/file", a.deleteFile)
		dataAPI.GET("/file/stat", a.statFile)
	}
}

func (a *API) listDir(ctx *gin.Context) {
	accountAddr, err := getAccountAddress(ctx)
	if err != nil {
		HandleError(ctx, err)
		return
	}
	path := ctx.Param("path")
	entries, err := a.controller.ListDir(ctx, path, accountAddr)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(200, types.NewOKResponse("OK", entries))
}

func (a *API) createDir(ctx *gin.Context) {
	accountAddr, err := getAccountAddress(ctx)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		HandleError(ctx, err)
		return
	}
	dirInfo := &model.DirInfo{}
	if err = json.Unmarshal(reqBody, dirInfo); err != nil {
		HandleError(ctx, err)
		return
	}
	newDirInfo, err := a.controller.Core.IPFSNodeAPI.CreateDir(context.Background(), dirInfo.Path, accountAddr)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(200, types.NewOKResponse("OK", newDirInfo))
}

func (a *API) deleteDir(ctx *gin.Context) {
	accountAddr, err := getAccountAddress(ctx)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		HandleError(ctx, err)
		return
	}
	dirInfo := &model.DirInfo{}
	if err = json.Unmarshal(reqBody, dirInfo); err != nil {
		HandleError(ctx, err)
		return
	}

	if err = a.controller.Core.IPFSNodeAPI.DeleteDir(context.Background(), dirInfo.Path, accountAddr); err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(200, types.NewOKResponse("OK", nil))
}

func (a *API) downloadFile(ctx *gin.Context) {
	accountAddr, err := getAccountAddress(ctx)
	if err != nil {
		HandleError(ctx, err)
		return
	}
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		HandleError(ctx, err)
		return
	}
	fileInfo := &model.FileInfo{}
	if err = json.Unmarshal(reqBody, fileInfo); err != nil {
		HandleError(ctx, err)
		return
	}
	data, err := a.controller.Core.IPFSNodeAPI.DownloadFile(context.Background(), fileInfo.Path, accountAddr)
	if err != nil {
		HandleError(ctx, err)
		return
	}
	_, fileName := path.Split(fileInfo.Path)
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/text/plain")
	ctx.Header("Accept-Length", fmt.Sprintf("%d", len(data)))
	_, err = ctx.Writer.Write(data)
	if err != nil {
		HandleError(ctx, err)
		return
	}
}

func (a *API) uploadFile(ctx *gin.Context) {
	accountAddr, err := getAccountAddress(ctx)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	multiForm, err := ctx.MultipartForm()
	if err != nil {
		HandleError(ctx, err)
		return
	}
	dstPaths, exists := multiForm.Value["path"]
	if !exists {
		HandleError(ctx, errors.New("key path does not exist in multiform"))
		return
	}
	dstPath := dstPaths[0]
	files, exists := multiForm.File["files"]
	if !exists {
		HandleError(ctx, errors.New("key files does not exist in multiform"))
		return
	}
	for _, file := range files {
		fo, err := file.Open()
		if err != nil {
			HandleError(ctx, err)
			return
		}
		fo.Close()
		err = a.controller.Core.IPFSNodeAPI.CreateFile(context.Background(), path.Join(dstPath, file.Filename), fo, accountAddr)
		if err != nil {
			HandleError(ctx, err)
			return
		}
	}

	ctx.JSON(200, types.NewOKResponse("OK", nil))
}

func (a *API) deleteFile(ctx *gin.Context) {
	accountAddr, err := getAccountAddress(ctx)
	if err != nil {
		HandleError(ctx, err)
		return
	}

	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		HandleError(ctx, err)
		return
	}
	fileInfo := &model.FileInfo{}
	if err = json.Unmarshal(reqBody, fileInfo); err != nil {
		HandleError(ctx, err)
		return
	}

	if err = a.controller.Core.IPFSNodeAPI.DeleteFile(context.Background(), fileInfo.Path, accountAddr); err != nil {
		HandleError(ctx, err)
		return
	}

	ctx.JSON(200, types.NewOKResponse("OK", nil))
}

func (a *API) statFile(ctx *gin.Context) {
	accountAddr, err := getAccountAddress(ctx)
	if err != nil {
		HandleError(ctx, err)
		return
	}
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		HandleError(ctx, err)
		return
	}
	fileInfo := &model.FileInfo{}
	if err = json.Unmarshal(reqBody, fileInfo); err != nil {
		HandleError(ctx, err)
		return
	}

	fileStat, err := a.controller.Core.IPFSNodeAPI.GetFileStat(context.Background(), fileInfo.Path, accountAddr)
	if err != nil {
		HandleError(ctx, err)
		return
	}
	ctx.JSON(200, types.NewOKResponse("OK", fileStat))
}

func getAccountAddress(ctx *gin.Context) (string, error) {
	accountAddr := ctx.Request.Header.Get("AccountAddr")
	if len(accountAddr) == 0 {
		return "", errors.New("account address is empty")
	}

	return accountAddr, nil
}

//
//import (
//	"errors"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/kenlabs/permastar/pkg/api/types"
//	"net/http"
//)
//
//func (a *API) registerDataAPI() {
//	a.router.MaxMultipartMemory = 8 << 20 // 8MB
//	dataAPI := a.router.Group("/data")
//	{
//		dataAPI.POST("/file", a.uploadFileToCollection)
//		dataAPI.DELETE("/file", a.removeFileFromCollection)
//	}
//}
//
//func (a *API) uploadFileToCollection(ctx *gin.Context) {
//	form, err := ctx.MultipartForm()
//	if err != nil {
//		logger.Errorf("get multipartform failed: %v", err)
//		HandleError(ctx, err)
//		return
//	}
//	// save the first file to local filesystem
//	files, exists := form.File["files"]
//	if !exists {
//		errStr := "key \"files\" does not exist in multipart-form"
//		logger.Errorf(errStr)
//		HandleError(ctx, errors.New(errStr))
//		return
//	}
//	if len(files) == 0 {
//		errStr := "key \"files\" does not contain any file"
//		logger.Errorf(errStr)
//		HandleError(ctx, errors.New(errStr))
//		return
//	}
//	tmpFilePath := "/tmp/permastar/" + files[0].Filename
//	err = ctx.SaveUploadedFile(files[0], tmpFilePath)
//	if err != nil {
//		errStr := fmt.Sprintf("save file to local failed: %v", err)
//		logger.Errorf(errStr)
//		HandleError(ctx, errors.New(errStr))
//		return
//	}
//
//	// upload & pin file to estuary
//	dstPaths, exists := form.Value["dstPath"]
//	if !exists {
//		errStr := "key \"dstPath\" does not exist in multipart-form"
//		logger.Errorf(errStr)
//		HandleError(ctx, errors.New(errStr))
//		return
//	}
//	if len(dstPaths) == 0 {
//		errStr := "key \"dstPath\" does not contain any value"
//		logger.Errorf(errStr)
//		HandleError(ctx, errors.New(errStr))
//		return
//	}
//	dstPath := dstPaths[0]
//	userTokens, exists := form.Value["userToken"]
//	if !exists {
//		errStr := "key \"userToken\" does not exist in multipart-form"
//		logger.Errorf(errStr)
//		HandleError(ctx, errors.New(errStr))
//		return
//	}
//	if len(userTokens) == 0 {
//		errStr := "key \"userToken\" does not contain any value"
//		logger.Errorf(errStr)
//		HandleError(ctx, errors.New(errStr))
//		return
//	}
//	userToken := userTokens[0]
//	colIDs, exists := form.Value["collectionID"]
//	if !exists {
//		errStr := "key \"collectionID\" does not exist in multipart-form"
//		logger.Errorf(errStr)
//		HandleError(ctx, errors.New(errStr))
//		return
//	}
//	if len(colIDs) == 0 {
//		errStr := "key \"collectionID\" does not contain any value"
//		logger.Errorf(errStr)
//		HandleError(ctx, errors.New(errStr))
//		return
//	}
//	colID := colIDs[0]
//
//	err = a.controller.UploadFileToCollection(ctx, tmpFilePath, dstPath, colID, userToken)
//	if err != nil {
//		logger.Errorf(err.Error())
//		HandleError(ctx, err)
//	}
//
//	ctx.JSON(http.StatusOK, types.NewOKResponse("OK", nil))
//	//file
//	//err := ctx.SaveUploadedFile(formFile, "/tmp/permastar/"+formFile.Filename)
//	//if err != nil {
//	//	logger.Errorf("upload file %v failed: %v", formFile.Filename, err)
//	//	HandleError(ctx, err)
//	//	return
//	//}
//	//
//	//uploadPath := ctx.PostForm()
//}
//
//func (a *API) removeFileFromCollection(ctx *gin.Context) {
//
//}
