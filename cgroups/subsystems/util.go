package subsystems

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	subSystem = "/([\\w|,]+?)/wwcdocker/"
)

func createCgroup(cpath string, pid int) error {
	if err := os.MkdirAll(cpath, 0644); err != nil {
		return fmt.Errorf("Failed to create cgroup %s, error: %v", cpath, err)
	}
	r := regexp.MustCompile(subSystem)
	subSystemName := r.FindStringSubmatch(cpath)[1]

	taskpath := path.Join(cpath, "tasks")
	if err := ioutil.WriteFile(taskpath,[]byte(strconv.Itoa(pid)),0644); err != nil {
		log.Error(err)
		return err
	}
	log.Debugf("Limit Resource in %s", subSystemName)
	return nil
}

func GetCgroupPath(subSysName string, containerId string) (string, error) {
	cgroupDir := FindCgroupMountPoint(subSysName)
	return path.Join(cgroupDir, "wwcdocker", containerId), nil
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
