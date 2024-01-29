package database

import (
    "fmt"

    "github.com/Manni-MinM/odin/internal/config"

	"github.com/go-redis/redis"
)

type RedisDB struct {
    client    *redis.Client
}

func RedisConn(conf config.Redis) (*RedisDB, error) {
    addr := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
    client := redis.NewClient(&redis.Options{
        Addr: addr,
        Password: "",
        DB: 0,
    })

    err := client.Ping().Err()
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

    if len(keys) == 0 {
        return []string{}, nil
    }

	values, err := rdb.client.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

    valueList := []string{}
    for _, val := range(values) {
        valueList = append(valueList, val.(string))
	}

	return valueList, nil
}

func (rdb *RedisDB) Get(key string) (string, error) {
    val, err := rdb.client.Get(key).Result()
    if err == redis.Nil || err != nil {
        return "", err
    }

    return val, nil
}

func (rdb *RedisDB) Set(key string, val string) (string, error) {
    err := rdb.client.Set(key, val, 0).Err()
    if err == redis.Nil || err != nil {
        return "", err
    }

    return key, nil
}
