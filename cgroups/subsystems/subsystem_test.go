package subsystems

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"testing"
)

func TestSubsystems(t *testing.T) {
	p := "/sys/fs/cgroup/memory/test1"
	f := "tasks"

	err := os.MkdirAll(p, 0644)

	if err != nil {
		t.Log(err.Error())
	}

	pid := os.Getpid()

	if err := ioutil.WriteFile(path.Join(p, f), []byte(strconv.Itoa(pid)), 0644); err != nil {
		t.Failed()
	}
}
