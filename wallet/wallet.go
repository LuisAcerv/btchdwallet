package wallet

import (
	"github.com/wemeetagain/go-hdwallet"
)

// CreateWallet is in charge of creating a new root wallet
func CreateWallet() (string, string) {
	// Generate a random 256 bit seed
	seed, _ := hdwallet.GenSeed(256)

	// Create a master private key
	masterprv := hdwallet.MasterKey(seed)

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	return masterpub.String(), masterprv.String()
}
