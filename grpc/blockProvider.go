package grpc

import (
	"context"
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

func GetBlockchainMetadata() {
	conn, err := grpc.Dial(BlockProviderIP, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := client.GetTransactionInfo(ctx, &Cid{
		Cid: "QmfJs6XRpf34TEK5TbTLFBDbeAbjgpYHHNdKJSrngkKfb7",
	})
	if err != nil {
		log.Fatalf("could not GetTransactionInfo: %v", err)
	}
	log.Printf("GetTransactionInfo reply: %s %d %s %s %s %d",
		r.GetBlockHash(),
		r.GetBlockNumber(),
		r.GetTxHash(),
		r.GetFromAddress(),
		r.GetToAddress(),
		r.GetNonce(),
	)
}
