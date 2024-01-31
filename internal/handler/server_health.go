package handler

import (
	metric "github.com/Manni-MinM/odin/internal/pkg/metrics"
	"net/http"
	"time"

	"github.com/Manni-MinM/odin/internal/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ServerHealthRequest struct {
	Address string `json:"address"`
}

type ServerHealthHandler struct {
	repo model.ServerHealthRepo
}

func NewHTTPServerHealthHandler(r model.ServerHealthRepo) *ServerHealthHandler {
	return &ServerHealthHandler{r}
}

func (h *ServerHealthHandler) Create(ctx echo.Context) error {
	createStart := time.Now()
	var req ServerHealthRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request"})
	}

	sh := model.ServerHealth{
		ID:           uuid.New().String(),
		Address:      req.Address,
		SuccessCount: 0,
		FailureCount: 0,
		LastFailure:  nil,
		CreatedAt:    uint64(time.Now().Unix()),
	}

	if err := h.repo.Create(&sh); err != nil {
		metric.MethodCount.WithLabelValues("Create", "failed").Inc()
		metric.MethodDuration.WithLabelValues("create_duration").Observe(float64(time.Since(createStart)))
		msg := map[string]string{"message": "Internal Server Error"}
		return ctx.JSON(http.StatusInternalServerError, msg)
	}

	metric.MethodDuration.WithLabelValues("create_duration").Observe(float64(time.Since(createStart)))
	metric.MethodCount.WithLabelValues("Create", "successful").Inc()

	return ctx.JSON(http.StatusCreated, sh.ID)
}

func (h *ServerHealthHandler) GetAll(ctx echo.Context) error {
	getAllStart := time.Now()

	serverHealthMap, err := h.repo.GetAll()

	metric.MethodDuration.WithLabelValues("get_all_duration").Observe(float64(time.Since(getAllStart)))
	if err != nil {
		metric.MethodCount.WithLabelValues("GetAll", "failed").Inc()
		msg := map[string]string{"message": "Internal Server Error"}
		return ctx.JSON(http.StatusInternalServerError, msg)
	}

	metric.MethodCount.WithLabelValues("GetAll", "successful").Inc()
	return ctx.JSON(http.StatusOK, serverHealthMap)
}

func (h *ServerHealthHandler) Get(ctx echo.Context) error {
	getStart := time.Now()
	id := ctx.QueryParam("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request"})
	}

	sh, err := h.repo.GetByID(id)
	metric.MethodDuration.WithLabelValues("get_duration").Observe(float64(time.Since(getStart)))
	if err != nil {
		metric.MethodCount.WithLabelValues("Get", "failed").Inc()
		msg := map[string]string{"message": "Internal Server Error"}
		return ctx.JSON(http.StatusInternalServerError, msg)
	}

	metric.MethodCount.WithLabelValues("Get", "successful").Inc()

	return ctx.JSON(http.StatusOK, sh)
}
