package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"l0/internal/errs"
	"l0/internal/services"
	"l0/pkg/responses"
	"net/http"
)

type orderHandler struct {
	service services.Orders
}

func InitOrderHandler(
	service services.Orders,
) Orders {
	return &orderHandler{
		service: service,
	}
}

// GetByID
// @Summary Get Order
// @Description Get order by ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Success 200 {object} dto.Order "Order data"
// @Failure 404 {object} responses.MessageResponse "Order with that ID not found"
// @Failure 500 "Internal server error"
// @Router /api/order/{order_id} [get]
func (o orderHandler) GetByID(c *gin.Context) {
	idStr := c.Param("order_id")

	order, err := o.service.GetByID(idStr)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrNoOrder):
			c.JSON(http.StatusNotFound, responses.MessageResponse{Message: err.Error()})
		default:
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, order)
}
