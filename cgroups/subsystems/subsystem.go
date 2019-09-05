package subsystems


var (
	Subsystems = []subsystem {
		&cpuSubsystem{},
		&memorySubsystem{},
	}
)

type subsystem interface{
	// Name returns name of cgroup
	Name() string
	// Set sets resource limits
	SetLimit(id string, c *ResourceConfig) error
	// Apply create a cgroup
	Apply(id string, pid int) error
	// Remove delete process from this cgroup
	Remove(id string) error
}

type ResourceConfig struct {
	CPUSet string
	CPUShares string
	MemLimit string
}
