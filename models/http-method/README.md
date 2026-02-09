# HTTP Method

Provides a standardized enumeration of HTTP methods used throughout the system. While Go's standard `http` package provides method constants, this model enables database storage and consistent usage across the codebase.

## Enum: `Method`

Represents the standard HTTP methods used in REST APIs and web services.

### Values

- **`GET`** (0): Retrieve data from the server (read-only, idempotent)
- **`POST`** (1): Create new resources or submit data (may have side effects)
- **`PUT`** (2): Update or replace existing resources (idempotent)
- **`DELETE`** (3): Remove resources from the server (idempotent)

## Use Cases

### Database Storage

When HTTP method information needs to be persisted:
- API request logging
- Route configuration storage
- Permission/authorization rules
- Audit trails
- API documentation storage

### Consistency

This model provides a consistent way to represent HTTP methods across:
- Different programming languages (Go, TypeScript)
- Database schemas
- API definitions
- Configuration files

## Usage in Go

Instead of using Go's `http.MethodGet`, `http.MethodPost`, etc., use this enum for consistency, especially when:
- Storing method information in the database
- Sharing method definitions with TypeScript code
- Maintaining consistency across the codebase

## Notes

- This enum covers the most commonly used HTTP methods
- For less common methods (PATCH, OPTIONS, HEAD), consider extending this enum if needed
- The enum values correspond to standard HTTP method names
- Use this model instead of string literals to ensure type safety and consistency
