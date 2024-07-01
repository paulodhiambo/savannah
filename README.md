# Go Application

This project consists of a simple Go service backend that manages customers and orders, with authentication via OpenID Connect, and sends SMS alerts when an order is placed. This project also includes unit tests, CI/CD setup, and deployment instructions.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Authentication and Authorization](#authentication-and-authorization)
- [SMS Notifications](#sms-notifications)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [License](#license)

## Features

- Manage customers and orders
- REST API for creating and retrieving customers and orders
- Authentication and authorization using OpenID Connect
- SMS notifications using Africa’s Talking SMS gateway
- Unit tests with coverage checking
- CI/CD pipeline setup
- Deployment instructions

## Prerequisites

- Go 1.22+
- Docker (optional, for database setup)
- Africa’s Talking account for SMS service
- OpenID Connect provider for authentication

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/paulodhiambo/savannah.git
    cd savannah
    ```

2. Install Go dependencies:
    ```bash
    go mod tidy
    ```

## Configuration

Create a `.env` file in the `backend` directory with the following content:
```env
DATABASE_URL=postgres://user:password@localhost:5432/go_app
AFRICAS_TALKING_API_KEY=your_africas_talking_api_key
AFRICAS_TALKING_USERNAME=your_africas_talking_username
OIDC_CLIENT_ID=your_oidc_client_id
OIDC_CLIENT_SECRET=your_oidc_client_secret
OIDC_ISSUER_URL=your_oidc_issuer_url
```

## Running the Application

1. Start the backend server:
    ```bash
    cd backend
    make run
    ```

## API Endpoints

### Customers
- **GET** `/customers` - Retrieve all customers
- **POST** `/customers` - Create a new customer
   - Request body:
       ```json
       {
           "name": "John Doe",
           "code": "CUST001"
       }
       ```

### Orders
- **GET** `/orders` - Retrieve all orders
- **POST** `/orders` - Create a new order
   - Request body:
       ```json
       {
           "customer_id": 1,
           "item": "Item Name",
           "amount": 100.0,
           "time": "2024-06-14T12:00:00Z"
       }
       ```

## Authentication and Authorization

The application uses OpenID Connect for authentication and authorization. Ensure you have configured the OIDC provider details in the `.env` file.

## SMS Notifications

The application sends SMS notifications using Africa’s Talking SMS gateway. Ensure you have configured your Africa’s Talking API key and username in the `.env` file.

## Testing

Run backend tests:
```bash
cd backend
make test
```

## Project Structure

```
go-app/
│
├── backend/
│   ├── main.go
│   ├── models.go
│   ├── handlers.go
│   ├── routes.go
│   ├── Makefile
│   ├── .env
│   └── ...
│
├── .github/
│   ├── workflows/
│   │   ├── ci.yml
│   │   └── cd.yml
│
└── README.md
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

- Deployed URL: [https://savannah-api-dot-streempoint.ue.r.appspot.com/](https://savannah-api-dot-streempoint.ue.r.appspot.com/)
- Docs URL: [https://savannah-api-dot-streempoint.ue.r.appspot.com/](https://savannah-api-dot-streempoint.ue.r.appspot.com/docs)