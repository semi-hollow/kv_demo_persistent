package raftcore

//type RaftLog struct {
//	mu         sync.RWMutex
//	firstIdx   uint64
//	lastIdx    uint64
//	appliedIdx int64
//	dbEng      storage_eng.KvStore
//	items      []*pb.Entry
//}

//// i think this interface is for, persister? but not applied on raftlog struct.
//type LogOp interface {
//	GetFirst() *pb.Entry
//
//	LogItemCount() int
//
//	EraseBefore(idx int64) []*pb.Entry
//
//	EraseAfter(idx int64) []*pb.Entry
//
//	GetRange(lo, hi int64) []*pb.Entry
//
//	Append(newEnt *pb.Entry)
//
//	GetEntry(idx int64) *pb.Entry
//
//	GetLast() *pb.Entry
//}

// these functions below are just for test, so commented here.

//
// Mem
//

//func MakeMemRaftLog() *RaftLog {
//	empEnt := &pb.Entry{}
//	newItems := []*pb.Entry{}
//	newItems = append(newItems, empEnt)
//	return &RaftLog{items: newItems, firstIdx: INIT_LOG_INDEX, lastIdx: INIT_LOG_INDEX + 1}
//}
//
//func (rfLog *RaftLog) GetMemFirst() *pb.Entry {
//	return rfLog.items[0]
//}
//
//func (rfLog *RaftLog) MemLogItemCount() int {
//	return len(rfLog.items)
//}
//
//func (rfLog *RaftLog) EraseMemBefore(idx int64) []*pb.Entry {
//	return rfLog.items[idx:]
//}
//
//func (rfLog *RaftLog) EraseMemAfter(idx int64) []*pb.Entry {
//	return rfLog.items[:idx]
//}
//
//func (rfLog *RaftLog) GetMemRange(lo, hi int64) []*pb.Entry {
//	return rfLog.items[lo:hi]
//}
//
//func (rfLog *RaftLog) MemAppend(newEnt *pb.Entry) {
//	rfLog.items = append(rfLog.items, newEnt)
//}
//
//func (rfLog *RaftLog) GetMemEntry(idx int64) *pb.Entry {
//	return rfLog.items[idx]
//}
//
//func (rfLog *RaftLog) GetMemLast() *pb.Entry {
//	return rfLog.items[len(rfLog.items)-1]
//}
