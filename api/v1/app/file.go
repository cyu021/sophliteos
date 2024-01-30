package app

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sophliteos/logger"

	"strconv"
	"strings"

	"sophliteos/pkg/buserr"
	"sophliteos/pkg/dto/request"
	"sophliteos/pkg/handle"
	"sophliteos/pkg/utils/files"

	"sophliteos/pkg/dto"

	"sophliteos/global"

	"github.com/gin-gonic/gin"
)

type FileApi struct{}

// 获取文件列表
func (b *FileApi) ListFiles(c *gin.Context) {
	var req request.FileOption
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	files, err := fileService.GetFileList(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Success(files))
}

// 分页获取上传文件
func (b *FileApi) SearchUploadWithPage(c *gin.Context) {
	var req request.SearchUploadWithPage
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	if err := global.VALID.Struct(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	total, files, err := fileService.SearchUploadWithPage(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	c.JSON(http.StatusOK, handle.Success(dto.PageResult{
		Items: files,
		Total: total,
	}))
}

// 加载文件树
func (b *FileApi) GetFileTree(c *gin.Context) {
	var req request.FileOption
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	tree, err := fileService.GetFileTree(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Success(tree))
}

// 创建文件/文件夹
func (b *FileApi) CreateFile(c *gin.Context) {
	var req request.FileCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	err := fileService.Create(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 删除文件/文件夹
func (b *FileApi) DeleteFile(c *gin.Context) {
	var req request.FileDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	err := fileService.Delete(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 批量删除文件/文件夹
func (b *FileApi) BatchDeleteFile(c *gin.Context) {
	var req request.FileBatchDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	err := fileService.BatchDelete(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 修改文件权限
func (b *FileApi) ChangeFileMode(c *gin.Context) {
	var req request.FileCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	err := fileService.ChangeMode(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 修改文件用户/组
func (b *FileApi) ChangeFileOwner(c *gin.Context) {
	var req request.FileRoleUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	if err := fileService.ChangeOwner(req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 压缩文件
func (b *FileApi) CompressFile(c *gin.Context) {
	var req request.FileCompress
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	err := fileService.Compress(req)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 解压文件
func (b *FileApi) DeCompressFile(c *gin.Context) {
	var req request.FileDeCompress
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	err := fileService.DeCompress(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 获取文件内容
func (b *FileApi) GetContent(c *gin.Context) {
	var req request.FileOption
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	info, err := fileService.GetContent(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Success(info))
}

// 更新文件内容
func (b *FileApi) SaveContent(c *gin.Context) {
	var req request.FileEdit
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	if err := fileService.SaveContent(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 检测文件是否存在
func (b *FileApi) CheckFile(c *gin.Context) {
	var req request.FilePathCheck
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}
	if err := global.VALID.Struct(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	if _, err := os.Stat(req.Path); err != nil && os.IsNotExist(err) {
		c.JSON(http.StatusOK, handle.Success(false))
		return
	}

	c.JSON(http.StatusOK, handle.Success(true))
}

// 上传文件
func (b *FileApi) UploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	files := form.File["file"]
	paths := form.Value["path"]
	if len(paths) == 0 || !strings.Contains(paths[0], "/") {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:error paths in request"))
		return
	}
	dir := path.Dir(paths[0])
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			errStr := fmt.Sprintf("mkdir %s failed, err: %v", dir, err)
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+errStr))
			return
		}
	}
	success := 0
	failures := make(buserr.MultiErr)
	for _, file := range files {
		if err := c.SaveUploadedFile(file, path.Join(paths[0], file.Filename)); err != nil {
			e := fmt.Errorf("upload [%s] file failed, err: %v", file.Filename, err)
			failures[file.Filename] = e
			logger.Error("%v", e)
			continue
		}
		success++
	}
	if success == 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, failures.Error()))
	} else {
		c.JSON(http.StatusOK, handle.Ok())
	}
}

// 分片上传文件
func (b *FileApi) UploadChunkFiles(c *gin.Context) {
	var err error
	fileForm, err := c.FormFile("chunk")
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	uploadFile, err := fileForm.Open()
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	chunkIndex, err := strconv.Atoi(c.PostForm("chunkIndex"))
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	chunkCount, err := strconv.Atoi(c.PostForm("chunkCount"))
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	fileOp := files.NewFileOp()
	tmpDir := path.Join("/data/sophliteos", "upload")
	if !fileOp.Stat(tmpDir) {
		if err := fileOp.CreateDir(tmpDir, 0755); err != nil {

			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
			return
		}
	}
	filename := c.PostForm("filename")
	fileDir := filepath.Join(tmpDir, filename)
	if chunkIndex == 0 {
		if fileOp.Stat(fileDir) {
			_ = fileOp.DeleteDir(fileDir)
		}
		_ = os.MkdirAll(fileDir, 0755)
	}
	filePath := filepath.Join(fileDir, filename)

	defer func() {
		if err != nil {
			_ = os.Remove(fileDir)
		}
	}()
	var (
		emptyFile *os.File
		chunkData []byte
	)

	emptyFile, err = os.Create(filePath)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	defer emptyFile.Close()

	chunkData, err = io.ReadAll(uploadFile)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"))
		return
	}

	chunkPath := filepath.Join(fileDir, fmt.Sprintf("%s.%d", filename, chunkIndex))
	err = os.WriteFile(chunkPath, chunkData, 0644)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"))
		return
	}

	if chunkIndex+1 == chunkCount {
		err = mergeChunks(filename, fileDir, c.PostForm("path"), chunkCount)
		if err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"))
			return
		}
		c.JSON(http.StatusOK, handle.Success(true))
	} else {
		return
	}
}

func mergeChunks(fileName string, fileDir string, dstDir string, chunkCount int) error {
	if _, err := os.Stat(path.Dir(dstDir)); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(path.Dir(dstDir), os.ModePerm); err != nil {
			return err
		}
	}

	targetFile, err := os.Create(filepath.Join(dstDir, fileName))
	if err != nil {
		return err
	}
	defer targetFile.Close()

	for i := 0; i < chunkCount; i++ {
		chunkPath := filepath.Join(fileDir, fmt.Sprintf("%s.%d", fileName, i))
		chunkData, err := os.ReadFile(chunkPath)
		if err != nil {
			return err
		}
		_, err = targetFile.Write(chunkData)
		if err != nil {
			return err
		}
	}

	return files.NewFileOp().DeleteDir(fileDir)
}

// 修改文件名称
func (b *FileApi) ChangeFileName(c *gin.Context) {
	var req request.FileRename
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	if err := fileService.ChangeName(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 下载远端文件
func (b *FileApi) WgetFile(c *gin.Context) {
	var req request.FileWget
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	err := fileService.Wget(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}

	c.JSON(http.StatusOK, handle.Ok())
}

// 移动文件
func (b *FileApi) MoveFile(c *gin.Context) {
	var req request.FileMove
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	if err := fileService.MvFile(req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Ok())
}

// 下载文件
func (b *FileApi) Download(c *gin.Context) {
	filePath := c.Query("path")
	file, err := os.Open(filePath)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
	}
	info, _ := file.Stat()
	c.Header("Content-Length", strconv.FormatInt(info.Size(), 10))
	c.Header("Content-Disposition", "attachment; filename*=utf-8''"+url.PathEscape(info.Name()))
	http.ServeContent(c.Writer, c.Request, info.Name(), info.ModTime(), file)
}

// 分片下载下载文件
func (b *FileApi) DownloadChunkFiles(c *gin.Context) {
	var req request.FileChunkDownload
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	fileOp := files.NewFileOp()
	if !fileOp.Stat(req.Path) {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "文件不存在"))
		return
	}
	filePath := req.Path
	fstFile, err := fileOp.OpenFile(filePath)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	info, err := fstFile.Stat()
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	if info.IsDir() {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "不支持下载文件夹"))
		return
	}

	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", req.Name))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Header().Set("Content-Length", strconv.FormatInt(info.Size(), 10))
	c.Writer.Header().Set("Accept-Ranges", "bytes")

	if c.Request.Header.Get("Range") != "" {
		rangeHeader := c.Request.Header.Get("Range")
		rangeArr := strings.Split(rangeHeader, "=")[1]
		rangeParts := strings.Split(rangeArr, "-")

		startPos, _ := strconv.ParseInt(rangeParts[0], 10, 64)

		var endPos int64
		if rangeParts[1] == "" {
			endPos = info.Size() - 1
		} else {
			endPos, _ = strconv.ParseInt(rangeParts[1], 10, 64)
		}

		c.Writer.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", startPos, endPos, info.Size()))
		c.Writer.WriteHeader(http.StatusPartialContent)

		buffer := make([]byte, 1024*1024)
		file, err := os.Open(filePath)
		if err != nil {

			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
			return
		}
		defer file.Close()

		_, _ = file.Seek(startPos, 0)
		reader := io.LimitReader(file, endPos-startPos+1)
		_, err = io.CopyBuffer(c.Writer, reader, buffer)
		if err != nil {

			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
			return
		}
	} else {
		c.File(filePath)
	}
}

// 获取文件夹大小
func (b *FileApi) Size(c *gin.Context) {
	var req request.DirSizeReq
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	res, err := fileService.DirSize(req)
	if err != nil {

		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "请求失败:"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, handle.Success(res))
}
