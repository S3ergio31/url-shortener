
# 🔗 url-shortener

A **URL shortening service**, written in pure Go with no frameworks. Inspired by [roadmap.sh/projects/url-shortening-service](https://roadmap.sh/projects/url-shortening-service), this project allowed me to learn how to build a REST API from scratch, including testing and database persistence.

## 🚀 What is url-shortener?

`url-shortener` is an API that allows generating short codes for long URLs. It provides functionality to redirect, count accesses, and retrieve statistics for each generated link.

## 📦 Versions

### ✅ v1.0.0
> First functional version.
- Full implementation of the REST API.
- In-memory repository.
- Endpoints:
  - `POST /shorten`
  - `GET /:shortcode`
  - `PUT /shorten/:shortcode`
  - `GET /:shortcode/stats`
  - `DELETE /shorten/:shortcode`

### 🧪 v2.0.0
> Added testing.
- Unit tests for handlers and repository logic.
- Code organized to facilitate testing.

### 🗃️ v3.0.0
> MySQL persistence.
- Replaced the in-memory repository with a MySQL implementation.
- SQL initialization script.
- Integrated with Docker Compose.

## 🛠️ Installation

```bash
git clone https://github.com/tuusuario/url-shortener.git
cd url-shortener
docker-compose up -d
```

```bash
Access the MySQL container

mysql -u root -p url_shortener (Password=root)

Run the following script to create the `shorts` table

CREATE TABLE IF NOT EXISTS shorts (
    id CHAR(36) NOT NULL PRIMARY KEY,
    url TEXT NOT NULL,
    short_code VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    access_count INT NOT NULL DEFAULT 0
);
```

```bash
Create a .env file with the following content

REPOSITORY_DRIVER="mysql"
DATABASE_USER=shortener
DATABASE_PASSWORD=shortener
DATABASE_HOST=127.0.0.1
DATABASE_PORT=3306
DATABASE_NAME="url_shortener"
```

```bash
Run the application

go run .
```

## ⚙️ Main Endpoints

- `POST /shorten` – Shortens a URL.
- `GET /shorten/:shortcode` – Retrieves the data of a shortened URL.
- `PUT /shorten/:shortcode` – Updates the original URL.
- `GET /shorten/:shortcode/stats` – Returns access statistics.
- `DELETE /shorten/:shortcode` – Deletes a shortened URL.

## 🧪 Tests

```bash
go test ./...
```

## 🌱 Learnings

This project allowed me to:
- Understand how to build a REST API in Go without frameworks.
- Design a modular architecture from scratch.
- Apply testing and progressive refactoring.
- Integrate a relational database using MySQL.

## 📚 Technologies

- [Go](https://golang.org/)
- [MySQL](https://www.mysql.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## 📄 License

MIT License © [Sergio Fidelis](https://github.com/S3ergio31)
