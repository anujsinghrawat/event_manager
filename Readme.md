# 🎉 Event Manager

<!-- Optional: Add your project logo or a relevant banner image here -->
<!-- <p align="center">
  <img src="path_to_your_logo.png" alt="Event Manager Logo" width="200"/>
</p> -->

<p align="center">
  <a href="https://goreportcard.com/report/github.com/anujsinghrawat/event-manager"><img src="https://goreportcard.com/badge/github.com/anujsinghrawat/event-manager" alt="Go Report Card" /></a>
  <a href="https://pkg.go.dev/github.com/anujsinghrawat/event-manager"><img src="https://pkg.go.dev/badge/github.com/anujsinghrawat/event-manager.svg" alt="Go Reference"></a>
  <!-- Replace with actual license badge once a LICENSE file is added -->
  <a href="LICENSE.md"><img src="https://img.shields.io/badge/license-Not%20Specified-lightgrey" alt="License" /></a>
</p>

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
    └── auth-protected.go  # Auth middleware
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
git clone https://github.com/YOUR_GITHUB_USERNAME/event-manager.git
# Replace YOUR_GITHUB_USERNAME with your actual GitHub username if you've forked it,
# or use the original repository URL: https://github.com/anujsinghrawat/event-manager.git
cd event-manager
```

### Set Up Environment Variables
Create a `.env` file in the project root with the following variables (adjust as needed):
```env
# Server Configuration
SERVER_PORT=3000
APP_ENV=dev # Use 'dev' for development, 'prod' for production

# Database Configuration
# If using Docker Compose (e.g., with 'make start'), DB_HOST should typically be the service name defined in docker-compose.yml (e.g., event-manager-db)
# If running the Go application locally and connecting to a local PostgreSQL instance, DB_HOST is usually 'localhost'
DB_HOST=event-manager-db 
DB_NAME=your_db_name 
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_SSLMODE=disable # Set to 'require' or other preferred modes in production if SSL is configured for your DB

# JWT Configuration
JWT_SECRET=your_super_secret_and_strong_jwt_key # Change this to a long, random string for security
```

### Running the Application
There are two main ways to run the application: using Docker (recommended for ease of setup) or running natively with Go.

#### Using Docker (Recommended)
This project includes a `Makefile` to simplify Docker operations. Ensure Docker and Docker Compose are installed on your system.

1.  **Set up Environment Variables**: Make sure you have created and configured your `.env` file as described in the "Set Up Environment Variables" section above. The `docker-compose.yml` file is configured to pass these variables to the application container.

2.  **Start the application**:
    ```bash
    make start
    ```
    This command builds the Docker images (if not already built) and starts the application and the PostgreSQL database in detached mode.

3.  **Access the app**:
    Open your browser and navigate to: `http://localhost:YOUR_SERVER_PORT` (e.g., `http://localhost:3000` if `SERVER_PORT=3000`).

4.  **View logs**:
    ```bash
    docker-compose logs -f
    ```

5.  **Stop the application**:
    ```bash
    make stop
    ```
    This command stops and removes the containers and the network. The database volume will persist by default.

**Alternative Docker Compose Commands:**
If you prefer not to use the `Makefile`, you can use `docker-compose` commands directly:
- To build and start: `docker-compose up --build -d`
- To stop: `docker-compose down` (add `-v` to remove volumes)

#### Natively with Go (Local Development)
If you prefer to run the application directly on your machine without Docker:

1.  **Install Go**: Ensure Go (version listed in `go.mod` or newer, e.g., 1.23.5+) is installed on your system.
2.  **Set up PostgreSQL**: Make sure you have a PostgreSQL instance running and accessible. Configure the connection details (`DB_HOST`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`, `DB_SSLMODE`) in your `.env` file. For `DB_HOST`, you'll likely use `localhost`.
3.  **Set up Environment Variables**: Create and configure your `.env` file as described earlier.
4.  **Install Dependencies**:
    ```bash
    go mod tidy
    ```
5.  **Run the Application with Hot Reloading (using Air)**:
    The project uses Air for hot reloading during development.
    ```bash
    # Ensure Air is installed: go get -u github.com/cosmtrek/air (or follow Air's official installation guide)
    air -c .air.toml
    ```
6.  **Alternatively, run without hot reloading**:
    ```bash
    go run cmd/api/main.go
    ```
7.  **Access the app**:
    Open your browser and navigate to: `http://localhost:YOUR_SERVER_PORT` (e.g., `http://localhost:3000`).

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
- Docker & Docker Compose installed on your production server.
- **Secure Environment Variable Management**: A system for securely managing all required environment variables. This includes `JWT_SECRET` (which **must** be set to a strong, unique random string) and appropriately configured `DB_SSLMODE` (e.g., `require`, `verify-full` depending on your database setup) for secure database connections. All variables listed in the "Set Up Environment Variables" section must be present.
- **Production-Ready PostgreSQL**: A robust PostgreSQL instance. While the `docker-compose.yml` can spin up a PostgreSQL container, for critical production environments, consider using a managed database service or a dedicated, properly secured PostgreSQL server.

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

### Configure Environment Variables for Production
Ensure the `.env` file (or your chosen secrets management system, such as HashiCorp Vault, AWS Secrets Manager, or platform-level environment variables) is correctly set up on your production server. 

**Key variables for production:**
-   `APP_ENV=prod`
-   A strong, unique `JWT_SECRET` (critical for security).
-   Appropriate `DB_SSLMODE` (e.g., `require`, `verify-full`) for secure database connections.
-   All other database credentials (`DB_HOST`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`) and `SERVER_PORT`.

Refer to the main "Set Up Environment Variables" section for the full list and general guidance, but pay special attention to the security implications of these variables in a production context. **Do not use development default values in production, especially for `JWT_SECRET`.**

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

---

## 🤝 Contributing

Contributions are welcome and greatly appreciated! If you have an idea for an improvement or have found a bug, please feel free to contribute.

Here's how you can help:

1.  **Fork the Repository**: Click the 'Fork' button at the top right of this page.
2.  **Clone Your Fork**:
    ```bash
    git clone https://github.com/YOUR_GITHUB_USERNAME/event-manager.git
    cd event-manager
    ```
    (Replace `YOUR_GITHUB_USERNAME` with your GitHub username)
3.  **Create a New Branch**:
    ```bash
    git checkout -b feature/your-amazing-feature 
    ```
    Or for bug fixes:
    ```bash
    git checkout -b fix/issue-description
    ```
4.  **Make Your Changes**: Implement your feature or fix the bug. Ensure your code follows the project's style and best practices.
5.  **Commit Your Changes**:
    ```bash
    git add .
    git commit -m "feat: Describe your amazing feature" 
    # Or "fix: Describe the bug fix"
    # Or "docs: Improved documentation"
    ```
    (Follow conventional commit messages if possible)
6.  **Push to Your Fork**:
    ```bash
    git push origin feature/your-amazing-feature
    ```
7.  **Open a Pull Request**: Go to the original repository on GitHub and open a pull request from your forked branch. Provide a clear description of your changes.

If you're reporting a bug, please include:
-   Steps to reproduce the behavior.
-   Expected behavior.
-   Actual behavior.
-   Screenshots (if applicable).

If you're suggesting an enhancement, please outline the potential benefits.

We appreciate your effort in making this project better!

---

## 📜 License

This project is currently not licensed.

It is recommended to add a `LICENSE` file to define how others can use, modify, and distribute the code. Common open-source licenses include [MIT License](https://opensource.org/licenses/MIT), [Apache License 2.0](https://opensource.org/licenses/Apache-2.0), or [GPLv3](https://www.gnu.org/licenses/gpl-3.0.html).

Once a license is chosen and a `LICENSE` file is added to the root of the project, the license badge at the top of this README should be updated accordingly (e.g., by changing `LICENSE.md` to the correct file name if different and updating the badge text/color). The current placeholder badge indicates "Not Specified".

---

Happy coding! 🚀✨
