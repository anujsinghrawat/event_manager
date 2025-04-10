# 🎉 Event Manager

**Event Manager** is a modern, robust, and fully-featured event management portal built in Go with [Fiber](https://gofiber.io). It allows users to authenticate, create and manage events, and even generate tickets with QR codes. Perfect for organizing events, tracking attendance, and providing smooth user experiences!

---

## 🚀 Features

-   **User Authentication**
    -   Register & Login endpoints 🔐
    -   Secure JWT authentication via custom middleware
-   **Event Management**
    -   Create, update, delete, and fetch events 🗓️
    -   RESTful API endpoints for listing and viewing event details
-   **Ticketing System**
    -   Ticket creation and validation 🎟️
    -   Integrated QR code generation for every ticket
-   **Beautiful UI**
    -   A welcoming root route that returns a styled HTML page
-   **Robust & Scalable Architecture**
    -   Clear separation of concerns with handlers, services, repositories, and middlewares
-   **Dockerized Deployment**
    -   Ready-to-use Docker and Docker Compose configurations for local and production environments
-   **Hot Reloading (Development)**
    -   [Air](https://github.com/cosmtrek/air) configuration for a seamless coding experience 💻

---

## 📂 Project Structure

```plaintext
.
├── cmd/                 # Entry point(s) for your application
├── config/              # Environment configuration and settings
├── db/                  # Database initialization and migration code
├── handlers/            # HTTP handler functions for API endpoints
│   ├── auth.go          # Authentication routes (login, register)
│   ├── events.go        # Event CRUD endpoints
│   └── tickets.go       # Ticket management and QR code generation
├── middlewares/         # Custom middleware such as auth protection
    └── auth-proctected.go  # Auth middleware
├── models/              # Data models (Event, Ticket, Auth, etc.)
├── repositories/        # Database interaction logic for various models
    ├── auth.go          # Auth repository
    ├── event.go         # Event repository
    └── ticket.go        # Ticket repository
├── services/            # Business logic and service layer (auth, etc.)
├── docker-compose.yml   # Docker Compose config for the app and PostgreSQL DB
├── Dockerfile           # Dockerfile to containerize the application
├── Makefile             # Useful commands for running tests, building, etc.
├── .env                 # Environment variables (DB credentials, etc.)
└── .air.toml            # Air configuration for hot reloading during development
```

---

## 🛠️ Getting Started

Follow these steps to set up the project on your local machine:

### Clone the Repository
```bash
git clone https://github.com/your_username/event-manager.git
cd event-manager
```

### Set Up Environment Variables
Create a `.env` file in the project root with the following variables (adjust as needed):
```env
DB_HOST=your_db_host
DB_NAME=your_db_name
DB_USER=your_db_user
DB_PASSWORD=your_db_password
SERVER_PORT=3000
APP_ENV=dev
```

### Install Dependencies
Make sure you have Go installed, then install the necessary packages:
```bash
go mod tidy
```

### Run the Application (Development Mode)
The project uses Air for hot reloading during development:
```bash
# Ensure Air is installed: go get -u github.com/cosmtrek/air
air -c .air.toml
```

Alternatively, you can run the app using:
```bash
go run main.go
```

### Access the App
Open your browser and navigate to: http://localhost:3000 to see the beautiful welcome page.

---

## 🔌 API Routes

All endpoints are prefixed with `/api/`. Below is an overview of the main routes:

### Authentication
- **POST /api/auth/login**: Login an existing user by providing credentials.
- **POST /api/auth/register**: Register a new user.

### Event Endpoints (Protected)
- **GET /api/events/**: Fetch all events.
- **GET /api/events/:eventId**: Fetch an event by its ID.
- **POST /api/events/**: Create a new event.
- **PUT /api/events/:eventId**: Update an existing event.
- **DELETE /api/events/:eventId**: Delete an event.

### Ticket Endpoints (Protected)
- **GET /api/tickets/**: Retrieve all tickets for the authenticated user.
- **GET /api/tickets/:ticket_id**: Fetch details of a specific ticket along with a generated QR code.
- **POST /api/tickets/**: Create a new ticket for an event.
- **POST /api/tickets/validate**: Validate a ticket at the event entrance.

---

## 🚢 Production Deployment

Follow these steps to deploy Event Manager in a production environment:

### Prerequisites
- Docker & Docker Compose installed on your production server
- Properly configured environment variables (use a secure mechanism for storing secrets)
- A production-ready PostgreSQL instance (or use the provided container if suitable)

### Build the Production Binary and Docker Image

**Option 1: Build Locally then Containerize**
```bash
# Build the Go binary
go build -o main .

# Build the Docker image
docker build -t event-manager .
```

**Option 2: Use Docker Compose for Build**
The provided `docker-compose.yml` will automatically build your image using the Dockerfile.

### Configure Environment Variables
Ensure the `.env` file (or your chosen secrets manager) is correctly set up on your production server with the required settings (DB credentials, SERVER_PORT, etc.).

### Deploy Using Docker Compose
In the project root, run:
```bash
docker-compose up --build -d
```

This command will:
- Build the application container (event-manager)
- Initialize the PostgreSQL container
- Run both containers in the background

### Verify Deployment
Confirm the containers are running:
```bash
docker ps
```

Check logs if needed:
```bash
docker-compose logs -f
```

Test the endpoints using a tool like Postman or cURL.

### Maintenance & Updates
To stop the application:
```bash
docker-compose down
```

To remove all containers and images (if needed):
```bash
docker-compose rm -v --force --stop
docker rmi event-manager
```

---

## 🎨 Custom Styles & Emojis

This README utilizes custom emojis and styling to create an engaging and clear guide. Use similar styling and icons within your project documentation and UI where possible to enhance the user experience!

## 📚 Additional Resources

- [Fiber Documentation](https://docs.gofiber.io/)
- [Go Documentation](https://golang.org/doc/)
- [Docker Documentation](https://docs.docker.com/)
- [Air Hot Reloading](https://github.com/cosmtrek/air)

Happy coding! 🚀✨

This documentation is maintained by the Event Manager team. For further questions or contributions, please submit an issue or pull request on GitHub.