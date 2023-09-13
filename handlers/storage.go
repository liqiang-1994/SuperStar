package handlers

import (
	"SuperStar/common"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type StorageHandler struct {
}

func NewStorageHandler() *StorageHandler {
	return &StorageHandler{}
}

// UploadAvatar 上传头像
// @Summary 上传头像
// @Description 上传头像
// @Tags storage
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// // @Security ApiKeyAuth
// @Router /api/storage/avatar/upload [post]
func (*StorageHandler) UploadAvatar(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(common.Fail(err))
	}
	log.Info(file.Filename)
	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		log.Error(err)
	}
	return c.JSON(common.Success(true))
}

// Download 下载文件
// @Summary 下载文件
// @Description 下载文件
// @Tags storage
// @Param filename query string true "file name"
// // @Security ApiKeyAuth
// @Router /api/storage/download [get]
func (*StorageHandler) Download(c *fiber.Ctx) error {
	return c.Download("./docs/docs.go", "focs.go")
}
