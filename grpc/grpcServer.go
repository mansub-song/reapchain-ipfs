package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port =50051
var (
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
        return &TransactionReply{Message: "Hello again " + in.GetBlockHash() + "  " + in.GetFromAddress()}, nil
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
