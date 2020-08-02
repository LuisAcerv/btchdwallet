package wallet

import (
	"fmt"

	pb "github.com/LuisAcerv/btchdwallet/proto/btchdwallet"

	"github.com/LuisAcerv/btchdwallet/config"
	"github.com/LuisAcerv/btchdwallet/crypt"
	"github.com/blockcypher/gobcy"
	"github.com/brianium/mnemonic"
	"github.com/wemeetagain/go-hdwallet"
)

var conf = config.ParseConfig()

// CreateWallet is in charge of creating a new root wallet
func CreateWallet() *pb.Response {
	// Generate a random 256 bit seed
	seed := crypt.CreateHash()
	mnemonic, _ := mnemonic.New([]byte(seed), mnemonic.English)

	// Create a master private key
	masterprv := hdwallet.MasterKey([]byte(mnemonic.Sentence()))

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	// Get your address
	address := masterpub.Address()

	return &pb.Response{Address: address, PubKey: masterpub.String(), PrivKey: masterprv.String(), Mnemonic: mnemonic.Sentence()}
}

// DecodeWallet is in charge of decoding wallet from mnemonic
func DecodeWallet(mnemonic string) *pb.Response {
	// Get private key from mnemonic
	masterprv := hdwallet.MasterKey([]byte(mnemonic))

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	// Get your address
	address := masterpub.Address()

	return &pb.Response{Address: address, PubKey: masterpub.String(), PrivKey: masterprv.String()}
}

// GetBalance is in charge of returning the given address balance
func GetBalance(address string) *pb.Response {
	btc := gobcy.API{conf.Blockcypher.Token, "btc", "main"}
	addr, err := btc.GetAddrBal(address, nil)
	if err != nil {
		fmt.Println(err)
	}

	balance := addr.Balance
	totalReceived := addr.TotalReceived
	totalSent := addr.TotalSent
	unconfirmedBalance := addr.UnconfirmedBalance

	return &pb.Response{Address: address, Balance: int64(balance), TotalReceived: int64(totalReceived), TotalSent: int64(totalSent), UnconfirmedBalance: int64(unconfirmedBalance)}

}
