package model

import (
	"time"
    "encoding/json"

    "github.com/Manni-MinM/odin/internal/database"

	"github.com/go-redis/redis"
)

type ServerHealth struct {
    ID              uint64      `json: "id"`
    Address         string      `json: "address"`
    SuccessCount    uint64      `json: "success_count"`
    FailureCount    uint64      `json: "failure_count"`
    LastFailure     time.Time   `json: "last_failure"`
    CreatedAt       time.Time   `json: "created_at"`
}

type ServerHealthRepo interface {
    Create(*ServerHealth) error
    GetAll() ([]ServerHealth, error)
    GetByID(uint64) (ServerHealth, error)
}

type RedisServerHealthRepo struct {
    db  *database.RedisDB
}

func NewRedisServerHealthRepo(db *redis.Client) *RedisServerHealthRepo {
    return &RedisServerHealthRepo{db}
}

func (r *RedisServerHealthRepo) Create(sh *ServerHealth) error {
    data, err := json.Marshal(sh)
	if err != nil {
		return err
	}

    id := string(sh.ID)
    jsonString := string(data)

    _, err = r.db.Set(id, jsonString)
    return err
}

func (r *RedisServerHealthRepo) GetAll() ([]ServerHealth, error) {
    values, err := r.db.GetAllValues()
    if err != nil {
        return []ServerHealth{}, err
    }

    serverHealthList := []ServerHealth{}
    for _, jsonString := range(values) {
        var sh ServerHealth

        err = json.Unmarshal([]byte(jsonString), &sh)
        if err != nil {
            return []ServerHealth{}, err
        }

        serverHealthList = append(serverHealthList, sh)
    }

    return serverHealthList, nil
}

func (r *RedisServerHealthRepo) GetByID(id uint64) (ServerHealth, error) {
    var sh ServerHealth

    key := string(id)
    jsonString, err := r.db.Get(key)
    if err != nil {
        return ServerHealth{}, err
    }

    err = json.Unmarshal([]byte(jsonString), &sh)
    if err != nil {
        return ServerHealth{}, err
    }

    return sh, nil
}
