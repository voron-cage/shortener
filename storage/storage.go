package storage

import (
	"context"
	"fmt"
	"github.com/tarantool/go-tarantool/v2"
	"go.uber.org/zap"
)

type TarantoolStorage struct {
	conn   *tarantool.Connection
	logger *zap.Logger
}

func NewTarantoolStorage(cfg *TarantoolConfig) *TarantoolStorage {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectTimeout.Duration)
	defer cancel()
	dialer := tarantool.NetDialer{
		Address: cfg.ListenAddress,
		User:    cfg.User,
	}
	opts := tarantool.Opts{
		Timeout: cfg.ResponseTimeout.Duration,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {
		fmt.Println("Connection refused:", err)
		return nil
	}
	return &TarantoolStorage{conn: conn, logger: zap.L()}
}

type TNTStorage interface {
	GetAuthUserIDSeqNext() (uint64, error)
	InsertUser(user *UserTnt) error
	GetUser(email string) (*UserTnt, error)
	InsertLink(tntLink *LinkTnt) error
	GetLinkSeqIDNext() (uint64, error)
	GetByShortURL(shortURL string) (*LinkTnt, error)
}
