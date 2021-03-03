package rstorage

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

//UserCred struct
type UserCred struct {
	UserName     string `json:"userName"`
	UserPassword string `json:"userPassword"`
}

func (t *UserCred) marshalBinary() []byte {
	val, _ := json.Marshal(t)
	return val
}

func (t *UserCred) unmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	return nil
}

// Auth structure
type Auth struct {
	rds *redis.Client
}

// NewAuth creates new auth
func NewAuth() Auth {
	rds := redisClient
	return Auth{rds: rds}
}

// FindUser finds already existing user
func (a Auth) FindUser(ctx context.Context, userName string) (*UserCred, error) {
	print("test inside Find User")
	print(userName)
	result, err := a.rds.HGet(ctx, userName, "creds").Result()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("find: redis error: %w", err)
	}
	if result == "" {
		return nil, fmt.Errorf("find: not found")
	}

	token := &UserCred{}
	if err := token.unmarshalBinary([]byte(result)); err != nil {
		return nil, fmt.Errorf("find: unmarshal error: %w", err)
	}
	return token, nil
}

// Create creats new user creds in the redis server
func (a Auth) Create(ctx context.Context, userCred UserCred) error {
	val, _ := json.Marshal(userCred)
	if _, err := a.rds.HSetNX(ctx, userCred.UserName, "creds", val).Result(); err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	// a.rds.Expire(ctx, userCred.UserName, time.Minute)
	return nil
}

// CreatUserCred creats user
func CreatUserCred(userName string, userPassword string) UserCred {
	return UserCred{
		UserName:     userName,
		UserPassword: userPassword,
	}

}
