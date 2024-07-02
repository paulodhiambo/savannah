package handlers

import (
	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	repo   repositories.CustomerRepositoryImpl
	logger *logrus.Logger
}

type CustomerHandlerImpl interface {
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	GetAllCustomers(c *gin.Context)
}

func NewCustomerHandler(repo repositories.CustomerRepositoryImpl, logger *logrus.Logger) *CustomerHandler {
	return &CustomerHandler{repo: repo, logger: logger}
}

// CreateCustomer @Summary Create a customer
// @Description Create a new customer
// @Tags Customers
// @Accept json
// @Produce json
// @Param name body string true "Customer name"
// @Param code body string true "Customer code"
// @Success 201 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Failure 400 {object} dto.BaseResponse
// @Router /api/v1/customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		h.logger.Warnf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Data:       nil,
			Message:    fmt.Sprintf("Failed to bind JSON: %v", err),
			StatusCode: http.StatusBadRequest,
		})
	}
	if err := h.repo.Create(&customer); err != nil {
		h.logger.Errorf("Failed to create customer: %v", err)
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Data:       nil,
			Message:    fmt.Sprintf("Failed to create customer: %v", err),
			StatusCode: http.StatusInternalServerError,
		})
	}

	h.logger.Infof("Created customer with ID: %d", customer.ID)

	c.JSON(http.StatusCreated, dto.BaseResponse{
		Data:       customer,
		Message:    fmt.Sprintf("Successfully created customer with ID: %d", customer.ID),
		StatusCode: http.StatusCreated,
	})
}

// UpdateCustomer @Summary Update a customer
// @Description Update an existing customer
// @Tags Customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param name body string true "Customer name"
// @Param code body string true "Customer code"
// @Success 200 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Failure 400 {object} dto.BaseResponse
// @Router /api/v1/customers/{id} [put]
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		h.logger.Warnf("Invalid customer ID: %s", c.Param("id"))
		c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Data:       nil,
			Message:    fmt.Sprintf("Invalid customer ID: %s", c.Param("id")),
			StatusCode: http.StatusBadRequest,
		})
	}

	customer.ID = id
	if err := h.repo.Update(&customer); err != nil {
		h.logger.Warnf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Data:       nil,
			Message:    fmt.Sprintf("Failed to bind JSON: %v", err),
			StatusCode: http.StatusInternalServerError,
		})
	}

	h.logger.Infof("Updated customer with ID: %d", customer.ID)
	c.JSON(http.StatusOK, dto.BaseResponse{
		Data:       customer,
		Message:    fmt.Sprintf("Successfully updated customer with ID: %d", customer.ID),
		StatusCode: http.StatusOK,
	})
}

// GetAllCustomers @Summary Get all customers
// @Description Get all customers
// @Tags Customers
// @Accept json
// @Produce json
// @Success 200 {array} dto.BaseResponse
// @Security ApiKeyAuth
// @Failure 400 {object} dto.BaseResponse
// @Router /api/v1/customers [get]
func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	customers, err := h.repo.GetAll()
	if err != nil {
		h.logger.Errorf("Failed to get all customers: %v", err)
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Data:       nil,
			Message:    fmt.Sprintf("Failed to get all customers: %v", err),
			StatusCode: http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, dto.BaseResponse{
		Data:       customers,
		Message:    fmt.Sprintf("Customers retried successfully"),
		StatusCode: http.StatusOK,
	})
}
