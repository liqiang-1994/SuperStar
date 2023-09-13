package repos

import (
	"SuperStar/internal/entity"
	"SuperStar/internal/services"
	"context"
)

func NewAccountRepo(db *Data) services.AccountRepo {
	return &accountRepo{db: db}
}

type accountRepo struct {
	db *Data
}

func (a *accountRepo) QueryByPhone(ctx context.Context, phone string) (*entity.User, error) {
	u := &entity.User{}
	result := a.db.DB(ctx).Where("phone = ?", phone).First(u)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return u, nil
}

func (a *accountRepo) CreateUser(ctx context.Context, u *entity.User) (int64, error) {
	u.Id = a.db.GID.Generate().Int64()
	result := a.db.DB(ctx).Create(u)
	return u.Id, result.Error
}

func (a *accountRepo) QueryById(ctx context.Context, id string) (*entity.User, error) {
	u := &entity.User{}
	err := a.db.DB(ctx).First(u, id).Error
	return u, err
}

func (a *accountRepo) UpdateUser(ctx context.Context, u *entity.User) error {
	return a.db.DB(ctx).Save(u).Error
}
