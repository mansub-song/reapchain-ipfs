/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/mansub1029/reapchain-ipfs/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "147.46.245.72:50051", "the address to connect to")
)

type TxInfo struct {
	BlockHash   string
	BlockNumber uint32
	TxHash      string
	FromAddress string
	ToAddress   string
	Nonce       uint32
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	testSayTransactionInfo(ctx, client)

	testGetTransactionInfo(ctx, client)

}
func testGetTransactionInfo(ctx context.Context, client pb.GreeterClient) {
	r, err := client.GetTransactionInfo(ctx, &pb.Cid{
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

func testSayTransactionInfo(ctx context.Context, client pb.GreeterClient) {
	r, err := client.SayTransactionInfo(ctx, &pb.TransactionRequest{
		BlockHash:   "0x9bb5a652cbbdb8f8b7e1cbbdb21264d7fa93983aada84d66272ab45233e740cf",
		BlockNumber: 91541744,
		TxHash:      "0xec166965eb8b5374f9ad52d1fa541de5e318d825242ed024d4a79d1b73e9fd59",
		FromAddress: "0xb2936f054560409973fffaa7d5fdae9e5c8b628e",
		ToAddress:   "0x3c817b136bad58d35c81bd1981b0151b7e07f21b",
		Nonce:       486,
	})
	if err != nil {
		log.Fatalf("could not SayTransactionInfo: %v", err)
	}
	log.Printf("SayTransactionInfo reply: %s", r.GetMessage())
}
