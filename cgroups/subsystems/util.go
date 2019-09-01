package subsystems

import (
	"bufio"
	"os"
	"strings"
)

// Find Cgroup path according to sub system type
func FindCgroupMountPoint(subsystem string) string {
	file, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	// 41 32 0:36 / /sys/fs/cgroup/memory rw,nosuid,nodev,noexec,relatime shared:19 - cgroup cgroup rw,memory
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(text, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	if err := scanner.Err(); err != nil{
		return ""
	}
	return ""
}
