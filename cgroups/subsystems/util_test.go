package subsystems

import "testing"

func TestFindCgroupMountPoint(t *testing.T) {
	sub := "cpu"
	t.Logf("mount point is %s",FindCgroupMountPoint(sub))
}
