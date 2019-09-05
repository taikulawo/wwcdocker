package subsystems

import (
	"io/ioutil"
	"os"
	"path"
)

type cpuSubsystem struct {
}

func (c *cpuSubsystem) SetLimit(id string, config *ResourceConfig) error {
	cpath, err := GetCgroupPath(c.Name(), id)
	if err != nil {
		return err
	}
	limit := config.CPUShares
	if limit == "" {
		return nil
	}
	if err := ioutil.WriteFile(path.Join(cpath, "cpu.shares"), []byte(limit), 0644); err != nil {
		return err
	}
	return nil
}

// Apply create a hierarchy
// id is container id
func (c *cpuSubsystem) Apply(id string, pid int) error {
	// cpath /sys/fs/cgroup/wwcdocker/1234/
	cpath, err := GetCgroupPath(c.Name(), id)
	if err != nil {
		return err
	}
	return createCgroup(cpath, pid)
}

// Remove delete process from cgroup
func (c *cpuSubsystem) Remove(id string) error {
	cpath, err := GetCgroupPath(c.Name(), id)
	if err != nil {
		return err
	}
	return os.RemoveAll(cpath)
}

func (c *cpuSubsystem) Name() string {
	return "cpu"
}
