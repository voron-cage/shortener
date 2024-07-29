package storage

import (
	"github.com/tarantool/go-tarantool/v2"
	"strings"
)

func (s *TarantoolStorage) GetAuthUserIDSeqNext() (uint64, error) {
	r := SeqNextResult{}
	err := s.conn.Do(
		tarantool.NewCallRequest("box.sequence.auth_user_id_seq:next"),
	).GetTyped(&r)
	if err != nil {
		return 0, err
	}
	return r.ID, nil
}

func (s *TarantoolStorage) InsertUser(user *UserTnt) error {
	_, err := s.conn.Do(tarantool.NewInsertRequest("auth_user").Tuple(user)).Get()
	return err
}

func (s *TarantoolStorage) GetUser(email string) (*UserTnt, error) {
	var userTuples []UserTnt
	req := tarantool.NewSelectRequest("auth_user").Index("email").Key([]interface{}{email}).Limit(1)
	err := s.conn.Do(req).GetTyped(&userTuples)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate ") {
			return nil, DuplicateError
		}
		return nil, err
	}
	for _, user := range userTuples {
		return &user, nil
	}
	return nil, NotFoundError
}
