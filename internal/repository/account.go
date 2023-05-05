package repository

import "time"

type AccountRepository interface {
	Set(username string, accessID interface{}, exp time.Time) error
	Exists(username, id string) (bool, error)
	Get(key string) (string, error)
	All(username string) (map[string]string, error)
	Delete(key string) error
	DeleteAll(field string) error
	Increment(key string) error
}
