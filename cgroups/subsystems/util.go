package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func GetCgroupPath(subSysName string, containerId string) (string, error) {
	cgroupDir := FindCgroupMountPoint(subSysName)
	containerPath := path.Join(cgroupDir, "wwcdocker", containerId)
	if err := os.MkdirAll(containerPath, 0644); err != nil {
		return "", fmt.Errorf("Fail to create cgroup %s, error: %v", containerPath, err)
	}
	return containerPath, nil
}

func FindCgroupMountPoint(subsystem string) string {
	file, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// / /sys/fs/cgroup/memory rw,nosuid,nodev,noexec,relatime shared:19 - cgroup cgroup rw,memory
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Split(text, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return ""
	}
	return ""
}
