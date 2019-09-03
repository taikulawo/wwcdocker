package subsystems

import (
	"io/ioutil"
	"path"
	"strconv"
)

type CPUSubsystem struct {
}

func (c *CPUSubsystem) SetLimit(id string, config *ResourceConfig) error {
	cpath, err := GetCgroupPath(c.Name(), id)
	if err != nil {
		return err
	}
	limit := config.CPUSet
	if err := ioutil.WriteFile(path.Join(cpath,"cpu.shares"),[]byte(limit),0644); err != nil {
		return err
	}
	return nil
}

// Apply create a hierarchy
// id is container id
func (c *CPUSubsystem) Apply(id string, pid int) error {
	// cpath /sys/fs/cgroup/wwcdocker/1234/
	cpath, err := GetCgroupPath(c.Name(), id)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path.Join(cpath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return err
	}
	return nil
}

// Remove delete process from cgroup
func (c *CPUSubsystem) Remove(id string) error {
	return nil
}

func (c *CPUSubsystem) Name() string {
	return "cpu"
}
