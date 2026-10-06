# OKOK Student Manager API

Minimal Echo API using ports and adapters. Health checks are available at `GET /health` and `GET /api/v1/health`.

## Project structure

```text
cmd/api/                              Application entrypoint
internal/adapter/http/                Echo handlers and HTTP DTOs
internal/adapter/process/             Process health probe
internal/bootstrap/                   HTTP setup and dependency wiring
internal/config/                      Viper configuration
internal/domain/                      Domain health types
internal/port/                        Application and dependency interfaces
internal/service/                     Use-case orchestration
docs/postman/                         Postman collection
config.yaml                           Root configuration file
```

The health request flow is `HTTP adapter → health service port → service → probe port → process adapter`. The HTTP adapter maps the domain result to the response DTO and adds the timestamp. Dependencies are constructed in `internal/bootstrap` and point toward the application core.

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
