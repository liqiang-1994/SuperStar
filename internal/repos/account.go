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

func (a *accountRepo) CreateUser(ctx context.Context, u *entity.TAccount) (int64, error) {
	u.Id = a.db.GID.Generate().Int64()
	result := a.db.DB(ctx).Create(u)
	return u.Id, result.Error
}
