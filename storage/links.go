package storage

import (
	"github.com/tarantool/go-tarantool/v2"
)

func (s *TarantoolStorage) InsertLink(tntLink *LinkTnt) error {
	_, err := s.conn.Do(tarantool.NewInsertRequest("links").Tuple(tntLink)).Get()
	return err
}

func (s *TarantoolStorage) GetLinkSeqIDNext() (uint64, error) {
	r := SeqNextResult{}
	err := s.conn.Do(
		tarantool.NewCallRequest("box.sequence.links_id_seq:next"),
	).GetTyped(&r)
	if err != nil {
		return 0, err
	}
	return r.ID, nil
}

func (s *TarantoolStorage) GetByShortURL(shortURL string) (*LinkTnt, error) {
	var linkTuples []LinkTnt
	req := tarantool.NewSelectRequest("links").Index("short_url_idx").Key([]interface{}{shortURL})
	err := s.conn.Do(req).GetTyped(&linkTuples)
	if err != nil {
		return nil, err
	}
	for _, link := range linkTuples {
		return &link, nil
	}
	return nil, NotFoundError
}
