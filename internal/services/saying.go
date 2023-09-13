package services

import (
	"SuperStar/internal/config"
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"context"
	"github.com/go-redis/redis"
)

type SayingRepo interface {
	QueryByParams(ctx context.Context, params *model.PoemListReq) ([]entity.Saying, error)
}

type SayingService struct {
	sayingRepo SayingRepo
	tm         Transaction
	rdb        *redis.Client
	cfg        *config.Config
}

func NewSayingService(saying SayingRepo, tm Transaction, client *redis.Client, cfg *config.Config) *SayingService {
	return &SayingService{sayingRepo: saying, tm: tm, rdb: client, cfg: cfg}
}

func (s *SayingService) QueryByParams(ctx context.Context, params *model.PoemListReq) (interface{}, error) {
	return s.sayingRepo.QueryByParams(ctx, params)
}
