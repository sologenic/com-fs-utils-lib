package metadata

import (
	"log"
	"os"
	"strings"
)

// ParseNetworkFromEnv parses the NETWORK environment variable to metadata.Network enum
// Defaults to MAINNET if the environment variable is not set
func ParseNetworkFromEnv() Network {
	networkStr := strings.ToUpper(os.Getenv("NETWORK"))

	switch networkStr {
	case "MAINNET":
		return Network_MAINNET
	case "TESTNET":
		return Network_TESTNET
	case "DEVNET":
		return Network_DEVNET
	default:
		log.Printf("Invalid or missing NETWORK env var '%s', defaulting to MAINNET", networkStr)
		return Network_MAINNET
	}
}
