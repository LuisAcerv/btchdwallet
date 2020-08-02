package main

import (
	"context"
	"fmt"
	"log"
	"net"

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
	fmt.Println("\nCreating new wallet")

	addr, pub, priv, mnemonic := wallet.CreateWallet()

	return &pb.Response{Address: addr, PubKey: pub, PrivKey: priv, Mnemonic: mnemonic, Balance: 0}, nil
}

func (s *server) GetWallet(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println()
	fmt.Println("\nGetting wallet data")

	address, pub, priv := wallet.DecodeWallet(in.Mnemonic)

	return &pb.Response{Address: address, PrivKey: priv, PubKey: pub}, nil
}

func (s *server) GetBalance(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println()
	fmt.Println("\nGetting Balance data")

	balance, totalReceived, totalSent, unconfirmedBalance := wallet.GetBalance(in.Address)

	return &pb.Response{Address: in.Address, Balance: int64(balance), TotalReceived: int64(totalReceived), TotalSent: int64(totalSent), UnconfirmedBalance: int64(unconfirmedBalance)}, nil
}

func main() {
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
