package handlers

import (
	"backend/internal/dto"
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

func TestCreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logging.GetLogger()

	mockRepo := mocks.NewMockCustomerRepositoryImpl(ctrl)
	handler := NewCustomerHandler(mockRepo, logger)

	customer := &models.Customer{Name: "Test Customer", Code: "TST123"}
	customerJSON, _ := json.Marshal(customer)

	mockRepo.EXPECT().Create(gomock.Any()).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/v1/customers", handler.CreateCustomer)

	req, err := http.NewRequest("POST", "/api/v1/customers", bytes.NewBuffer(customerJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestUpdateCustomer(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logging.GetLogger()

	mockRepo := mocks.NewMockCustomerRepositoryImpl(ctrl)
	handler := NewCustomerHandler(mockRepo, logger)

	router := gin.New()
	router.PUT("/api/v1/customers/:id", handler.UpdateCustomer)

	// Test case: Successful update
	customer := &models.Customer{ID: 1, Name: "John Doe", Code: "C123"}
	mockRepo.EXPECT().Update(gomock.Any()).Return(nil)

	reqBody, _ := json.Marshal(customer)
	req, _ := http.NewRequest("PUT", "/api/v1/customers/1", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dto.BaseResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)
}

func TestGetAllCustomers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logging.GetLogger()

	mockRepo := mocks.NewMockCustomerRepositoryImpl(ctrl)
	handler := NewCustomerHandler(mockRepo, logger)

	router := gin.New()
	router.GET("/api/v1/customers", handler.GetAllCustomers)

	// Test case: Successful fetch
	customers := []models.Customer{
		{ID: 1, Name: "John Doe", Code: "C123"},
		{ID: 2, Name: "Jane Doe", Code: "C124"},
	}
	mockRepo.EXPECT().GetAll().Return(customers, nil)

	req, _ := http.NewRequest("GET", "/api/v1/customers", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dto.BaseResponse
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, response.StatusCode)
}
