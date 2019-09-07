package common

import (
	"fmt"
	"io/ioutil"
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/google/uuid"
)

func GetRandomNumber() string {
	s, err := uuid.NewRandom()
	if err != nil {
		log.Warnf("UUID error. %v",err)
	}
	return s.String()
}

func NameExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	log.Error(err)
	return false
}

func ReadFromFd(fd uintptr) (string, error) {
	f := os.NewFile(fd, "cmdInit")
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("Failed to read from fd %d, error: %v", fd, err)
	}
	return string(b), nil
}


const (
	quiet = "-q"
	directory = "-P"
	verbose = "-v"
)

// Use wget download images
func DownloadFromUrl(url string, savedDir string) {
	if _, err := Exec("wget",verbose,directory,savedDir); err != nil {
		log.Error(err)
	}
}