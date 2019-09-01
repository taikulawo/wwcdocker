package subsystems


var (
	Subsystems = []Subsystem {
		&CPUSubsystem{},
	}
)

type Subsystem interface{
	// Name returns name of cgroup
	Name() string
	// Set sets resource limits
	SetLimit(c *ResourceConfig) error
	// Apply create a cgroup
	Apply(cname string, pid int) error
	// Remove delete process from this cgroup
	Remove(cname string) error
}

type ResourceConfig struct {
	CPUSet string
	MemLimit string
}
