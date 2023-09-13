package services

import (
	"SuperStar/internal/entity"
	"context"
)

type TagRepo interface {
	CreateTag(ctx context.Context, u *entity.Tag) (bool, error)
	QueryAllTag(ctx context.Context, name string) ([]*entity.Tag, error)
	CreateTagRel(ctx context.Context, m *entity.TagRel) (bool, error)
}

type TagService struct {
	tagRepo TagRepo
	tm      Transaction
}

func NewTagService(tag TagRepo, tm Transaction) *TagService {
	return &TagService{tagRepo: tag, tm: tm}
}

func (s *TagService) CreateTag(ctx context.Context, u *entity.Tag) (bool, error) {
	return s.tagRepo.CreateTag(ctx, u)
}

func (s *TagService) QueryAllTag(ctx context.Context, name string) ([]*entity.Tag, error) {
	return s.tagRepo.QueryAllTag(ctx, name)
}

func (s *TagService) CreateTagRel(ctx context.Context, u *entity.TagRel) (bool, error) {
	return s.tagRepo.CreateTagRel(ctx, u)
}
