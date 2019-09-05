package subsystems

import (
	"regexp"
	"testing"
)

func TestFindCgroupMountPoint(t *testing.T) {
	sub := "cpu"
	t.Logf("mount point is %s", FindCgroupMountPoint(sub))
}

func TestRegexp(t *testing.T) {
	cpath := "/sys/fs/cgroup/cpu,cpuacct/wwcdocker/iwehjwqkenkwq"
	r := regexp.MustCompile(subSystem)
	subSystemName := r.FindStringSubmatch(cpath)
	t.Log(subSystemName)
}
