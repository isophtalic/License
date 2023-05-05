package redis

import (
	"fmt"
	"time"

	"git.cyradar.com/license-manager/backend/internal/configs"
)

var UserRedisAccessIdPrefix = "user_access"

type UserRedisRepository struct {
	redisClient *RedisClient
}

func NewUserAccessIDRedisRepository(config *configs.Configure) *UserRedisRepository {
	return &UserRedisRepository{
		redisClient: NewRedisProviderFromURL(config).RedisClient(),
	}
}

func (repo *UserRedisRepository) redisKey(username string) string {
	return fmt.Sprintf("%s:%s", UserRedisAccessIdPrefix, username)
}

func (repo *UserRedisRepository) Set(username string, accessID interface{}, exp time.Time) error {
	err := repo.redisClient.HSet(UserRedisAccessIdPrefix, username, accessID)
	if err != nil {
		return err
	}
	err = repo.redisClient.ExpiresAt(UserRedisAccessIdPrefix, exp)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRedisRepository) Exists(username, accessID string) (bool, error) {
	return repo.redisClient.HExists(repo.redisKey(username), accessID)
}

func (repo *UserRedisRepository) Get(key string) (result string, err error) {
	val, err := repo.redisClient.HGet(UserRedisAccessIdPrefix, key)
	if err != nil {
		return
	}
	return val, nil
}

func (repo *UserRedisRepository) All(username string) (map[string]string, error) {
	return repo.redisClient.HGetAll(UserRedisAccessIdPrefix)
}

func (repo *UserRedisRepository) Delete(key string) error {
	err := repo.redisClient.HDel(UserRedisAccessIdPrefix, key)
	return err
}

func (repo *UserRedisRepository) Increment(key string) error {
	return repo.redisClient.Increment(key)
}

func (repo *UserRedisRepository) DeleteAll(field string) error {
	return repo.redisClient.DelAllByFields(UserRedisAccessIdPrefix, field)
}
