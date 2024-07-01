COVERAGE_PROFILE := coverage.out
COVERAGE_REPORT := coverage.html

# Targets
.PHONY: test coverage clean mock document run help

# Default target
all: test coverage

# Generate Mock files
mock:
	mockgen -destination=mocks/mock_customer_repository.go -package=mocks backend/internal/repositories CustomerRepositoryImpl
	mockgen -destination=mocks/mock_order_repository.go -package=mocks backend/internal/repositories OrderRepositoryImpl

# Run tests with coverage
test:
	go test ./... -coverprofile=$(COVERAGE_PROFILE)

# Generate coverage report
coverage: test
	go tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_REPORT)
	@echo "Coverage report generated at $(COVERAGE_REPORT)"

# Clean coverage files
clean:
	rm -f $(COVERAGE_PROFILE) $(COVERAGE_REPORT)

# Generate Swagger documentation
document:
	swag init -g main.go

# Run the server
run:
	go run main.go

# Help message
help:
	@echo "Usage:"
	@echo "  make test        Run tests with coverage"
	@echo "  make coverage    Generate HTML coverage report"
	@echo "  make clean       Remove coverage files"
	@echo "  make mock        Generate mock files"
	@echo "  make document    Generate Swagger documentation"
	@echo "  make run         Run the server"
	@echo "  make help        Display this help message"
