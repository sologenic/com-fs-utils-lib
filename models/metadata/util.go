package metadata

import (
	"os"
	"strings"

	"github.com/sologenic/com-fs-utils-internal-lib/go/logger"
)

var network Network

func init() {
	value := strings.ToUpper(os.Getenv("NETWORK"))
	if value != "MAINNET" && value != "TESTNET" {
		logger.Fatalf("NETWORK must be 'MAINNET' or 'TESTNET'")
	}
	network = Network(Network_value[value])
}

// GetNetwork Returns metadata.Network based on the NETWORK environment variable
func GetNetwork() Network {
	return network
}
