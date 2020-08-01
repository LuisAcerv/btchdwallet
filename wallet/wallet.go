package wallet

import (
	"github.com/LuisAcerv/btchdwallet/crypt"
	"github.com/brianium/mnemonic"
	"github.com/wemeetagain/go-hdwallet"
)

// CreateWallet is in charge of creating a new root wallet
func CreateWallet() (string, string, string) {
	// Generate a random 256 bit seed
	//seed, _ := hdwallet.GenSeed(256)
	seed := crypt.CreateHash()
	mnemonic, _ := mnemonic.New([]byte(seed), mnemonic.English)

	// Create a master private key
	masterprv := hdwallet.MasterKey([]byte(mnemonic.Sentence()))

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	return masterpub.String(), masterprv.String(), mnemonic.Sentence()
}
