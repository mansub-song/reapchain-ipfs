package grpc

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	BlockProvider   string // ex) ip4/147.46.240.229/tcp/4001
	BlockProviderIP string // ex) 147.46.240.229
	RootCid         string // ex) QmRwEqSE9VrAThDccfNPePL14z34zjm7RX7MH4xWsdn96U
)

func GetBlockchainMetadata(blockProviderIP string, cid string) {
	blockchainName := "ETH"
	conn, err := grpc.Dial(blockProviderIP+":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := client.GetTransactionInfo(ctx, &MetadataKey{
		Cid: cid, BlockchainName: blockchainName,
	}) //cid : QmcxuB3wB9fSPyimMLfqD9CxLSPQ3NgM7Es97aLrGT1nwe
	if err != nil {
		log.Fatalf("could not GetTransactionInfo: %v", err)
	}

	tx := &TxInfo{
		BlockHash:      r.GetBlockHash(),
		BlockNumber:    r.GetBlockNumber(),
		TxHash:         r.GetTxHash(),
		FromAddress:    r.GetFromAddress(),
		ToAddress:      r.GetToAddress(),
		Nonce:          r.GetNonce(),
		Cid:            r.GetCid(),
		BlockchainName: r.GetBlockchainName(),
	}

	key := []byte(blockchainName + cid)
	value, err := json.Marshal(tx)
	AddTransactionInfo(key, value)
	log.Printf("GetTransactionInfo reply: %s %d %s %s %s %d",
		r.GetBlockHash(),
		r.GetBlockNumber(),
		r.GetTxHash(),
		r.GetFromAddress(),
		r.GetToAddress(),
		r.GetNonce(),
	)
}
