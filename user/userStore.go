package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

//RedisClient redis client
var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

// Store holding a redis client instance
type Store struct {
	rds *redis.Client
}

// NewUserStore creates new redis client
func NewUserStore() Store {
	rds := redisClient
	return Store{rds: rds}
}

// FindUser finds already existing user
func (a Store) FindUser(ctx context.Context, userName string) (*User, error) {
	result, err := a.rds.HGet(ctx, userName, "resourceRates").Result()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("find: redis error: %w", err)
	}
	if result == "" {
		return nil, fmt.Errorf("find: not found")
	}

	user := &User{}
	if err := json.Unmarshal([]byte(result), user); err != nil {
		return nil, fmt.Errorf("find: unmarshal error: %w", err)
	}
	return user, nil
}

// CreateUser creates new user object in the redis server
func (a Store) CreateUser(ctx context.Context, user User) error {
	val, _ := json.Marshal(user)
	if _, err := a.rds.HSetNX(ctx, user.UserName, "resourceRates", val).Result(); err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	return nil
}

// UpdateUser updates the world object in the redis server
func (a Store) UpdateUser(ctx context.Context, user User) error {
	val, _ := json.Marshal(user)
	if _, err := a.rds.HSet(ctx, user.UserName, "resourceRates", val).Result(); err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	return nil
}
