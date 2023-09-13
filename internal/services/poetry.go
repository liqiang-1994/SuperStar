package services

import (
	"SuperStar/internal/config"
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"context"
	"github.com/go-redis/redis"
)

type PoetryRepo interface {
	QueryByParams(ctx context.Context, params *model.PoemListReq) ([]entity.Poetry, error)
}

func NewPoetryService(poetry PoetryRepo, tm Transaction, client *redis.Client, cfg *config.Config) *PoetryService {
	return &PoetryService{poetryRepo: poetry, tm: tm, rdb: client, cfg: cfg}
}

type PoetryService struct {
	poetryRepo PoetryRepo
	tm         Transaction
	rdb        *redis.Client
	cfg        *config.Config
}

func (s *PoetryService) QueryByParams(ctx context.Context, params *model.PoemListReq) (interface{}, error) {
	return s.poetryRepo.QueryByParams(ctx, params)
}
