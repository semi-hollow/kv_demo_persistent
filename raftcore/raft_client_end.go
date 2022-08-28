package raftcore

import (
	"fmt"
	"google.golang.org/grpc"
	raftpb "raft_grpc_leveldb/raftpb"
)

type RaftClientEnd struct {
	id             uint64
	addr           string
	conns          []*grpc.ClientConn
	raftServiceCli *raftpb.RaftServiceClient
}

//type RaftClientEndwithResolver struct {
//	id             uint64
//	addr           string
//	conns          []*grpc.ClientConn
//	raftServiceCli []*raftpb.RaftServiceClient
//}

func (rfEnd *RaftClientEnd) Id() uint64 {
	return rfEnd.id
}

func (rfEnd *RaftClientEnd) GetRaftServiceCli() *raftpb.RaftServiceClient {
	return rfEnd.raftServiceCli
}

func MakeRaftClientEnd(addr string, id uint64) *RaftClientEnd {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	//fmt.Printf("faild to connect: %v", err)

	conns := []*grpc.ClientConn{}
	conns = append(conns, conn)
	rpcClient := raftpb.NewRaftServiceClient(conn)
	return &RaftClientEnd{
		id:             id,
		addr:           addr,
		conns:          conns,
		raftServiceCli: &rpcClient,
	}
}

//func MakeRaftClientEndwithResolver(addr string, id uint64) *RaftClientEnd {
//	addrs := strings.Split(addr, ",")
//	//fmt.Printf("addrs:", addrs)
//	conns := []*grpc.ClientConn{}
//	var conn *grpc.ClientConn
//	var err error
//
//	rpcCli := &RaftClientEndwithResolver{
//		id:    id,
//		addr:  addr,
//		conns: conns,
//	}
//
//	for _, v := range addrs {
//		conn, err = grpc.Dial(v, grpc.WithInsecure())
//		if err != nil {
//			fmt.Printf("failed to connect: %v", err)
//		}
//		conns = append(conns, conn)
//		cli := raftpb.NewRaftServiceClient(conn)
//		rpcCli.raftServiceCli = append(rpcCli.raftServiceCli, cli)
//	}
//
//	//rpcClient := raftpb.NewRaftServiceClient(conn)
//	return rpcCli
//}

func (rfEnd *RaftClientEnd) CloseAllConn() {
	// PrintDebugLog(fmt.Sprintf("%s close rpc connect", rfEnd.addr))
	for _, conn := range rfEnd.conns {
		conn.Close()
	}
}
