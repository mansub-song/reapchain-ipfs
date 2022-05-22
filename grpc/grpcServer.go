package grpc

import (
	"context"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
)


var (
	port =50051
	tx TxInfo
)

type TxInfo struct {
	BlockHash string
	BlockNumber uint32
	TxHash string
	FromAddress string
	ToAddress string
	Nonce uint32
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	fmt.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}
// func (s *server) SayHelloAgain(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
//         return &HelloReply{Message: "Hello again " + in.GetName()}, nil
// }

func (s *server) SayTransactionInfo(ctx context.Context, in *TransactionInfo) (*HelloReply, error) {
		tx := &TxInfo {
			BlockHash: in.GetBlockHash(),
			BlockNumber: in.GetBlockNumber(),
			TxHash: in.GetTxHash(),
			FromAddress:in.GetFromAddress(),
			ToAddress:in.GetToAddress(),
			Nonce:in.GetNonce(),

		}
		// fmt.Printf("Received: %v", in)
		fmt.Printf("txInfo: %+v",*tx)
		fmt.Println(tx.BlockHash)
        return &HelloReply{Message: "Hello again " + in.GetBlockHash() + "  " + in.GetFromAddress()}, nil
}

func ServerInit() {
	flag.Parse()
	// lis, err := net.Listen("tcp", fmt.Sprintf("147.46.240.229:%d", port))
	lis, err := net.Listen("tcp", fmt.Sprintf("147.46.240.229:%d", port))
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
