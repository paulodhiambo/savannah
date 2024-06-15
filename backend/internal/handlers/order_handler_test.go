package handlers_test

import (
	"backend/internal/dto"
	"backend/internal/handlers"
	"backend/internal/models"
	"backend/mocks"
	"backend/pkg/logging"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrderHandler_CreateOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
	logger := logging.GetLogger()
	handler := handlers.NewOrderHandler(mockRepo, logger)

	router := gin.Default()
	router.POST("/api/v1/orders", handler.CreateOrder)

	order := dto.CreateOrderRequest{ProductID: 1, Quantity: 2, UserId: 1}
	mockRepo.EXPECT().Create(gomock.Any()).Return(nil)

	reqBody, _ := json.Marshal(order)
	req, _ := http.NewRequest("POST", "/api/v1/orders", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response dto.BaseResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Order created successfully", response.Message)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, float64(order.ProductID), response.Data.(map[string]interface{})["product_id"])
}

func TestOrderHandler_UpdateOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
	logger := logging.GetLogger()
	handler := handlers.NewOrderHandler(mockRepo, logger)

	router := gin.Default()
	router.PUT("/api/v1/orders/:id", handler.UpdateOrder)

	order := models.Order{ProductID: 1, Quantity: 2, UserId: 1, Total: 20.0}
	mockRepo.EXPECT().Update(gomock.Any()).Return(nil)

	reqBody, _ := json.Marshal(order)
	req, _ := http.NewRequest("PUT", "/api/v1/orders/1", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dto.BaseResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Order updated successfully", response.Message)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(order.ProductID), response.Data.(map[string]interface{})["product_id"])
}

func TestOrderHandler_DeleteOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
	logger := logging.GetLogger()
	handler := handlers.NewOrderHandler(mockRepo, logger)

	router := gin.Default()
	router.DELETE("/api/v1/orders/:id", handler.DeleteOrder)

	mockRepo.EXPECT().Delete(1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/api/v1/orders/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	var response dto.BaseResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestOrderHandler_GetOrderByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
	logger := logging.GetLogger()
	handler := handlers.NewOrderHandler(mockRepo, logger)

	router := gin.Default()
	router.GET("/api/v1/orders/:id", handler.GetOrderByID)

	order := models.Order{ProductID: 1, Quantity: 2, UserId: 1, Total: 20.0}
	mockRepo.EXPECT().GetByID(1).Return(&order, nil)

	req, _ := http.NewRequest("GET", "/api/v1/orders/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dto.BaseResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Order fetched successfully", response.Message)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(order.ProductID), response.Data.(map[string]interface{})["product_id"])
}

func TestOrderHandler_GetAllOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
	logger := logging.GetLogger()
	handler := handlers.NewOrderHandler(mockRepo, logger)

	router := gin.Default()
	router.GET("/api/v1/orders", handler.GetAllOrders)

	orders := []models.Order{
		{ProductID: 1, Quantity: 2, UserId: 1, Total: 20.0},
		{ProductID: 2, Quantity: 1, UserId: 1, Total: 10.0},
	}
	mockRepo.EXPECT().GetAll().Return(orders, nil)

	req, _ := http.NewRequest("GET", "/api/v1/orders", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dto.BaseResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Orders fetched successfully", response.Message)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(orders[0].ProductID), response.Data.([]interface{})[0].(map[string]interface{})["product_id"])
}

func TestOrderHandler_GetOrdersByUserID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
	logger := logging.GetLogger()
	handler := handlers.NewOrderHandler(mockRepo, logger)

	router := gin.Default()
	router.GET("/api/v1/users/:user_id/orders", handler.GetOrdersByUserID)

	orders := []models.Order{
		{ProductID: 1, Quantity: 2, UserId: 1, Total: 20.0},
		{ProductID: 2, Quantity: 1, UserId: 1, Total: 10.0},
	}
	mockRepo.EXPECT().GetOrdersByUserID(1).Return(orders, nil)

	req, _ := http.NewRequest("GET", "/api/v1/users/1/orders", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dto.BaseResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Orders fetched successfully", response.Message)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(orders[0].ProductID), response.Data.([]interface{})[0].(map[string]interface{})["product_id"])
}
