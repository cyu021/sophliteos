package app

import (
	v1 "sophliteos/api/v1"

	"github.com/gin-gonic/gin"
)

type FileRouter struct{}

func (s *FileRouter) InitFileRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	fileRouter := Router.Group("api/files")
	baseApi := v1.ApiGroupApp.AppGroup.FileApi
	{
		fileRouter.POST("/search", baseApi.ListFiles)
		fileRouter.POST("/upload/search", baseApi.SearchUploadWithPage)
		fileRouter.POST("/tree", baseApi.GetFileTree)
		fileRouter.POST("", baseApi.CreateFile)
		fileRouter.POST("/del", baseApi.DeleteFile)
		fileRouter.POST("/batch/del", baseApi.BatchDeleteFile)
		fileRouter.POST("/mode", baseApi.ChangeFileMode)
		fileRouter.POST("/owner", baseApi.ChangeFileOwner)
		fileRouter.POST("/compress", baseApi.CompressFile)
		fileRouter.POST("/decompress", baseApi.DeCompressFile)
		fileRouter.POST("/content", baseApi.GetContent)
		fileRouter.POST("/save", baseApi.SaveContent)
		fileRouter.POST("/check", baseApi.CheckFile)
		fileRouter.POST("/upload", baseApi.UploadFiles)
		fileRouter.POST("/chunkupload", baseApi.UploadChunkFiles)
		fileRouter.POST("/rename", baseApi.ChangeFileName)
		// fileRouter.POST("/wget", baseApi.WgetFile)
		fileRouter.POST("/move", baseApi.MoveFile)
		fileRouter.GET("/download", baseApi.Download)
		fileRouter.POST("/chunkdownload", baseApi.DownloadChunkFiles)
		fileRouter.POST("/size", baseApi.Size)

	}

	return fileRouter
}
