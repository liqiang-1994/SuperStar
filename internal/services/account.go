package services

import (
	"SuperStar/internal/entity"
	"context"
)

type AccountRepo interface {
	CreateUser(ctx context.Context, u *entity.TAccount) (int64, error)
}

type AccountService struct {
	accountRepo AccountRepo
	tm          Transaction
}

func NewAccountService(account AccountRepo, tm Transaction) *AccountService {
	return &AccountService{accountRepo: account, tm: tm}
}

func (s *AccountService) CreateAccount(ctx context.Context, m *entity.Account) (int, error) {
	var (
		err error
		id  int64
	)
	//err = s.tm.ExecTx(ctx, func(ctx context.Context) error {
	//	id, err = s.accountRepo.CreateUser(ctx, m)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return 0, err
	//}
	var req = entity.TAccount{
		Id:       m.Id,
		UserName: m.Name,
	}
	id, err = s.accountRepo.CreateUser(ctx, &req)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
