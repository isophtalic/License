package redis

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/isophtalic/License/internal/configs"

	"gopkg.in/redis.v5"
)

type PubSub *redis.PubSub

type RedisProvider struct {
	url      string
	server   string
	password string
	db       int
	redis    *RedisClient
}

type RedisClient struct {
	client *redis.Client
}

type DatabaseExecutionError struct {
	Message string
}

func (e DatabaseExecutionError) Error() string {
	return e.Message
}

func NewRedisProviderFromURL(config *configs.Configure) *RedisProvider {
	client := newRedisClientFromURL(config.RedisPort)
	if client == nil {
		log.Fatalln("Redis server connected unsuccessfully")
	}
	return &RedisProvider{
		url:   config.RedisPort,
		redis: client,
	}
}
func newRedisClientFromURL(u string) *RedisClient {
	tempU, err := url.Parse(u)
	if err != nil {
		log.Fatalln("Redis server 1 connected unsuccessfully")
	}
	result := new(RedisClient)

	switch tempU.Scheme {
	case "redis-sentinel":
		options, err := parseURLSentinel(u)
		if err != nil {
			log.Fatalln("Redis server 2 connected unsuccessfully")
		}

		result.client = redis.NewFailoverClient(options)

	default:
		options, err := redis.ParseURL(u)
		if err != nil {
			log.Fatalln("Redis server 3 connected unsuccessfully")
		}

		result.client = redis.NewClient(options)
	}

	return result
}

func parseURLSentinel(redisURL string) (*redis.FailoverOptions, error) {
	result := new(redis.FailoverOptions)

	o := &redis.Options{Network: "tcp"}
	u, err := url.Parse(redisURL)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "redis-sentinel" {
		return nil, errors.New("invalid redis sentinel URL scheme: " + u.Scheme)
	}

	if u.User == nil {
		return nil, errors.New("invalid redis sentinel URL User: " + u.Scheme)
	}

	if masterName := u.User.Username(); masterName == "" {
		return nil, errors.New("invalid redis sentinel Master Name: " + u.Scheme)
	}

	result.MasterName = u.User.Username()

	if p, ok := u.User.Password(); ok {
		result.Password = p
	}

	if len(u.Query()) > 0 {
		return nil, errors.New("no options supported")
	}

	sentinelHostPorts := strings.Split(u.Host, ",")

	var sentinelAddresses []string

	for _, v := range sentinelHostPorts {
		h, p, err := net.SplitHostPort(v)
		if err != nil {
			return nil, errors.New("error sentinel address")
		}

		sentinelAddresses = append(sentinelAddresses, net.JoinHostPort(h, p))
	}
	result.SentinelAddrs = sentinelAddresses

	f := strings.FieldsFunc(u.Path, func(r rune) bool {
		return r == '/'
	})
	switch len(f) {
	case 0:
		o.DB = 0
	case 1:
		if o.DB, err = strconv.Atoi(f[0]); err != nil {
			return nil, fmt.Errorf("invalid redis database number: %q", f[0])
		}
	default:
		return nil, errors.New("invalid redis URL path: " + u.Path)
	}

	result.ReadTimeout = 5 * time.Minute
	result.WriteTimeout = 5 * time.Minute

	return result, nil
}

func (provider *RedisProvider) RedisClient() *RedisClient {
	return provider.redis
}

func (provider *RedisProvider) NewRedisClient() *RedisClient {
	if provider.url != "" {
		return newRedisClientFromURL(provider.url)
	}
	return newRedisClient(provider.server, provider.password, provider.db)
}

func newRedisClient(server string, password string, db int) *RedisClient {
	result := new(RedisClient)
	result.client = redis.NewClient(&redis.Options{
		Addr:         server,
		Password:     password,
		DB:           db,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	})

	return result
}

func (provider *RedisProvider) NewError(e error) error {
	if e == nil {
		return nil
	}
	return DatabaseExecutionError{Message: fmt.Sprintf("Redis execution error: %s", e.Error())}
}

func (r *RedisClient) HSet(key, field string, value interface{}) error {
	return r.client.HSet(key, field, value).Err()
}

func (r *RedisClient) HDel(key string, fields ...string) error {
	return r.client.HDel(key, fields...).Err()
}

func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	cmd := r.client.HGetAll(key)
	if err := cmd.Err(); err != nil {
		return map[string]string{}, err
	}

	return cmd.Val(), nil
}

func (r *RedisClient) HGet(key, field string) (string, error) {
	cmd := r.client.HGet(key, field)
	if err := cmd.Err(); err != nil {
		return "", err
	}

	return cmd.Val(), nil
}

func (r *RedisClient) HExists(key, field string) (bool, error) {
	cmd := r.client.HExists(key, field)
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val(), nil
}

func (r *RedisClient) Get(key string) (string, error) {
	cmd := r.client.Get(key)
	if err := cmd.Err(); err != nil {
		return "", err
	}
	return cmd.Val(), nil
}

func (r *RedisClient) Set(key string, value string, exp time.Duration) error {
	return r.client.Set(key, value, exp).Err()

}

func (r *RedisClient) ExpiresAt(key string, exp time.Time) error {
	return r.client.ExpireAt(key, exp).Err()
}

func (r *RedisClient) Increment(key string) error {
	return r.client.Incr(key).Err()
}

func (r *RedisClient) DelAllByFields(accessID, field string) error {
	cursor := uint64(0)
	keys, cursor, err := r.client.HScan(accessID, cursor, field, 0).Result()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(keys); i += 2 {
		_, err := r.client.HDel(accessID, keys[i]).Result()
		if err != nil {
			return err
		}
	}
	return nil
}
