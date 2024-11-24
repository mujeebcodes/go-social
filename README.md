# GoSocial API

GoSocial is a robust and scalable social media API built with Go (Golang). It provides backend services and functionalities required for building a modern social media platform, including user authentication, posts, comments, likes, and more.

---

## Features

- **User Authentication:**
  - Register, log in, and log out.
  - Password encryption and secure session management.
- **Posts and Media:**

  - Create, edit, and delete posts.
  - Upload and manage images/videos for posts.

- **Engagement:**

  - Like and comment on posts.
  - Follow/unfollow users.

- **Notifications:**

  - Real-time notifications for likes, comments, and follows.

- **Search and Discover:**

  - Search for users and hashtags.
  - Trending posts and popular hashtags.

- **Security:**
  - Token-based authentication with JWT.
  - Rate limiting to prevent abuse.

---

## Getting Started

### Prerequisites

- Go 1.20+ installed
- PostgreSQL database (or your preferred database system)
- Redis (for caching and real-time notifications)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-social.git
   cd go-social
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up environment variables. Create a `.env` file in the project root:

   ```
   DATABASE_URL=postgres://username:password@localhost:5432/gosocial
   REDIS_URL=localhost:6379
   JWT_SECRET=your_jwt_secret
   ```

4. Run database migrations:

   ```bash
   go run cmd/migrate/main.go
   ```

5. Start the server:
   ```bash
   go run cmd/api/main.go
   ```

---

## API Endpoints

### Authentication

| Method | Endpoint         | Description         |
| ------ | ---------------- | ------------------- |
| POST   | `/auth/register` | Register a new user |
| POST   | `/auth/login`    | Log in a user       |
| POST   | `/auth/logout`   | Log out a user      |

### Posts

| Method | Endpoint     | Description       |
| ------ | ------------ | ----------------- |
| GET    | `/posts`     | Get all posts     |
| POST   | `/posts`     | Create a new post |
| PUT    | `/posts/:id` | Edit a post       |
| DELETE | `/posts/:id` | Delete a post     |

### Comments

| Method | Endpoint              | Description             |
| ------ | --------------------- | ----------------------- |
| POST   | `/posts/:id/comments` | Add a comment to a post |
| GET    | `/posts/:id/comments` | Get comments for a post |

### Likes

| Method | Endpoint           | Description   |
| ------ | ------------------ | ------------- |
| POST   | `/posts/:id/likes` | Like a post   |
| DELETE | `/posts/:id/likes` | Unlike a post |

### Users

| Method | Endpoint            | Description         |
| ------ | ------------------- | ------------------- |
| GET    | `/users/:id`        | Get user profile    |
| PUT    | `/users/:id`        | Update user profile |
| POST   | `/users/:id/follow` | Follow a user       |
| DELETE | `/users/:id/follow` | Unfollow a user     |

---

## Technologies Used

- **Backend Framework:** Go (Golang)
- **Database:** PostgreSQL
- **Caching:** Redis
- **Authentication:** JWT (JSON Web Tokens)
- **Testing:** Go's built-in testing package
- **API Documentation:** Swagger/OpenAPI

---

## Contributing

We welcome contributions! To get started:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m "Add feature"`).
4. Push to the branch (`git push origin feature-name`).
5. Open a Pull Request.

---
