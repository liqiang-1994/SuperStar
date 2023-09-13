package handlers

import (
	"SuperStar/common"
	"SuperStar/internal/config"
	"SuperStar/internal/entity"
	"SuperStar/internal/middlemare"
	"SuperStar/internal/model"
	"SuperStar/internal/services"
	"github.com/gofiber/fiber/v2"
	"time"
)

type TagHandler struct {
	tagSrv *services.TagService
	cfg    *config.Config
}

func NewTagHandler(tagService *services.TagService, cfg *config.Config) *TagHandler {
	return &TagHandler{tagSrv: tagService, cfg: cfg}
}

// CreateTag 新建标签
// @Summary 新建标签
// @Description 新建标签
// @Tags tag
// @Accept json
// @Produce json
// @Param param body object true "param"
// @Security ApiKeyAuth
// @Router /api/tag/create [post]
func (t *TagHandler) CreateTag(c *fiber.Ctx) error {
	createTagReq := &model.CreateTagReq{}
	if err := c.BodyParser(createTagReq); err != nil {
		return c.JSON(common.Fail(err))
	}
	user := middlemare.ExtractTokenMetadata(c, t.cfg)
	tag := &entity.Tag{
		CreateId: user.Id,
		Name:     createTagReq.Name,
	}
	result, err := t.tagSrv.CreateTag(c.UserContext(), tag)
	if err != nil {
		return c.JSON(common.Fail(err))
	}
	return c.JSON(common.Success(result))
}

// QueryAllTag 查询标签列表
// @Summary 查询标签列表
// @Description 查询标签列表
// @Tags tag
// @Accept json
// @Produce json
// @Param param body object true "param"
// @Security ApiKeyAuth
// @Router /api/tag/all [post]
func (t *TagHandler) QueryAllTag(c *fiber.Ctx) error {
	createTagReq := &model.CreateTagReq{}
	if err := c.BodyParser(createTagReq); err != nil {
		return c.JSON(common.Fail(err))
	}
	result, err := t.tagSrv.QueryAllTag(c.UserContext(), createTagReq.Name)
	if err != nil {
		return c.JSON(common.Fail(err))
	}
	return c.JSON(common.Success(result))
}

// CreateTagRel 新建标签关系
// @Summary 新建标签关系
// @Description 新建标签关系
// @Tags tag
// @Accept json
// @Produce json
// @Param param body object true "param"
// @Security ApiKeyAuth
// @Router /api/tag/createTagRel [post]
func (t *TagHandler) CreateTagRel(c *fiber.Ctx) error {
	createTagRelReq := &model.CreateTagRelReq{}
	if err := c.BodyParser(createTagRelReq); err != nil {
		return c.JSON(common.Fail(err))
	}
	user := middlemare.ExtractTokenMetadata(c, t.cfg)
	u := &entity.TagRel{
		TagId:        createTagRelReq.TagId,
		ResourceId:   createTagRelReq.ResourceId,
		ResourceType: createTagRelReq.ResourceType,
		CreateId:     user.Id,
		CreateTime:   time.Now(),
	}
	result, err := t.tagSrv.CreateTagRel(c.UserContext(), u)
	if err != nil {
		return c.JSON(common.Fail(err))
	}
	return c.JSON(common.Success(result))
}
