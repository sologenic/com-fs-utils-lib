# Order Properties

This proto contains order-related properties and enums that are reused across multiple models in the trading system. These properties define order characteristics, execution parameters, and processing states.

## Enum: `Side`

Represents the order side (buy or sell).

### Values

- **`SIDE_NOT_APPLICABLE`** (0): Reserved default value to protect against unexpected settings
- **`SIDE_BUY`** (1): Buy order - purchases base_denom quantity at the specified price
- **`SIDE_SELL`** (2): Sell order - sells base_denom quantity at the specified price

### Usage

The side determines the direction of the trade:
- **BUY**: Acquiring the base denomination by spending the quote denomination
- **SELL**: Disposing of the base denomination to receive the quote denomination

## Enum: `TimeInForce`

Defines how long an order remains active in the market.

### Values

- **`NOT_USED_TIME_IN_FORCE`** (0): Default/unused value
- **`DAY`** (1): Order is valid for the trading day only
- **`GOOD_TIL_CANCELED`** (2): GTC - Order remains active until filled or manually canceled
- **`AT_THE_OPENING`** (3): OPG - Order executes at the market opening
- **`AT_THE_CLOSE`** (4): CLS - Order executes at the market close
- **`IMMEDIATE_OR_CANCEL`** (5): IOC - Order must be filled immediately or canceled
- **`FILL_OR_KILL`** (6): FOK - Order must be filled completely immediately or canceled
- **`GOOD_TIL_TIME`** (7): GTT - Order is valid until a specified time

## Enum: `OrderType`

Represents the type of order being placed.

### Values

- **`NOT_APPLICABLE_ORDER_TYPE`** (0): Default/unused value (changed from "not used" to avoid naming conflicts)
- **`ORDER_TYPE_PURCHASE`** (1): Purchase order type (renamed to avoid conflicts with PURCHASE in action enum)
- **`ORDER_TYPE_SELL`** (2): Sell order type

### Naming Note

The enum values use the `ORDER_TYPE_` prefix to avoid naming conflicts with other enums (specifically the action enum's `PURCHASE` value).

## Message: `ProcessInfo`

Tracks the processing state and timestamp of an order.

### Fields

- **`ProcessState`** (ProcessState, required): Current state of order processing
- **`ProcessedAt`** (google.protobuf.Timestamp, required): Timestamp when the order reached its current state

## Enum: `ProcessState`

Represents the current processing state of an order.

### Values

- **`NOT_USED_PROCESS_STATE`** (0): Default/unused value
- **`PROCESS_STATE_OPEN`** (1): Order is open and active
- **`PROCESS_STATE_LOCKED`** (2): Order is locked (cannot be modified)
- **`PROCESS_STATE_FAILED`** (3): Order processing failed
- **`PROCESS_STATE_PROCESSED`** (4): Order has been successfully processed

## Enum: `ClearingBroker`

Identifies the clearing broker responsible for order settlement.

### Values

- **`NOT_USED_CLEARING_BROKER`** (0): Default/unused value
- **`ALPACA`** (1): Alpaca clearing broker
- **`RQD`** (2): RQD clearing broker

## Use Cases

- Order creation and management
- Trade execution tracking
- Order book operations
- Clearing and settlement
- Order lifecycle management
- Multi-broker support

## Notes

- These enums are designed to be embedded in various order-related messages
- The `SIDE_NOT_APPLICABLE` default helps catch uninitialized order sides
- `TimeInForce` values follow standard trading industry conventions
- `ProcessState` enables tracking of order lifecycle from creation to completion
