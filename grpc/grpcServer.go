package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc"
)

const port = 50051
const metadataPath = "/home/mssong/.ipfs/"

var (
	tx TxInfo
)

type TxInfo struct {
	BlockHash      string
	BlockNumber    uint32
	TxHash         string
	FromAddress    string
	ToAddress      string
	Nonce          uint32
	Cid            string
	BlockchainName string
}

// server is used to implement reapGRPC.GreeterServer.
type server struct {
	UnimplementedGreeterServer
}

// Get preferred outbound ip of this machine
func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func AddTransactionInfo(key []byte, value []byte) {
	db, err := leveldb.OpenFile(metadataPath, nil) //저장할 path
	if err != nil {
		log.Fatal("Failed leveldb open!")
	}
	err = db.Put(key, value, nil)
	// newValue, err := db.Get(key, nil)
	defer db.Close()
}

//client 코드의 testSayTransactionInfo() 함수에 대한 reply 해주는 함수
func (s *server) SayTransactionInfo(ctx context.Context, in *TransactionRequest) (*TransactionReply, error) {
	tx := &TxInfo{
		BlockHash:      in.GetBlockHash(),
		BlockNumber:    in.GetBlockNumber(),
		TxHash:         in.GetTxHash(),
		FromAddress:    in.GetFromAddress(),
		ToAddress:      in.GetToAddress(),
		Nonce:          in.GetNonce(),
		Cid:            in.GetCid(),
		BlockchainName: in.GetBlockchainName(),
	}
	// fmt.Printf("Received: %v", in)
	fmt.Printf("txInfo: %+v", *tx)
	fmt.Println(tx.BlockHash)

	blockchainName := in.GetBlockchainName()
	key := []byte(blockchainName + in.GetCid())
	fmt.Println("key:", blockchainName+in.GetCid())
	fmt.Printf("key bytes: %x\n", key)
	value, err := json.Marshal(tx)
	// fmt.Println("value:", value)
	if err != nil {
		log.Fatal("Failed marshaling!")
	}
	AddTransactionInfo(key, value)
	// var newTx *TxInfo
	// err = json.Unmarshal(newValue, newTx)
	// fmt.Println("newValue:", newValue)
	// fmt.Printf("newTx: %+v", *newTx)

	return &TransactionReply{Message: in.GetBlockHash() + "  " + in.GetFromAddress()}, nil
}

//client 코드에서 testGetTransactionInfo() 함수에 대한 reply를 해주는 함수
func (s *server) GetTransactionInfo(ctx context.Context, in *MetadataKey) (*MetadataValue, error) {
	cid := in.GetCid()
	blockchainName := in.GetBlockchainName()
	key := []byte(blockchainName + cid)
	// fmt.Println("GetTransactionInfo cid:", cid)
	db, err := leveldb.OpenFile(metadataPath, nil) //저장할 path
	if err != nil {
		panic(err)
	}

	fmt.Println("key:", blockchainName+cid)
	fmt.Printf("key bytes: %x\n", key)

	value, err := db.Get(key, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("leveldb value:", string(value))

	tx := &TxInfo{
		BlockHash:      "0x9bb5a652cbbdb8f8b7e1cbbdb21264d7fa93983aada84d66272ab45233e740cf",
		BlockNumber:    91541744,
		TxHash:         "0xec166965eb8b5374f9ad52d1fa541de5e318d825242ed024d4a79d1b73e9fd59",
		FromAddress:    "0xb2936f054560409973fffaa7d5fdae9e5c8b628e",
		ToAddress:      "0x3c817b136bad58d35c81bd1981b0151b7e07f21b",
		Nonce:          486,
		Cid:            "QmcxuB3wB9fSPyimMLfqD9CxLSPQ3NgM7Es97aLrGT1nwe",
		BlockchainName: "ETH",
	}
	return &MetadataValue{
		BlockHash:      tx.BlockHash,
		BlockNumber:    tx.BlockNumber,
		TxHash:         tx.TxHash,
		FromAddress:    tx.FromAddress,
		ToAddress:      tx.ToAddress,
		Nonce:          tx.Nonce,
		Cid:            tx.Cid,
		BlockchainName: tx.BlockchainName,
	}, nil
}

func ServerInit() {
	localIP := GetLocalIP().String()
	lis, err := net.Listen("tcp", fmt.Sprintf(localIP+":%d", port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &server{})
	fmt.Printf("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
