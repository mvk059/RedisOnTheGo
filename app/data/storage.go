package data

import "errors"

type StorageHelper interface {
	Set(key string, value string)
	Get(key string) (string, error)
}

type Storage struct {
	Data map[string]string
}

func NewStorage() *Storage {
	return &Storage{Data: make(map[string]string)}
}

func (s *Storage) Set(key string, value string) {
	s.Data[key] = value
}

func (s *Storage) Get(key string) (string, error) {
	value, ok := s.Data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}
