package common

import (
	"testing"
)

func TestFile(t *testing.T) {
	path := "/proc/1/fd"
	_ = GetAllFdsOfProcess(path)
}
