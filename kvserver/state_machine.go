package kvserver

import "errors"

type StateMachine interface {
	Get(key string) (string, error)
	Put(key, value string) error
	Append(key, value string) error
}

type MemKV struct {
	KV map[string]string
}

func NewMemKV() *MemKV {
	return &MemKV{make(map[string]string)}
}

func (memKv *MemKV) Get(key string) (string, error) {
	if v, ok := memKv.KV[key]; ok {
		return v, nil
	}
	return "", errors.New("KeyNotFound")
}

func (memKv *MemKV) Put(key, value string) error {
	memKv.KV[key] = value
	return nil
}

func (memKv *MemKV) Append(key, value string) error {
	memKv.KV[key] += value
	return nil
}
