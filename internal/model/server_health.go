package model

import (
    "encoding/json"

    "github.com/Manni-MinM/odin/internal/database"
)

type ServerHealth struct {
    ID              string      `json:"id"`
    Address         string      `json:"address"`
    SuccessCount    uint64      `json:"success_count"`
    FailureCount    uint64      `json:"failure_count"`
    LastFailure     *uint64     `json:"last_failure"`
    CreatedAt       uint64      `json:"created_at"`
}

type ServerHealthRepo interface {
    Create(*ServerHealth) (error)
    GetAll() (map[string]ServerHealth, error)
    GetByID(string) (ServerHealth, error)
}

type RedisServerHealthRepo struct {
    db  *database.RedisDB
}

func NewRedisServerHealthRepo(db *database.RedisDB) *RedisServerHealthRepo {
    return &RedisServerHealthRepo{db}
}

func (r *RedisServerHealthRepo) Create(sh *ServerHealth) (error) {
    data, err := json.Marshal(sh)
	if err != nil {
		return err
	}

    id := string(sh.ID)
    jsonString := string(data)

    _, err = r.db.Set(id, jsonString)
    return err
}

func (r *RedisServerHealthRepo) GetAll() (map[string]ServerHealth, error) {
    values, err := r.db.GetAllValues()
    if err != nil {
        return map[string]ServerHealth{}, err
    }

    serverHealthMap := make(map[string]ServerHealth)
    for _, jsonString := range(values) {
        var sh ServerHealth

        err = json.Unmarshal([]byte(jsonString), &sh)
        if err != nil {
            return map[string]ServerHealth{}, err
        }

        serverHealthMap[sh.ID] = sh
    }

    return serverHealthMap, nil
}

func (r *RedisServerHealthRepo) GetByID(id string) (ServerHealth, error) {
    var sh ServerHealth

    jsonString, err := r.db.Get(id)
    if err != nil {
        return ServerHealth{}, err
    }

    err = json.Unmarshal([]byte(jsonString), &sh)
    if err != nil {
        return ServerHealth{}, err
    }

    return sh, nil
}
