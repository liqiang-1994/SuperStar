package handlers

import (
	"SuperStar/common"
	"SuperStar/internal/config"
	"SuperStar/internal/middlemare"
	"SuperStar/internal/model"
	"SuperStar/internal/services"
	"github.com/gofiber/fiber/v2"
)

type CircleHandler struct {
	circleSrv *services.CircleService
	cfg       *config.Config
}

func NewCircleHandler(circleService *services.CircleService, cfg *config.Config) *CircleHandler {
	return &CircleHandler{circleSrv: circleService, cfg: cfg}
}

// CreateCircle 创建圈子
// @Summary 创建圈子
// @Description 创建圈子
// @Tags circle
// @Accept json
// @Produce json
// @Param param body object true "param"
// @Security ApiKeyAuth
// @Router /api/circle/create [post]
func (s *CircleHandler) CreateCircle(c *fiber.Ctx) error {
	createCircleReq := &model.CreateCircleReq{}
	if err := c.BodyParser(createCircleReq); err != nil {
		return c.JSON(common.Fail(err))
	}
	user := middlemare.ExtractTokenMetadata(c, s.cfg)
	result, err := s.circleSrv.CreateCircle(c.Context(), createCircleReq, user)
	if err != nil || !result {
		return c.JSON(common.Fail(err))
	}
	return c.JSON(common.Success(true))
}

// QueryCircleDetail 查询圈子详情
// @Summary 查询圈子详情
// @Description 查询圈子详情
// @Tags circle
// @Accept json
// @Produce json
// @Param id path int true "Circle ID"
// @Security ApiKeyAuth
// @Router /api/circle/{id} [get]
func (s *CircleHandler) QueryCircleDetail(c *fiber.Ctx) error {
	circleId := c.Params("id")
	result, err := s.circleSrv.QueryCircleDetail(c.Context(), circleId)
	if err != nil {
		return c.JSON(common.Fail(err))
	}
	return c.JSON(common.Success(result))
}
