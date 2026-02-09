# com-fs-utils-lib

Contains generic models to create a composable proto domain

## Common build issues


### File not found

A common issue is that a file is not found. This has to do with how protoc resolves imports.

Example:

```
decimal/decimal.proto: File not found.
/Users/someuser/go/src/github.com/sologenic/com-fs-utils-lib/models/commission/commission.proto:7:1: Import "decimal/decimal.proto" was not found or had errors.
/Users/someuser/go/src/github.com/sologenic/com-fs-utils-lib/models/commission/commission.proto:11:12: "decimal.Decimal" is not defined.
```

Solution: Include the com-fs-utils-lib/models directory in the proto path.

```
protoc \
    --proto_path=. "user.proto" \
    --proto_path=$(dirname $(dirname "$rd")) \
    --proto_path=$(dirname "$rd")/com-fs-utils-lib/models \
    "--go_out=." --go_opt=paths=source_relative
```