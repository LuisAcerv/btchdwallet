package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/LuisAcerv/btchdwallet/crypt"
	"github.com/LuisAcerv/btchdwallet/db"
	"github.com/LuisAcerv/btchdwallet/models"
	pb "github.com/LuisAcerv/btchdwallet/proto/btchdwallet"
	"github.com/LuisAcerv/btchdwallet/wallet"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50055"
)

type server struct {
	pb.UnimplementedWalletServer
}

func (s *server) CreateWallet(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println()
	fmt.Printf("\nCreating new wallet: %s\n", in.Name)
	fmt.Println()

	// 1. Generate phrase
	pub, priv, mnemonic := wallet.CreateWallet()

	// 2. Generate the encryption hash from pin
	encryptionHash := crypt.GenerateHashFromPin(mnemonic)

	// 4. Encrypt the generated KeyPair
	walletString := fmt.Sprintf(`{"name":"%s" ,"pub": "%s", "priv": "%s"}`, in.Name, pub, priv)
	encrypted := crypt.Encrypt(walletString, encryptionHash)

	// 5. Save the to DB
	db.SaveWallet(in.Name, []byte(in.Name), encrypted)

	return &pb.Response{Name: in.Name, PubKey: pub, PrivKey: priv, Mnemonic: mnemonic, Balance: 0}, nil
}

func (s *server) GetWallet(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println()
	fmt.Printf("\nGetting wallet: %s\n", in.Name)
	fmt.Println()

	// 2. Generate the encryption hash from pin
	encryptionHash := crypt.GenerateHashFromPin(in.Mnemonic)

	// 3. GetWallet
	encryptedWallet := db.GetWallet([]byte(in.Name))

	// 4. Decrypt Wallet
	decryptedWallet, _ := crypt.Decrypt(encryptedWallet.Encrypted, encryptionHash)

	// 5. Unmarshal to struct
	wallet := &models.Wallet{}
	json.Unmarshal([]byte(decryptedWallet), wallet)

	return &pb.Response{Name: wallet.Name, PrivKey: wallet.PrivKey, PubKey: wallet.PubKey}, nil
}

func main() {
	db.Initialize()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWalletServer(s, &server{})
	reflection.Register(s)
	fmt.Printf("Service running at port: %s", port)
	fmt.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
