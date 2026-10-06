# OKOK Student Manager API

Minimal Echo API organized into handler and service layers. Health checks are available at `GET /health` and `GET /api/v1/health`.

## Project structure

```text
cmd/api/             Application entrypoint
internal/bootstrap/  Echo setup and dependency wiring
internal/config/     Viper configuration
internal/handler/    HTTP handlers and response DTOs
internal/service/    Health check logic
docs/postman/        Postman collection
config.yaml          Root configuration file
```

The health request flow is `handler → service`. The service performs the health check; the handler builds the HTTP response DTO and timestamp.

## Run

```sh
go mod tidy
go run ./cmd/api
```

Configuration is read from the root `config.yaml` file and can be overridden with `APP_HOST`, `APP_PORT`, and `APP_ENV`. Defaults are `0.0.0.0:8080` and `development` if the config file is absent.

## Postman

Import [`docs/postman/okok-student-manager.postman_collection.json`](docs/postman/okok-student-manager.postman_collection.json) into Postman. Its `baseUrl` collection variable defaults to `http://localhost:8080`.

The collection includes both health endpoints and an unknown route request for the JSON 404 response.

## Responses

Successful health response:

```json
{"code":"OK","message":"service is healthy","data":{"status":"ok","timestamp":"2026-01-01T00:00:00Z"}}
```

Unknown routes return JSON with the HTTP status in `code` and its standard text in `message`.
