package database

import (
    "fmt"
    "context"

    "github.com/Manni-MinM/odin/internal/config"

	"github.com/go-redis/redis"
)

type RedisDB struct {
    client    *redis.Client
}

func RedisConn(conf config.RedisConfig) (Database, error) {
    addr := fmt.Sprintf("%v:%v", conf.Addr, conf.Port)
    client := redis.NewClient(&redis.Options{
        Addr: addr,
        Password: "",
        DB: 0,
    })

    ctx := context.Background()

    err := client.Ping(ctx).Err()
    if err != nil {
        return nil, err
    }

    return &RedisDB{client}, nil
}

func (rdb *RedisDB) getAllKeys() ([]string, error) {
	keys, err := rdb.client.Keys("*").Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (rdb *RedisDB) GetAllValues() ([]string, error) {
	keys, err := rdb.getAllKeys()
	if err != nil {
		return nil, err
	}

	values, err := rdb.client.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (rdb *RedisDB) Get(key string) (string, error) {
    ctx := context.Background()

    val, err := rdb.client.Get(ctx, key).Result()
    if err == redis.Nil || err != nil {
        return "", err
    }

    return val, nil
}

func (rdb *RedisDB) Set(key string, val string) (string, error) {
    ctx := context.Background()

    err := rdb.client.Set(ctx, key, val, 0).Err()
    if err == redis.Nil || err != nil {
        return "", err
    }

    return key, nil
}
