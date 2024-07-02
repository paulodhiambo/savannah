package handlers

import (
	"backend/internal/config"
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	repo   repositories.OrderRepositoryImpl
	logger *logrus.Logger
}

func NewOrderHandler(repo repositories.OrderRepositoryImpl, logger *logrus.Logger) *OrderHandler {
	return &OrderHandler{repo: repo, logger: logger}
}

// CreateOrder @Summary Create a new order
// @Description Create a new order
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body dto.CreateOrderRequest true "Order"
// @Success 201 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Failure 400 {object} dto.BaseResponse
// @Failure 500 {object} dto.BaseResponse
// @Router /api/v1/orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var createOrder dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&createOrder); err != nil {
		h.logger.Warnf("invalid order data: %v", err)
		c.JSON(http.StatusBadRequest, dto.BaseResponse{Message: "Invalid order data", StatusCode: http.StatusBadRequest})
		return
	}
	order := models.Order{
		ProductID: createOrder.ProductID,
		Quantity:  createOrder.Quantity,
		UserId:    createOrder.UserId,
	}

	if err := h.repo.Create(&order); err != nil {
		h.logger.Warnf("failed to create order: %v", err)
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{Message: "Failed to create order", StatusCode: http.StatusInternalServerError})
		return
	}
	err := utils.SendSMS(config.AppConfig.SMSSandboxAPIKey, config.AppConfig.SMSSandboxUserName, "+254722123123", "Order created successfully")
	if err != nil {
		h.logger.Warnf("failed to create order: %v", err)
	}
	c.JSON(http.StatusCreated, dto.BaseResponse{Data: order, Message: "Order created successfully", StatusCode: http.StatusCreated})
}

// UpdateOrder @Summary Update an existing order
// @Description Update an existing order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body dto.CreateOrderRequest true "Order"
// @Success 200 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Failure 400 {object} dto.BaseResponse
// @Failure 500 {object} dto.BaseResponse
// @Router /api/v1/orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	var createOrder dto.CreateOrderRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Warnf("invalid order ID: %v", err)
		c.JSON(http.StatusBadRequest, dto.BaseResponse{Message: "Invalid order ID", StatusCode: http.StatusBadRequest})
		return
	}

	if err := c.ShouldBindJSON(&createOrder); err != nil {
		h.logger.Warnf("invalid order data: %v", err)
		c.JSON(http.StatusBadRequest, dto.BaseResponse{Message: "Invalid order data", StatusCode: http.StatusBadRequest})
		return
	}

	order := models.Order{
		ProductID: createOrder.ProductID,
		Quantity:  createOrder.Quantity,
		UserId:    createOrder.UserId,
	}

	order.ID = uint(id)
	if err := h.repo.Update(&order); err != nil {
		h.logger.Warnf("failed to update order: %v", err)
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{Message: "Failed to update order", StatusCode: http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse{Data: order, Message: "Order updated successfully", StatusCode: http.StatusOK})
}

// DeleteOrder @Summary Delete an order
// @Description Delete an existing order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 204 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Failure 400 {object} dto.BaseResponse
// @Failure 500 {object} dto.BaseResponse
// @Router /api/v1/orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Warnf("invalid order ID: %v", err)
		c.JSON(http.StatusBadRequest, dto.BaseResponse{Message: "Invalid order ID", StatusCode: http.StatusBadRequest})
		return
	}

	if err := h.repo.Delete(id); err != nil {
		h.logger.Warnf("failed to delete order: %v", err)
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{Message: "Failed to delete order", StatusCode: http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusNoContent, dto.BaseResponse{Message: "Order deleted successfully", StatusCode: http.StatusNoContent})
}

// GetOrderByID @Summary Get an order by ID
// @Description Get an order by ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Failure 400 {object} dto.BaseResponse
// @Failure 500 {object} dto.BaseResponse
// @Router /api/v1/orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Warnf("invalid order ID: %v", err)
		c.JSON(http.StatusBadRequest, dto.BaseResponse{Message: "Invalid order ID", StatusCode: http.StatusBadRequest})
		return
	}

	order, err := h.repo.GetByID(id)
	if err != nil {
		h.logger.Warnf("failed to get order by ID: %v", err)
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{Message: "Failed to get order", StatusCode: http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse{Data: *order, Message: "Order fetched successfully", StatusCode: http.StatusOK})
}

// GetAllOrders @Summary Get all orders
// @Description Get all orders
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Failure 500 {object} dto.BaseResponse
// @Router /api/v1/orders [get]
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.repo.GetAll()
	if err != nil {
		h.logger.Warnf("failed to get all orders: %v", err)
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{Message: "Failed to get orders", StatusCode: http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse{Data: orders, Message: "Orders fetched successfully", StatusCode: http.StatusOK})
}

// GetOrdersByUserID
// @Summary Get orders by user ID
// @Description Get orders by user ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Failure 400 {object} dto.BaseResponse
// @Failure 500 {object} dto.BaseResponse
// @Router /api/v1/users/{user_id}/orders [get]
func (h *OrderHandler) GetOrdersByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		h.logger.Warnf("invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, dto.BaseResponse{Message: "Invalid user ID", StatusCode: http.StatusBadRequest})
		return
	}

	orders, err := h.repo.GetOrdersByUserID(userID)
	if err != nil {
		h.logger.Warnf("failed to get orders by user ID: %v", err)
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{Message: "Failed to get orders", StatusCode: http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse{Data: orders, Message: "Orders fetched successfully", StatusCode: http.StatusOK})
}
