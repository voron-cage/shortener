package storage

import "errors"

var (
	DuplicateError = errors.New("duplicate key exists")
	NotFoundError  = errors.New("record not found")
)

type SeqNextResult struct {
	ID uint64
}

type UserTnt struct {
	_msgpack     struct{} `msgpack:",asArray"`
	ID           uint64
	Email        string
	HashPassword string
	Username     string
}

type LinkTnt struct {
	_msgpack    struct{} `msgpack:",asArray"`
	ID          uint64
	UserID      uint64
	OriginalURL string
	ShortURL    string
}
