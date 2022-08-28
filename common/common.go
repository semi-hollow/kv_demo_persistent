package common

import (
	"hash/crc32"
	"math/rand"
	"time"
)

const NBuckets = 10

const UN_UNSED_TID = 999

const (
	ErrDefault int64 = iota
	ErrCodeNoErr
	ErrCodeWrongLeader
	ErrCodeExecTimeout
)

func Key2BucketID(key string) int {
	return CRC32KeyHash(key, NBuckets)
}

func CRC32KeyHash(k string, base int) int {
	bucketId := 0
	crc32q := crc32.MakeTable(0xD5828281)
	sum := crc32.Checksum([]byte(k), crc32q)
	bucketId = int(sum) % NBuckets
	return bucketId
}

func Int64ArrToIntArr(in []int64) []int {
	out := []int{}
	for _, item := range in {
		out = append(out, int(item))
	}
	return out
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
