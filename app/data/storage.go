package data

import "errors"

type StorageHelper interface {
	Set(key string, value Data)
	Get(key string) (Data, error)
}

type Storage struct {
	Data map[string]Data
}

type Data struct {
	Value          string
	ExpiryTimeNano int64
}

func NewStorage() *Storage {
	return &Storage{Data: make(map[string]Data)}
}

func (s *Storage) Set(key string, value Data) {
	s.Data[key] = value
}

func (s *Storage) Get(key string) (Data, error) {
	value, ok := s.Data[key]
	if !ok {
		return Data{}, errors.New("key not found")
	}
	return value, nil
}
