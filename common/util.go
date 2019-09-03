package common

import "crypto/rand"
import "os"

func GetRandomNumber(l int) (b []byte) {
	b = make([]byte, l,l)
	rand.Read(b)
	return 
}

func NameExists(path string ) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return true, nil
	}
	return false, err
}

