package common

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
)

func GetRandomNumber(l int) (b []byte) {
	b = make([]byte, l, l)
	rand.Read(b)
	return
}

func NameExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return true, nil
	}
	return false, err
}

func ReadFromFd(fd uintptr) (string, error) {
	f := os.NewFile(fd, "cmdInit")
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("Failed to read from fd %d, error: %v", fd, err)
	}
	return string(b), nil
}
