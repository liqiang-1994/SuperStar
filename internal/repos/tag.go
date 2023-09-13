package repos

import (
	"SuperStar/internal/entity"
	"SuperStar/internal/services"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

func NewTagRepo(db *Data) services.TagRepo {
	return &tagRepo{db: db}
}

type tagRepo struct {
	db *Data
}

func (t *tagRepo) QueryAllTag(ctx context.Context, name string) ([]*entity.Tag, error) {
	var tags []*entity.Tag
	sql := t.db.DB(ctx).Order("id asc")
	if name != "" {
		sql = sql.Where("name LIKE ?", "%"+name+"%")
	}
	err := sql.Find(&tags).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return tags, err
}

func (t *tagRepo) CreateTag(ctx context.Context, u *entity.Tag) (bool, error) {
	u.Id = int64(t.db.GID.Generate())
	u.CreateTime = time.Now()
	result := t.db.DB(ctx).Create(u)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (t *tagRepo) CreateTagRel(ctx context.Context, m *entity.TagRel) (bool, error) {
	m.Id = int64(t.db.GID.Generate())
	result := t.db.DB(ctx).Create(m)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
