package storage_eng

//
// If you want to contribute a new engine implementation, you need to implement these interfaces
//
type KvStore interface {
	Put(string, string) error
	Get(string) (string, error)
	Delete(string) error
	DumpPrefixKey(string) (map[string]string, error)
	PutBytesKv(k []byte, v []byte) error
	DeleteBytesK(k []byte) error
	GetBytesValue(k []byte) ([]byte, error)
	SeekPrefixLast(prefix []byte) ([]byte, []byte, error)
	SeekPrefixFirst(prefix string) ([]byte, []byte, error)
	DelPrefixKeys(prefix string) error
	SeekPrefixKeyIdMax(prefix []byte) (uint64, error)
	FlushDB()
}
