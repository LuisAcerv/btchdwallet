package wallet

import (
	"fmt"

	"github.com/LuisAcerv/btchdwallet/config"
	"github.com/LuisAcerv/btchdwallet/crypt"
	"github.com/blockcypher/gobcy"
	"github.com/brianium/mnemonic"
	"github.com/wemeetagain/go-hdwallet"
)

var conf = config.ParseConfig()

// CreateWallet is in charge of creating a new root wallet
func CreateWallet() (string, string, string, string) {
	// Generate a random 256 bit seed
	seed := crypt.CreateHash()
	mnemonic, _ := mnemonic.New([]byte(seed), mnemonic.English)

	// Create a master private key
	masterprv := hdwallet.MasterKey([]byte(mnemonic.Sentence()))

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	// Get your address
	address := masterpub.Address()

	return address, masterpub.String(), masterprv.String(), mnemonic.Sentence()
}

// DecodeWallet is in charge of decoding wallet from mnemonic
func DecodeWallet(mnemonic string) (string, string, string) {
	// Get private key from mnemonic
	masterprv := hdwallet.MasterKey([]byte(mnemonic))

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	// Get your address
	address := masterpub.Address()

	return address, masterpub.String(), masterprv.String()
}

// GetBalance is in charge of returning the given address balance
func GetBalance(address string) (int, int, int, int) {
	btc := gobcy.API{conf.Blockcypher.Token, "btc", "main"}
	addr, err := btc.GetAddrBal(address, nil)
	if err != nil {
		fmt.Println(err)
	}

	balance := addr.Balance
	totalReceived := addr.TotalReceived
	totalSent := addr.TotalSent
	unconfirmedBalance := addr.UnconfirmedBalance

	return balance, totalReceived, totalSent, unconfirmedBalance
}
