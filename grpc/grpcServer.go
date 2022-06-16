package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = 50051

var (
	tx TxInfo
)

type TxInfo struct {
	BlockHash   string
	BlockNumber uint32
	TxHash      string
	FromAddress string
	ToAddress   string
	Nonce       uint32
}

// server is used to implement reapGRPC.GreeterServer.
type server struct {
	UnimplementedGreeterServer
}

// Get preferred outbound ip of this machine
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func (s *server) SayTransactionInfo(ctx context.Context, in *TransactionRequest) (*TransactionReply, error) {
	tx := &TxInfo{
		BlockHash:   in.GetBlockHash(),
		BlockNumber: in.GetBlockNumber(),
		TxHash:      in.GetTxHash(),
		FromAddress: in.GetFromAddress(),
		ToAddress:   in.GetToAddress(),
		Nonce:       in.GetNonce(),
	}
	// fmt.Printf("Received: %v", in)
	fmt.Printf("txInfo: %+v", *tx)
	fmt.Println(tx.BlockHash)
	return &TransactionReply{Message: in.GetBlockHash() + "  " + in.GetFromAddress()}, nil
}

func (s *server) GetTransactionInfo(ctx context.Context, in *Cid) (*CidReply, error) {
	cid := in.GetCid()
	fmt.Println("GetTransactionInfo cid:", cid)
	tx := &TxInfo{
		BlockHash:   "0x9bb5a652cbbdb8f8b7e1cbbdb21264d7fa93983aada84d66272ab45233e740cf",
		BlockNumber: 91541744,
		TxHash:      "0xec166965eb8b5374f9ad52d1fa541de5e318d825242ed024d4a79d1b73e9fd59",
		FromAddress: "0xb2936f054560409973fffaa7d5fdae9e5c8b628e",
		ToAddress:   "0x3c817b136bad58d35c81bd1981b0151b7e07f21b",
		Nonce:       486,
	}
	return &CidReply{
		BlockHash:   tx.BlockHash,
		BlockNumber: tx.BlockNumber,
		TxHash:      tx.TxHash,
		FromAddress: tx.FromAddress,
		ToAddress:   tx.ToAddress,
		Nonce:       tx.Nonce,
	}, nil
}

func ServerInit() {
	localIP := getOutboundIP().String()
	lis, err := net.Listen("tcp", fmt.Sprintf(localIP+":%d", port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &server{})
	fmt.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
