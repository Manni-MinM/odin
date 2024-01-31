package database

import (
	"fmt"
	metric "github.com/Manni-MinM/odin/internal/pkg/metrics"
	"time"

	"github.com/Manni-MinM/odin/internal/config"

	"github.com/go-redis/redis"
)

type RedisDB struct {
	client *redis.Client
}

func RedisConn(conf config.Redis) (*RedisDB, error) {
	addr := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password,
		DB:       conf.DBName,
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
	getAllValuesTime := time.Now()
	keys, err := rdb.getAllKeys()
	if err != nil {
		metric.MethodCount.WithLabelValues("redis_get_all_values", "failed").Inc()
		return nil, err
	}

	if len(keys) == 0 {
		return []string{}, nil
	}

	values, err := rdb.client.MGet(keys...).Result()
	if err != nil {
		metric.MethodCount.WithLabelValues("redis_get_all_values", "failed").Inc()
		return nil, err
	}

	valueList := []string{}
	for _, val := range values {
		valueList = append(valueList, val.(string))
	}

	metric.MethodDuration.WithLabelValues("redis_get_all_values_duration").Observe(float64(time.Since(getAllValuesTime)))
	metric.MethodCount.WithLabelValues("redis_get_all_values", "successful").Inc()

	return valueList, nil
}

func (rdb *RedisDB) Get(key string) (string, error) {
	getTime := time.Now()

	val, err := rdb.client.Get(key).Result()
	metric.MethodDuration.WithLabelValues("redis_get_duration").Observe(float64(time.Since(getTime)))
	if err == redis.Nil || err != nil {
		metric.MethodCount.WithLabelValues("redis_get", "failed").Inc()
		return "", err
	}

	metric.MethodCount.WithLabelValues("redis_get", "successful").Inc()

	return val, nil
}

func (rdb *RedisDB) Set(key string, val string) (string, error) {
	setTime := time.Now()
	err := rdb.client.Set(key, val, 0).Err()
	metric.MethodDuration.WithLabelValues("redis_set_duration").Observe(float64(time.Since(setTime)))
	if err == redis.Nil || err != nil {
		metric.MethodCount.WithLabelValues("redis_set", "failed").Inc()
		return "", err
	}

	metric.MethodCount.WithLabelValues("redis_set", "successful").Inc()

	return key, nil
}
