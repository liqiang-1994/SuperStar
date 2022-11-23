package services

import (
	"context"
	"github.com/google/wire"
)

var Provider = wire.NewSet(NewAccountService)

// Transaction 新增事务接口方法
type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}
