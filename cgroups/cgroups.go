package cgroups

import sub "github.com/iamwwc/wwcdocker/cgroups/subsystems"

type SubSystem interface {
	Set()
}

type CgroupManager struct {
	path string
}

// NewManager return a new manager
func NewManager(path string) *CgroupManager {
	return &CgroupManager{
		path: path,
	}
}

// Add process to special cgroup
func (c *CgroupManager) Add(pid int) bool {
	// create new cgroup for this process

	
}

func (c *CgroupManager) SetResourceLimit(config *sub.ResourceConfig) error {
	for _, sys := range sub.Subsystems {
		sys.SetLimit(config)
	}
	return nil
}

// Remove process from special cgroup
func (c *CgroupManager) Remove(pid int) bool {

}

func (c *CgroupManager) Destroy() bool {

}
