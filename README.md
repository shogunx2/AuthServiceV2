# AuthServiceV2

AuthServiceV2 is a microservice designed as a gateway for authenticating users into other products or websites. Built primarily for learning and experimentation, it demonstrates a Go backend and a JavaScript frontend with future plans for a React-based UI. The service supports user authentication via password or API key and stores user data in a PostgreSQL database.

## Features

- Acts as an authentication gateway for other services
- Users can authenticate via:
  - UserID & Password
  - UserID & API Key
- Backend written in Go
- Frontend written in JavaScript (planned migration to React)
- PostgreSQL integration for persistent storage

## Tech Stack

- **Backend:** Go
- **Frontend:** JavaScript (planned React migration)
- **Database:** PostgreSQL

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed
- [Node.js](https://nodejs.org/) and npm (for frontend)
- [PostgreSQL](https://www.postgresql.org/) running and accessible

### Setup

1. **Clone the repository**
   ```sh
   git clone https://github.com/shogunx2/AuthServiceV2.git
   cd AuthServiceV2
   ```

2. **Start the PostgreSQL server**
   - Ensure your DB connection details are configured properly (see backend config).

3. **Run the backend**
   ```sh
   cd backend
   go run server.go
   ```

4. **Run the frontend**
   ```sh
   cd ../frontend
   # If using plain JS:
   # Open index.html in your browser

   # (If/when React is implemented:)
   npm install
   npm start
   ```

### Configuration

- Adjust database credentials and other settings as needed in backend configuration files or environment variables.

## Usage

- Access the web frontend to log in or register via UserID & Password or UserID & API Key.
- The current frontend is a basic JavaScript interface; a React UI is planned to provide a more modern user experience.

## Roadmap

- [ ] Upgrade frontend to React and improve user experience
- [ ] Enable dynamic choice between UserID/Password and UserID/API Key on the frontend
- [ ] Add frontend validation and better error feedback
- [ ] Dockerize backend and database for easier deployment
- [ ] Add API documentation (OpenAPI/Swagger)
- [ ] Improve security and production readiness

## Intended Use

This project was created as a learning tool for Go and web microservices. While functional, it is not intended for production use without further hardening and improvements.

## Contributing

Contributions and suggestions are welcome! Please open issues or pull requests as you experiment with or extend the service.

## License

This project is open source and available under the [MIT License](LICENSE).
