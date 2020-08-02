package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/LuisAcerv/btchdwallet/proto/btchdwallet"
	"google.golang.org/grpc"
)

var address = "localhost:50055"

func createWallet() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Body read error, %v", err)
		fmt.Println(err, "Error!")
		return
	}

	defer conn.Close()
	c := pb.NewWalletClient(conn)
	clientDeadline := time.Now().Add(time.Duration(10000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	r, err := c.CreateWallet(ctx, &pb.Request{})
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	log.Printf("\n\nNew Wallet >>\n\n > Public Key: %s\n\n > Private Key: %s\n\n > Mnemonic: %s", r.PubKey, r.PrivKey, r.Mnemonic)
}

func getWallet(mnemonic string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Body read error, %v", err)
		//w.WriteHeader(500) // Return 500 Internal Server Error.

		fmt.Println(err, "Error!")
		return
	}

	defer conn.Close()
	c := pb.NewWalletClient(conn)
	clientDeadline := time.Now().Add(time.Duration(10000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	r, err := c.GetWallet(ctx, &pb.Request{Mnemonic: mnemonic})
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	log.Printf("\n\nWallet >>\n\n > Address: %s\n\n > Public Key: %s\n\n > Private Key: %s\n\n", r.Address, r.PubKey, r.PrivKey)
}

func getBalance(addr string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Body read error, %v", err)
		//w.WriteHeader(500) // Return 500 Internal Server Error.

		fmt.Println(err, "Error!")
		return
	}

	defer conn.Close()
	c := pb.NewWalletClient(conn)
	clientDeadline := time.Now().Add(time.Duration(10000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	r, err := c.GetBalance(ctx, &pb.Request{Address: addr})
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	log.Printf("\n\nBalance >>\n\n > Address: %s\n\n > Balance: %v\n\n > Total Received: %v\n\n > Total Sent: %v\n\n > Unconfirmed: %v\n\n ", r.Address, r.Balance, r.TotalReceived, r.TotalSent, r.UnconfirmedBalance)
}

func main() {
	method := flag.String("m", "default", "Method to be executed")
	mnemonic := flag.String("mne", "default", "Encryption Pin")
	address := flag.String("addr", "default", "Bitcoin Address")

	flag.Parse()

	switch *method {
	case "create-wallet":
		createWallet()
		return

	case "get-wallet":
		getWallet(*mnemonic)
		return

	case "get-balance":
		getBalance(*address)
		return
	}

}
