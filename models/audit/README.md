# Audit

The audit proto is used to represent generic audit information (who, when, why) reusable across different subjects (e.g., Role, Rolepaths, User Status, etc.).

## Message: `Audit`

The `Audit` message provides a standardized way to track changes across the system.

### Fields

- **`ChangedBy`** (string, required): The identifier of the user or system that made the change
- **`ChangedAt`** (google.protobuf.Timestamp, required): The timestamp when the change occurred
- **`Reason`** (string, optional): An optional explanation or reason for the change

### Usage

This model can be embedded in other proto messages that need audit trail functionality. It provides a consistent way to track:
- Who made a change
- When the change was made
- Why the change was made (if provided)

### Example Use Cases

- Tracking role changes
- Recording user status updates
- Auditing configuration modifications
- Logging administrative actions

### Notes

- The `Reason` field is optional, allowing for cases where an explanation may not be required or available
- The `ChangedAt` timestamp uses Google's protobuf Timestamp type for cross-language compatibility