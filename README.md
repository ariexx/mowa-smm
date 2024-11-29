
# Mowa Backend

The backend service for Mowa, built with **Golang**, providing robust and efficient API functionalities. This repository serves as the backend logic for the Mowa application, handling data processing, authentication, and API endpoints.

---

## Features

- RESTful API architecture
- Authentication and authorization
- Data processing and validation
- Scalable and modular design
---

## Prerequisites

Ensure you have the following installed:

- [Golang](https://golang.org/dl/) (Version 1.19 or higher recommended)
- [Docker](https://www.docker.com/) (Optional but recommended for deployment)

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ariexx/mowa-backend.git
   cd mowa-backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the environment variables:  
   Create a `.env` file in the root directory and configure the following variables:
   ```env
   PORT=8080
   APP_ENV=local
   BLUEPRINT_DB_HOST=localhost
   BLUEPRINT_DB_PORT=3306
   BLUEPRINT_DB_DATABASE=mowa
   BLUEPRINT_DB_USERNAME=root  
   BLUEPRINT_DB_PASSWORD=
   BLUEPRINT_DB_ROOT_PASSWORD= 
   JWT_SECRET=
   ```

4. Run the application:
   ```bash
   go run main.go
   ```
---

## Testing

Run unit tests to ensure functionality:
```bash
go test ./...
```

---

## Deployment

### Docker

1. Build the Docker image:
   ```bash
   docker build -t mowa-backend .
   ```

2. Run the container:
   ```bash
   docker run -d -p 8080:8080 --env-file .env mowa-backend
   ```

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add new feature"
   ```
4. Push to your fork:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Author

**Arief**  
Connect with me on [GitHub](https://github.com/ariexx) or explore my other [projects](https://github.com/ariexx).
