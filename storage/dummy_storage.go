package storage

type DummyTarantoolStorage struct {
}

func (ds *DummyTarantoolStorage) GetAuthUserIDSeqNext() (uint64, error) {
	return 0, nil
}
func (ds *DummyTarantoolStorage) InsertUser(user *UserTnt) error {
	return nil
}
func (ds *DummyTarantoolStorage) GetUser(email string) (*UserTnt, error) {
	return nil, nil
}
func (ds *DummyTarantoolStorage) InsertLink(tntLink *LinkTnt) error {
	return nil
}
func (ds *DummyTarantoolStorage) GetLinkSeqIDNext() (uint64, error) {
	return 0, nil
}
func (ds *DummyTarantoolStorage) GetByURL(userID uint64, originalURL string) (*LinkTnt, error) {
	return nil, nil
}