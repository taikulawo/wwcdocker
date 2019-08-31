package main

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestMain(t *testing.T) {
	var i *int32
	var oldV, newV, origin int32
	oldV = 3
	newV = 2
	origin = 3
	i = &origin
	a := atomic.CompareAndSwapInt32(i, oldV, newV)
	fmt.Println(a)
}
