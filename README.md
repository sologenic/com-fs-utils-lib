# com-fs-utils-lib

Contains generic models to create a composable proto domain

## Common build issues


### File not found

A common issue is that a file is not found. This has to do with how protoc resolves imports.

Example:

```
google/type/decimal.proto: File not found.
/Users/someuser/go/src/github.com/sologenic/com-fs-utils-lib/models/commission/commission.proto:7:1: Import "google/type/decimal.proto" was not found or had errors.
/Users/someuser/go/src/github.com/sologenic/com-fs-utils-lib/models/commission/commission.proto:11:12: "google.type.Decimal" is not defined.
```

Solution: Include the dependencies directory (where `google/type/decimal.proto` is fetched by `models/bin/build.sh`) in the proto path, and include com-fs-utils-lib/models for other shared models.

```
protoc \
    --proto_path=. "user.proto" \
    --proto_path=$(dirname $(dirname "$rd")) \
    --proto_path=$(dirname "$rd")/com-fs-utils-lib/models \
    --proto_path=$rd/dependencies \
    "--go_out=." --go_opt=paths=source_relative
```
