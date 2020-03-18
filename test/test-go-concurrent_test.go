package test

import (
	"fmt"
	"testing"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func TestGo(t *testing.T) {
	// go关键字直接开启线程
	go say("world")
	say("hello")
}
