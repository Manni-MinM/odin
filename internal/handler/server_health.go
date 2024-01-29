package handler

import (
	"time"
	"net/http"

	"github.com/Manni-MinM/odin/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ServerHealthRequest struct {
    Address     string      `json:"address"`
}

type ServerHealthHandler struct {
    repo     model.ServerHealthRepo
}

func NewHTTPServerHealthHandler(r model.ServerHealthRepo) *ServerHealthHandler {
    return &ServerHealthHandler{r}
}

func (h *ServerHealthHandler) Create(ctx echo.Context) error {
    var req ServerHealthRequest

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

    return ctx.JSON(http.StatusCreated, sh.ID)
}

func (h *ServerHealthHandler) GetAll(ctx echo.Context) error {
    serverHealthMap, err := h.repo.GetAll()
    if err != nil {
        msg := map[string]string{"message": "Internal Server Error"}
        return ctx.JSON(http.StatusInternalServerError, msg)
    }

    return ctx.JSON(http.StatusOK, serverHealthMap)
}


func (h *ServerHealthHandler) Get(ctx echo.Context) error {
    id := ctx.QueryParam("id")
    if id == "" {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request"})
    }

    sh, err := h.repo.GetByID(id)
    if err != nil {
        msg := map[string]string{"message": "Internal Server Error"}
        return ctx.JSON(http.StatusInternalServerError, msg)
    }

    return ctx.JSON(http.StatusOK, sh)
}
