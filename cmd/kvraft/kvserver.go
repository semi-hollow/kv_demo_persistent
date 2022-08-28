package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	kvsvr "raft_grpc_leveldb/kvserver"
	pb "raft_grpc_leveldb/raftpb"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: server [nodeId]")
		return
	}
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	nodeIdStr := os.Args[1]
	nodeId, err := strconv.Atoi(nodeIdStr)
	if err != nil {
		panic(err)
	}

	kvServer := kvsvr.MakeKvServer(nodeId)
	lis, err := net.Listen("tcp", kvsvr.PeersMap[nodeId])
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	fmt.Printf("server listen on: %s \n", kvsvr.PeersMap[nodeId])
	s := grpc.NewServer()
	pb.RegisterRaftServiceServer(s, kvServer)

	sigChan := make(chan os.Signal)

	signal.Notify(sigChan)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		kvServer.Rf.CloseEndsConn()
		os.Exit(-1)
	}()

	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
