# History Hub

A lightweight REST API that serves **"On This Day"** historical events by proxying and caching the [Wikimedia REST API](https://api.wikimedia.org/wiki/Feed_API/Reference/On_this_day). Built with Go, Gin, and Redis.

## Features

- Browse historical events, births, deaths, and holidays for any date
- Multi-language support (English & Turkish)
- Redis caching with configurable TTLs (shorter for today, longer for past dates)
- Request timeout middleware
- Docker-ready with multi-stage build

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go 1.26 |
| HTTP Framework | [Gin](https://github.com/gin-gonic/gin) |
| Cache | [Redis](https://redis.io/) via go-redis |
| Data Source | [Wikimedia Feed API](https://api.wikimedia.org/wiki/Feed_API/Reference/On_this_day) |
| Config | godotenv |

## Project Structure

```
history-hub/
├── cmd/server/             # Application entrypoint
├── internal/
│   ├── cache/              # Redis client wrapper & cache key builder
│   ├── config/             # Environment config loader
│   ├── http/
│   │   ├── handlers/       # HTTP handlers (events, health)
│   │   └── router.go       # Gin router & middleware setup
│   ├── model/              # Request/response DTOs
│   ├── provider/wikimedia/ # Wikimedia API client
│   └── service/            # Business logic (cache-aside pattern)
├── pkg/utils/              # Shared helpers (validation, formatting)
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

## API Endpoints

### `GET /health`

Health check endpoint.

**Response:**
```json
{
  "timestamp": "2026-04-15T12:00:00Z",
  "message": "success"
}
```

### `GET /events/today`

Returns historical events for today's date.

| Query Param | Required | Default | Values |
|-------------|----------|---------|--------|
| `type` | No | `events` | `events`, `births`, `deaths`, `holidays`, `selected`, `all` |
| `language` | No | `en` | `en`, `tr` |

### `GET /events/list`

Returns historical events for a specific date.

| Query Param | Required | Default | Values |
|-------------|----------|---------|--------|
| `month` | Yes | - | `1`–`12` |
| `day` | Yes | - | `1`–`31` |
| `type` | No | `events` | `events`, `births`, `deaths`, `holidays`, `selected`, `all` |
| `language` | No | `en` | `en`, `tr` |

### Response Format

```json
{
  "date": "2026-04-15",
  "lang": "en",
  "type": "events",
  "cached": true,
  "events": [
    {
      "year": 1912,
      "text": "The RMS Titanic sinks...",
      "title": "Sinking of the Titanic",
      "url": "https://en.wikipedia.org/wiki/Sinking_of_the_Titanic"
    }
  ]
}
```

## Getting Started

### Prerequisites

- Go 1.26+
- Redis

### Environment Variables

Copy the example file and fill in the values:

```bash
cp .env.example .env
```

| Variable | Default | Description |
|----------|---------|-------------|
| `APP_PORT` | `8080` | Application port |
| `APP_ENV` | `development` | Environment name |
| `REDIS_HOST` | - | Redis host address |
| `REDIS_PORT` | - | Redis port |
| `REDIS_PASSWORD` | - | Redis password (optional) |
| `WIKIMEDIA_BASE_URL` | `https://api.wikimedia.org` | Wikimedia API base URL |
| `CACHE_TTL_TODAY_HOURS` | `12` | Cache TTL for today's events (hours) |
| `CACHE_TTL_PAST_HOURS` | `168` | Cache TTL for past dates (hours) |

### Run Locally

```bash
go run ./cmd/server
```

The server starts on `http://localhost:8080`.

### Run with Docker Compose

```bash
docker compose up --build
```

This starts both the API (port 8080) and Redis (port 6379).

## Caching Strategy

Responses are cached in Redis with the key pattern:

```
historyhub:onthisday:{lang}:{type}:{month}:{day}
```

- **Today's date**: cached for `CACHE_TTL_TODAY_HOURS` (default 12h) since new events can be curated.
- **Past dates**: cached for `CACHE_TTL_PAST_HOURS` (default 168h / 7 days) since they rarely change.

## Related

- [history-hub-ui](https://github.com/kaderkaya/history-hub-ui) — Next.js frontend for this API with a story-style scrolling feed.
