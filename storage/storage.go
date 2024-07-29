package storage

import (
	"context"
	"github.com/tarantool/go-tarantool/v2"
	"go.uber.org/zap"
)

type TarantoolStorage struct {
	conn   *tarantool.Connection
	logger *zap.Logger
}

func NewTarantoolStorage(cfg *TarantoolConfig) *TarantoolStorage {
	logger := zap.L()
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
		logger.Fatal("Connection refused:", zap.Error(err))
	}
	return &TarantoolStorage{conn: conn, logger: logger}
}

type TNTStorage interface {
	GetAuthUserIDSeqNext() (uint64, error)
	InsertUser(user *UserTnt) error
	GetUser(email string) (*UserTnt, error)
	InsertLink(tntLink *LinkTnt) error
	GetLinkSeqIDNext() (uint64, error)
	GetByShortURL(shortURL string) (*LinkTnt, error)
}
