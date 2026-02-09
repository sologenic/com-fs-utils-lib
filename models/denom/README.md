# Denom (Denomination)

Represents cryptocurrency denominations (tokens) used in trading and financial operations. While denominations are typically represented as strings in standard notation, this model provides a structured format for easier data processing.

## Message: `Denom`

Represents a single denomination (token) with its identifying information.

### Fields

- **`Currency`** (string, required): The currency code or symbol (e.g., "XRP", "BTC", "USD")
- **`Issuer`** (string, required): The issuer address or identifier for the token
- **`Precision`** (int32, optional): The decimal precision (number of decimal places) for this denomination
- **`IsIBC`** (bool, required): Indicates whether this is an Inter-Blockchain Communication (IBC) token

### Standard Notation

In standard notation, a denomination is typically represented as a string combining currency and issuer:
- Format: `{Currency}.{Issuer}` or `{Currency}/{Issuer}`
- Example: `XRP.rIssuerAddress123...` or `USD/rBankAddress456...`

This model splits that string into structured fields for easier processing and validation.

## Message: `TokenPair`

Represents a trading pair between two denominations, commonly used in order books and trading operations.

### Fields

- **`Maker`** (Denom, required): The maker side denomination (base currency in the pair)
- **`Taker`** (Denom, required): The taker side denomination (quote currency in the pair)

### Usage

Token pairs are used to represent trading pairs like:
- XRP/USD
- BTC/ETH
- Custom token pairs

The `Maker` and `Taker` fields define the two sides of the trading pair.

## Use Cases

- Order book management
- Trading pair definitions
- Token identification and validation
- Cross-chain token handling (IBC)
- Price calculations
- Balance tracking

## Notes

- The `Precision` field is optional, allowing flexibility for different token types
- `IsIBC` helps identify tokens that are part of the Inter-Blockchain Communication protocol
- When working with native tokens (like XRP), the issuer may be empty or a special value