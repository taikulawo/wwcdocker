package common

import "testing"

func TestFindMountPoint(t *testing.T) {
	result, err := FindMountPoint()
	Must(err)
	t.Logf("Mountinfo are %s",result)
}
