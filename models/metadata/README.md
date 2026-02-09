# Metadata

Provides a standardized set of metadata fields that are commonly required across multiple proto messages. This model helps avoid repetition and ensures consistency in tracking network context, timestamps, and change tracking.

## Message: `MetaData`

Contains common metadata fields used across various messages in the system.

### Fields

- **`Network`** (Network, required): The blockchain network context (mainnet, testnet, devnet, or custom virtual networks)
- **`UpdatedAt`** (google.protobuf.Timestamp, required): Timestamp of the last update
- **`CreatedAt`** (google.protobuf.Timestamp, required): Timestamp of creation (internal to listener)
- **`UpdatedByAccount`** (string, optional): Account identifier that performed the last update

### Network Context

The `Network` field supports:
- Standard networks: mainnet, testnet, devnet
- Custom virtual networks: Extra devnet instances for testing, additional mainnet nodes for historical block scanning, etc.

## Enum: `Network`

Represents the blockchain network environment.

### Values

- **`NETWORK_DO_NOT_USE`** (0): Reserved default value to avoid bugs from unset network values
- **`MAINNET`** (1): Production mainnet network
- **`TESTNET`** (2): Test network for development and testing
- **`DEVNET`** (3): Development network for local development

### Why Avoid Default Value

The `NETWORK_DO_NOT_USE` value at 0 ensures that an unset network is distinguishable from a bug. This prevents accidental use of default values that could lead to production data being written to test networks or vice versa.

## Usage

This metadata model can be embedded in other proto messages that need:
- Network context tracking
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
- Multi-network support
- Change tracking

## Notes

- `CreatedAt` is marked as "internal to listener" - it's typically set by the system/listener, not by user input
- `UpdatedByAccount` is optional, allowing for system-initiated updates where no account is involved
- The network field enables multi-network deployments and testing scenarios
- Always use `NETWORK_DO_NOT_USE` as the default to catch uninitialized values
