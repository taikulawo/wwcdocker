package cgroups

type SubSystem interface {
	Set()
}

type ResourceConfig struct {
	MemLimit string
}

type CgroupManager struct{
	path string
	
}

// NewManager return a new manager
func NewManager(path string) *CgroupManager {
	return & CgroupManager {
		path : path,
	}
}

// Add process to special cgroup
func (c *CgroupManager) Add(pid int ) bool{

}

// Remove process from special cgroup
func (c *CgroupManager) Remove(pid int) bool {

}

func (c *CgroupManager) Destroy() bool{

}