# Go Glimpse Backend

Backend service for the Glimpse platform, built with Go.

## Prerequisites

Before getting started, ensure the following tools are installed on your machine:

* Go
* Docker & Docker Compose
* Taskfile
* Bun

## Running the Application

The project uses Docker for local development. To start all required services (application, database, cache, and supporting infrastructure), run:

```bash
task docker:up-dev
```

Once the services are running, the API will be available on the configured application port.

## Stopping the Application

To stop and remove all running containers and related resources, run:

```bash
task docker:down-dev
```

## API Documentation

The OpenAPI specification is maintained in the `packages/openapi` directory.

Whenever API endpoints, request/response schemas, or documentation are updated, regenerate the API documentation by running:

```bash
cd packages/openapi
bun gen
```

This command updates the generated OpenAPI artifacts and ensures the documentation remains in sync with the codebase.

## Development Workflow

1. Start the development environment:

   ```bash
   task docker:up-dev
   ```

2. Make your code changes.

3. If API contracts have changed, regenerate the OpenAPI documentation:

   ```bash
   cd packages/openapi
   bun gen
   ```

4. Commit your changes and generated documentation files.

5. Stop the development environment when finished:

   ```bash
   task docker:down-dev
   ```

## Project Structure

```text
.
├── backend/
│   ├── cmd/                # Application entrypoints
│   ├── internal/           # Internal application code
│   ├── docker-compose.yml  # Local development services
│   └── Taskfile.yml        # Project task definitions
│
└── packages/
    └── openapi/            # OpenAPI specification and generation scripts
```