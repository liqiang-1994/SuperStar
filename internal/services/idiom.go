package services

import (
	"SuperStar/internal/config"
	"SuperStar/internal/entity"
	"SuperStar/internal/model"
	"context"
	"github.com/go-redis/redis"
)

type IdiomRepo interface {
	QueryByParams(ctx context.Context, params *model.PoemListReq) ([]entity.Idiom, error)
}

type IdiomService struct {
	idiomRepo IdiomRepo
	tm        Transaction
	rdb       *redis.Client
	cfg       *config.Config
}

func NewIdiomService(idiom IdiomRepo, tm Transaction, client *redis.Client, cfg *config.Config) *IdiomService {
	return &IdiomService{idiomRepo: idiom, tm: tm, rdb: client, cfg: cfg}
}

func (s *IdiomService) QueryByParams(ctx context.Context, params *model.PoemListReq) (interface{}, error) {
	return s.idiomRepo.QueryByParams(ctx, params)
}
