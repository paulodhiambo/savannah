# Go-Nuxt Application

This project consists of a simple Go service backend and a Nuxt frontend. It manages customers and orders, with authentication via OpenID Connect, and sends SMS alerts when an order is placed. This project also includes unit tests, CI/CD setup, and deployment instructions.

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
- [CI/CD Setup](#cicd-setup)
- [Deployment](#deployment)
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

- Go 1.16+
- Node.js 14+
- Docker (optional, for database setup)
- Africa’s Talking account for SMS service
- OpenID Connect provider for authentication

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/go-nuxt-app.git
    cd go-nuxt-app
    ```

2. Install Go dependencies:
    ```bash
    go mod tidy
    ```

3. Install Node.js dependencies for the Nuxt frontend:
    ```bash
    cd frontend
    npm install
    ```

## Configuration

1. **Backend Configuration:**

   Create a `.env` file in the root directory with the following content:
    ```env
    DATABASE_URL=postgres://user:password@localhost:5432/go_nuxt_app
    AFRICAS_TALKING_API_KEY=your_africas_talking_api_key
    AFRICAS_TALKING_USERNAME=your_africas_talking_username
    OIDC_CLIENT_ID=your_oidc_client_id
    OIDC_CLIENT_SECRET=your_oidc_client_secret
    OIDC_ISSUER_URL=your_oidc_issuer_url
    ```

2. **Frontend Configuration:**

   Update the `nuxt.config.js` file in the `frontend` directory with your configuration details.

## Running the Application

1. Start the backend server:
    ```bash
    go run main.go
    ```

2. Start the Nuxt frontend:
    ```bash
    cd frontend
    npm run dev
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

The application sends SMS notifications using the Africa’s Talking SMS gateway. Ensure you have configured your Africa’s Talking API key and username in the `.env` file.

## Testing

1. Run backend tests:
    ```bash
    go test ./... -cover
    ```

2. Run frontend tests:
    ```bash
    cd frontend
    npm run test
    ```

## CI/CD Setup

1. **CI Setup:**

   Add the following `.github/workflows/ci.yml` for GitHub Actions:
    ```yaml
    name: CI

    on:
      push:
        branches: [ main ]
      pull_request:
        branches: [ main ]

    jobs:
      build:

        runs-on: ubuntu-latest

        steps:
        - uses: actions/checkout@v2
        - name: Set up Go
          uses: actions/setup-go@v2
          with:
            go-version: 1.16
        - name: Install dependencies
          run: go mod tidy
        - name: Run tests
          run: go test ./... -cover
    ```

2. **CD Setup:**

   For CD, you can use platforms like Heroku, AWS, or any other PAAS/FAAS/IAAS provider. Example for Heroku:

    ```yaml
    name: CD to Heroku

    on:
      push:
        branches:
          - main

    jobs:
      deploy:
        runs-on: ubuntu-latest

        steps:
        - uses: actions/checkout@v2
        - name: Install Heroku CLI
          uses: akhileshns/heroku-deploy@v3.10.9
          with:
            heroku_api_key: ${{secrets.HEROKU_API_KEY}}
            heroku_app_name: "your-heroku-app-name"
            heroku_email: "your-heroku-email"
    ```

## Deployment

1. **Deploy Backend:**

   You can deploy the Go backend to any server or PAAS. For example, using Heroku:
    ```bash
    git push heroku main
    ```

2. **Deploy Frontend:**

   Build and deploy the Nuxt frontend. For example, using Vercel:
    ```bash
    cd frontend
    vercel
    ```

## Project Structure

```
go-nuxt-app/
│
├── backend/
│   ├── main.go
│   ├── models.go
│   ├── handlers.go
│   ├── routes.go
│   ├── .env
│   └── ...
│
├── frontend/
│   ├── pages/
│   ├── components/
│   ├── nuxt.config.js
│   ├── package.json
│   ├── ...
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
