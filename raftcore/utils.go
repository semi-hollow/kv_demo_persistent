package raftcore

import (
	"fmt"
	"math/rand"
	"time"
)

func RandIntRange(min int, max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(max-min) + int(min)
}

func MakeAnRandomElectionTimeout(base int) int {
	return RandIntRange(base, 2*base)
}

func PrintDebugLog(msg string) {
	fmt.Printf("%s %s \n", time.Now().Format("2006-01-02 15:04:05"), msg)
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
