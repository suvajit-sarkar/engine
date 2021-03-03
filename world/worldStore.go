package world

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

// NewWorldStore creates new auth
func NewWorldStore() Store {
	rds := redisClient
	return Store{rds: rds}
}

// FindWorld finds already existing user
func (a Store) FindWorld(ctx context.Context, worldName string) (*World, error) {
	print("test inside Find User")
	print(worldName)
	result, err := a.rds.HGet(ctx, worldName, "creds").Result()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("find: redis error: %w", err)
	}
	if result == "" {
		return nil, fmt.Errorf("find: not found")
	}

	// token := &World{}
	// if err := json.Unmarshal([]byte(result)); err != nil {
	// 	return nil, fmt.Errorf("find: unmarshal error: %w", err)
	// }
	return nil, nil
}

// GetNode finds node info
func (a Store) GetNode(ctx context.Context, nodeCoordinate string) (*Node, error) {
	result, err := a.rds.HGet(ctx, nodeCoordinate, "creds").Result()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("find: redis error: %w", err)
	}
	if result == "" {
		return nil, fmt.Errorf("find: not found")
	}
	node := &Node{}
	if err := json.Unmarshal([]byte(result), node); err != nil {
		return nil, fmt.Errorf("find: unmarshal error: %w", err)
	}
	return node, nil
}

// GetVaccantCityList finds node info
func (a Store) GetVaccantCityList(ctx context.Context) ([]string, error) {
	result, err := a.rds.HGet(ctx, "NodeLists", "VaccantCityList").Result()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("find: redis error: %w", err)
	}
	if result == "" {
		return nil, fmt.Errorf("find: not found")
	}
	var list []string
	if err := json.Unmarshal([]byte(result), &list); err != nil {
		return nil, fmt.Errorf("find: unmarshal error: %w", err)
	}
	return list, nil
}

// CreateWorld creates new world object in the redis server
func (a Store) CreateWorld(ctx context.Context, obj World) error {
	val, _ := json.Marshal(obj)
	if _, err := a.rds.HSetNX(ctx, obj.WorldName, "world", val).Result(); err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	return nil
}

// CreateNode creates new world object in the redis server
func (a Store) CreateNode(ctx context.Context, obj *Node) error {
	val, _ := json.Marshal(obj)
	if _, err := a.rds.HSetNX(ctx, "coordinate", obj.Coordinate, val).Result(); err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	return nil
}

// CreateVaccantCityList creates new world object in the redis server
func (a Store) CreateVaccantCityList(ctx context.Context, obj []string) error {
	val, _ := json.Marshal(obj)
	if _, err := a.rds.HSetNX(ctx, "NodeLists", "VaccantCityList", val).Result(); err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	return nil
}

// UpdateWorld updates the world object in the redis server
func (a Store) UpdateWorld(ctx context.Context, worldObj *World) error {
	val, _ := json.Marshal(worldObj)
	if _, err := a.rds.HSet(ctx, worldObj.WorldName, "world", val).Result(); err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	return nil
}

// UpdateNode updates new world object in the redis server
func (a Store) UpdateNode(ctx context.Context, obj Node) error {
	val, _ := json.Marshal(obj)
	if _, err := a.rds.HSet(ctx, "coordinate", obj.Coordinate, val).Result(); err != nil {
		return fmt.Errorf("create: redis error: %w", err)
	}
	return nil
}
