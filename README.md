# Email Campaign Service (email-go)

This project is a Go application for managing email campaigns. It includes several components such as campaign services, email sending infrastructure, and route handling.

## Project Structure

- `cmd/api/main.go`: The main entry point of the application.
- `internal/domain/campaign`: Contains the business logic for campaigns.
- `internal/infrastructure/mail`: Handles email sending functionalities.
- `internal/infrastructure/repository`: Manages database interactions.
- `internal/routes`: Defines HTTP routes and handlers.

## Prerequisites

- Go (1.18 or later)
- Docker (optional, for running a local database)
- Make sure to have a `.env` file at the root of your project directory with the necessary environment variables.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/thiagoadsix/email-go.git
    cd email-campaign-service
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

3. Set up your environment variables in a `.env` file:
    ```sh
    touch .env
    ```

    Add the following content to the `.env` file:
    ```env
    DATA_BASE_URL="host=localhost user=emailn_dev password=12345678 port=5432 sslmode=disable"
    KEYCLOAK_URL="http://localhost:8080/realms/emailn_realm"
    GOMAIL_SMTP="smtp.gmail.com"
    GOMAIL_USER="your_email_test@gmail.com"
    GOMAIL_PASS="your_app_password_google_account"
    ```

## Running the Project

1. Start the application:
    ```sh
    go run cmd/api/main.go
    ```

2. The application will be running on `http://localhost:8080`.

## API Endpoints

- `POST /campaigns`: Create a new campaign.
- `GET /campaigns`: Retrieve a list of campaigns.
- `GET /campaigns/{id}`: Retrieve a specific campaign by ID.
- `PUT /campaigns/{id}`: Update a specific campaign by ID.
- `DELETE /campaigns/{id}`: Delete a specific campaign by ID.
- `POST /campaigns/{id}/send`: Send a campaign.

## Usage

- Make sure your database is up and running.
- Use an API client like Postman to interact with the endpoints.

## Keycloak for Authentication

We are using [Keycloak](https://www.keycloak.org/) for authentication and authorization. Keycloak is an open-source Identity and Access Management solution aimed at modern applications and services.

## Libraries Used

- `github.com/coreos/go-oidc/v3 v3.10.0`: OpenID Connect (OIDC) client implementation.
- `github.com/go-chi/chi/v5 v5.0.14`: Lightweight, idiomatic, and composable router for building Go HTTP services.
- `github.com/go-chi/render v1.0.3`: Rendering utilities for HTTP responses.
- `github.com/go-playground/validator/v10 v10.22.0`: Go (golang) package for data validation.
- `github.com/golang-jwt/jwt/v5 v5.2.1`: A Go implementation of JSON Web Tokens (JWT).
- `github.com/jaswdr/faker v1.19.1`: Faker is a library for generating fake data.
- `github.com/joho/godotenv v1.5.1`: Go port of Ruby's dotenv library (loads environment variables from `.env`).
- `github.com/rs/xid v1.5.0`: Xid is a globally unique id generator.
- `github.com/stretchr/testify v1.9.0`: A toolkit with common assertions and mocks that plays nicely with the standard library.
- `gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df`: A simple and efficient package to send emails.
- `gorm.io/driver/postgres v1.5.9`: GORM PostgreSQL driver.
- `gorm.io/gorm v1.25.10`: The fantastic ORM library for Golang.


## Contributing

1. Fork the repository.
2. Create a new feature branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.