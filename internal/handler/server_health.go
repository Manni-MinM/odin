package handler

import (
	"time"
	"net/http"
    "encoding/json"

	"github.com/Manni-MinM/odin/internal/model"
	"github.com/Manni-MinM/odin/internal/request"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ServerHealthHandler struct {
    repo     model.ServerHealthRepo
}

func NewHTTPServerHealthHandler(r model.ServerHealthRepo) *ServerHealthHandler {
    return &ServerHealthHandler{r}
}

func (h *ServerHealthHandler) Create(ctx echo.Context) error {
    var req request.ServerHealth

    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request"})
    }

    sh := model.ServerHealth {
        ID: uuid.New().String(),
        Address: req.Address,
        SuccessCount: 0,
        FailureCount: 0,
        LastFailure: nil,
        CreatedAt: uint64(time.Now().Unix()),
    }

    if err := h.repo.Create(&sh); err != nil {
        msg := map[string]string{"message": "Internal Server Error"}
        return ctx.JSON(http.StatusInternalServerError, msg)
    }

    resp, err := json.Marshal(sh)
	if err != nil {
		return err
	}

    return ctx.JSON(http.StatusCreated, resp)
}

func (h *ServerHealthHandler) GetAll(ctx echo.Context) error {
    serverHealthMap, err := h.repo.GetAll()
    if err != nil {
        msg := map[string]string{"message": err.Error()}
        // msg := map[string]string{"message": "Internal Server Error"}
        return ctx.JSON(http.StatusInternalServerError, msg)
    }

    resp, err := json.Marshal(serverHealthMap)
	if err != nil {
		return err
	}

    return ctx.JSON(http.StatusOK, resp)
}


func (h *ServerHealthHandler) Get(ctx echo.Context) error {
    id := ctx.Param("id")
    if id == "" {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request"})
    }

    sh, err := h.repo.GetByID(id)
    if err != nil {
        msg := map[string]string{"message": "Internal Server Error"}
        return ctx.JSON(http.StatusInternalServerError, msg)
    }

    resp, err := json.Marshal(sh)
	if err != nil {
		return err
	}

    return ctx.JSON(http.StatusOK, resp)
}
