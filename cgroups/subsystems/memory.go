package subsystems

import (
	"io/ioutil"
	"os"
	"path"
)

type memorySubsystem struct{}

// Apply create new cgroup
func (m *memorySubsystem) Apply(id string, pid int) error {
	cpath, err := GetCgroupPath(m.Name(), id)
	if err != nil {
		return err
	}
	return createCgroup(cpath, pid)
}

func (m *memorySubsystem) Remove(id string) error {
	cpath, err := GetCgroupPath(m.Name(), id)
	if err != nil {
		return err
	}
	return os.RemoveAll(cpath)
}

func (m *memorySubsystem) SetLimit(id string, config *ResourceConfig) error {
	cpath, err := GetCgroupPath(m.Name(), id)
	if err != nil {
		return err
	}
	limitInBytes := config.MemLimit
	if limitInBytes == "" {
		return nil
	}
	if err := ioutil.WriteFile(path.Join(cpath, "memory.limit_in_bytes"), []byte(limitInBytes),0644); err != nil {
		return err
	}
	return nil
}

func (m *memorySubsystem) Name() string {
	return "memory"
}
