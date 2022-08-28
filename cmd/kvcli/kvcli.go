package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"raft_grpc_leveldb/common"
	"raft_grpc_leveldb/raftcore"
	pb "raft_grpc_leveldb/raftpb"
	"strconv"
	"time"
)

type KvClient struct {
	rpcCli    *raftcore.RaftClientEnd
	leaderId  int64
	clientId  int64
	commandId int64
}

func (kvCli *KvClient) Close() {
	kvCli.rpcCli.CloseAllConn()
}

func nrand() int64 {
	max := big.NewInt(int64(1) << 62)
	bigx, _ := rand.Int(rand.Reader, max)
	return bigx.Int64()
}

func MakeKvClient(targetId int, targetAddr string) *KvClient {
	cli := raftcore.MakeRaftClientEnd(targetAddr, uint64(targetId))
	return &KvClient{
		rpcCli:    cli,
		leaderId:  0,
		clientId:  nrand(),
		commandId: 0,
	}
}

func (kvCli *KvClient) Get(key string) string {
	cmdReq := &pb.CommandRequest{
		Key:      key,
		OpType:   pb.OpType_OpGet,
		ClientId: kvCli.clientId,
	}
	//fmt.Println("putbefore do command")
	resp, err := (*kvCli.rpcCli.GetRaftServiceCli()).DoCommand(context.Background(), cmdReq)
	if err != nil {
		return "err"
	}
	fmt.Println("key value is:", resp.Value)
	//fmt.Println("leader_id is:", resp.LeaderId)
	//fmt.Println("error_code is:", resp.ErrCode)

	return resp.Value
}

func (kvCli *KvClient) Put(key, value string) string {
	cmdReq := &pb.CommandRequest{
		Key:      key,
		Value:    value,
		ClientId: kvCli.clientId,
		OpType:   pb.OpType_OpPut,
	}
	_, err := (*kvCli.rpcCli.GetRaftServiceCli()).DoCommand(context.Background(), cmdReq)
	if err != nil {
		return "err"
	}
	//return "put ok, now the key value is updated to:"
	return "ok"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: kvcli [serveraddr] [count] [op] [options]")
		return
	}
	sigs := make(chan os.Signal, 1)

	kvCli := MakeKvClient(99, os.Args[1])

	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	op := os.Args[3]

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		kvCli.rpcCli.CloseAllConn()
		os.Exit(-1)
	}()

	keys := make([]string, count)
	vals := make([]string, count)

	for i := 0; i < count; i++ {
		rndK := common.RandStringRunes(8)
		rndV := common.RandStringRunes(8)
		keys[i] = rndK
		vals[i] = rndV
	}

	startTs := time.Now()
	for i := 0; i < count; i++ {
		//fmt.Println(kvCli.Put(keys[i], vals[i]))
		//fmt.Println(kvCli.Get(keys[i]))
		kvCli.Put(keys[i], vals[i])
		kvCli.Get(keys[i])
	}
	elapsed := time.Since(startTs).Seconds()
	fmt.Printf("total cost %f s\n", elapsed)

	switch op {
	case "get":
		//fmt.Println("23333get")
		kvCli.Get(os.Args[4])

	case "put":
		//fmt.Println("23333put")
		fmt.Println(kvCli.Put(os.Args[4], os.Args[5]))
		kvCli.Get("testkey")
	}

	//fmt.Println("run test get value -> " + kvCli.Get("testkey"))
	//fmt.Println("run test get value -> " + kvCli.Get(keys[0]))
}
