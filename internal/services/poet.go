package services

import (
	"SuperStar/internal/config"
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"context"
	"github.com/go-redis/redis"
)

type PoetRepo interface {
	QueryByParams(ctx context.Context, params *model.PoemListReq) ([]entity.Poet, error)
}

type PoetService struct {
	poetRepo PoetRepo
	tm       Transaction
	rdb      *redis.Client
	cfg      *config.Config
}

func NewPoetService(poet PoetRepo, tm Transaction, client *redis.Client, cfg *config.Config) *PoetService {
	return &PoetService{poetRepo: poet, tm: tm, rdb: client, cfg: cfg}
}

func (s *PoetService) QueryByParams(ctx context.Context, params *model.PoemListReq) (interface{}, error) {
	return s.poetRepo.QueryByParams(ctx, params)
}
