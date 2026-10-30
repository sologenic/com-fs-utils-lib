# Metadata

Provides a standardized set of metadata fields that are commonly required across multiple proto messages. This model helps avoid repetition and ensures consistency in tracking timestamps and change tracking.

## Message: `MetaData`

Contains common metadata fields used across various messages in the system.

### Fields

- **`UpdatedAt`** (google.protobuf.Timestamp, required): Timestamp of the last update
- **`CreatedAt`** (google.protobuf.Timestamp, required): Timestamp of creation (internal to listener)
- **`UpdatedByAccount`** (string, optional): Account identifier that performed the last update

## Enum: `Network`

Represents the blockchain network environment. Used with the `NETWORK` environment variable and `GetNetwork()`, including Datastore key segments that already embed the network integer.

### Values

- **`NETWORK_DO_NOT_USE`** (0): Reserved default value to avoid bugs from unset network values
- **`MAINNET`** (1): Production mainnet network
- **`TESTNET`** (2): Test network for development and testing

### Why Avoid Default Value

The `NETWORK_DO_NOT_USE` value at 0 ensures that an unset network is distinguishable from a bug. This prevents accidental use of default values that could lead to production data being written to test networks or vice versa.

## Usage

This metadata model can be embedded in other proto messages that need:
- Creation and update timestamps
- Change tracking (who updated)

### Example Embedding

```protobuf
message MyMessage {
    // ... other fields ...
    metadata.MetaData Metadata = 10;
}
```

## Use Cases

- Database record metadata
- API response metadata
- Event tracking
- Audit logging
- Change tracking

## Notes

- `CreatedAt` is marked as "internal to listener" - it's typically set by the system/listener, not by user input
- `UpdatedByAccount` is optional, allowing for system-initiated updates where no account is involved
- Always use `NETWORK_DO_NOT_USE` as the default to catch uninitialized values
