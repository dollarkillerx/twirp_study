package server

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	add := time.Now().Add(time.Second * 3).Unix()
	time.Sleep(time.Second * 4)
	now := time.Now().Unix()
	if add > now {
		fmt.Println("Success")
	}
	fmt.Println("Add: ",add)
	fmt.Println("Now: ",now)
}
