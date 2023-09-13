package handlers

import (
	"SuperStar/common"
	"SuperStar/internal/model"
	"SuperStar/internal/services"
	"github.com/gofiber/fiber/v2"
)

type PoemHandler struct {
	poetSrv   *services.PoetService
	poetrySrv *services.PoetryService
	idiomSrv  *services.IdiomService
	sayingSrv *services.SayingService
}

func NewPoemHandler(
	poetSrv *services.PoetService,
	poetrySrv *services.PoetryService,
	idiomSrv *services.IdiomService,
	sayingSrv *services.SayingService) *PoemHandler {
	return &PoemHandler{poetSrv: poetSrv, poetrySrv: poetrySrv, idiomSrv: idiomSrv, sayingSrv: sayingSrv}
}

// QueryPoemList 诗集首页列表
// @Summary 诗集首页列表
// @Description 诗集首页列表
// @Tags poem
// @Accept json
// @Produce json
// @Param param body object true "param"
// @Security ApiKeyAuth
// @Router /api/poem/list [post]
func (s *PoemHandler) QueryPoemList(c *fiber.Ctx) error {
	poemListReq := &model.PoemListReq{}
	if err := c.BodyParser(poemListReq); err != nil {
		return c.JSON(common.Fail(err))
	}
	var result interface{}
	var err error
	if poemListReq.Type == 0 {
		result, err = s.poetSrv.QueryByParams(c.UserContext(), poemListReq)
		if err != nil {
			return c.JSON(common.Fail(err))
		}
	} else if poemListReq.Type == 1 {
		result, err = s.poetrySrv.QueryByParams(c.UserContext(), poemListReq)
		if err != nil {
			return c.JSON(common.Fail(err))
		}
	} else if poemListReq.Type == 2 {
		result, err = s.idiomSrv.QueryByParams(c.UserContext(), poemListReq)
		if err != nil {
			return c.JSON(common.Fail(err))
		}
	} else if poemListReq.Type == 3 {
		result, err = s.sayingSrv.QueryByParams(c.UserContext(), poemListReq)
		if err != nil {
			return c.JSON(common.Fail(err))
		}
	}
	return c.JSON(common.Success(result))
}
