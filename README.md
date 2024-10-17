# Comic Chest - Backend

Comic Chest is a backend server designed to manage a comic, manga, and graphic novel application. This project provides a RESTful API that allows users to perform CRUD operations on graphic novels and their chapters, offering a complete solution for managing and viewing comic-related content.

## Features

- **CRUD for Graphic Novels and Chapters:** Users can create, read, update, and delete graphic novels (comics, manga, manhwa) and manage their respective chapters.
- **RESTful API:** The architecture follows a REST-based design, ensuring a simple and standardized interface for server interaction.
- **Authentication and Authorization:** Implemented with stateful tokens, providing secure user session management.
- **Email Management:** Integrated support for sending emails, used for account activation.
- **CORS Enabled:** Cross-Origin Resource Sharing (CORS) is configured to allow requests from different domains, which is essential for interacting with frontend applications.
- **Database Management:** Persistent storage for graphic novel and user data, using Postgres as database.
- **Docker Compose:** The project includes a docker-compose.yml file for easy setup and deployment of the application. Users only need to define their environment variables for development and production environments.

## Installation Requirements

- **1 - Docker and Docker Compose:** Ensure Docker and Docker Compose are installed on your system.
- **2 - Environment Variables:** Set up the following environment variables before launching the containers:
    - DB_DSN (no need to change)
    - SMTP_HOST
    - SMTP_PORT
    - SMTP_USERNAME
    - SMTP_PASSWORD
    - SMTP_SENDER

- **3 - Database Variables:** User, password and database name (no need to change) 

## Getting started

```bash
    git clone https://github.com/Juanmagc99/comic-chest.git
    docker compose up
```
## Available Endpoints
The API includes the following main endpoints:

### Graphic Novels (gnovels):
- `GET /gnovels`: Retrieves a list of graphic novels.
- `POST /gnovels`: Creates a new graphic novel.
- `GET /gnovels/{id}`: Retrieves a graphic novel by ID.
- `PATCH /gnovels/{id}`: Updates an existing graphic novel.
- `DELETE /gnovels/{id}`: Deletes a graphic novel.

### Chapters:
- `POST /gnovels/{id}/chapter`: Uploads a chapter for a specific graphic novel.
- `GET /gnovels/{id}/chapter/{chapter_num}`: Retrieves a chapter by number.
- `DELETE /gnovels/{id}/chapter/{chapter_num}`: Deletes a specific chapter.

### Users:
- `POST /users`: Registers a new user.
- `PUT /users/activated`: Activates a user account.

### Authentication:
- `POST /tokens/authentication`: Generates an authentication token to access protected resources.

### Health Check:
- `GET /healthcheck`: Checks that the server is running properly.

## License

[MIT](https://choosealicense.com/licenses/mit/)