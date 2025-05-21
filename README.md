# 🔗 url-shortener

Un **servicio para acortar URLs**, escrito en Go puro y sin frameworks. Inspirado en [roadmap.sh/projects/url-shortening-service](https://roadmap.sh/projects/url-shortening-service), este proyecto me permitió aprender cómo construir una API REST desde cero, incluyendo pruebas y persistencia en base de datos.

## 🚀 ¿Qué es url-shortener?

`url-shortener` es una API que permite generar códigos cortos para URLs largas. Ofrece funcionalidades para redireccionar, contar accesos y consultar estadísticas por cada enlace generado.

## 📦 Versiones

### ✅ v1.0.0
> Primera versión funcional.
- Implementación completa de la API REST.
- Repositorio en memoria.
- Endpoints:
  - `POST /shorten`
  - `GET /:shortcode`
  - `PUT /shorten/:shortcode`
  - `GET /:shortcode/stats`
  - `DELETE /shorten/:shortcode`

### 🧪 v2.0.0
> Agregado de testing.
- Tests unitarios para handlers y lógica del repositorio.
- Código organizado para facilitar pruebas.

### 🗃️ v3.0.0
> Persistencia con MySQL.
- Reemplazo del repositorio en memoria por una implementación con MySQL.
- Script SQL de inicialización.
- Integración con Docker Compose.

## 🛠️ Instalación

```bash
git clone https://github.com/tuusuario/url-shortener.git
cd url-shortener
docker-compose up -d
```

```bash
Ingresar al container de mysql 

mysql -u root -p url_shortener (Password=root)

Ejecutar el siguiente script para crear la tabla shorts

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
Crear archivo .env con el siguiente contenido

REPOSITORY_DRIVER="mysql"
DATABASE_USER=shortener
DATABASE_PASSWORD=shortener
DATABASE_HOST=127.0.0.1
DATABASE_PORT=3306
DATABASE_NAME="url_shortener"
```

```bash
Ejecutar para lanzar la aplicación

go run .
```

## ⚙️ Endpoints principales

- `POST /shorten` – Acorta una URL.
- `GET /shorten/:shortcode` – Permite buscar los datos de una URL acortada.
- `PUT /shorten/:shortcode` – Permite actualizar la url original.
- `GET /shorten/:shortcode/stats` – Devuelve estadísticas de acceso.
- `DELETE /shorten/:shortcode` – Elimina una URL acortada.

## 🧪 Tests

```bash
go test ./...
```

## 🌱 Aprendizajes

Este proyecto me permitió:
- Entender cómo construir una API REST en Go sin frameworks.
- Diseñar una arquitectura modular desde cero.
- Aplicar testing y refactorización progresiva.
- Integrar una base de datos relacional usando MySQL.

## 📚 Tecnologías

- [Go](https://golang.org/)
- [MySQL](https://www.mysql.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## 📄 Licencia

MIT License © [Sergio Fidelis]
(https://github.com/S3ergio31)
