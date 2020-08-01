build_protoc:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative proto/btchdwallet/wallet.proto
