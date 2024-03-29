package common

import (
	"fmt"
	"strings"
	"syscall"

)

var (
	// Mount is shortcut of syscall.Mount
	Mount = syscall.Mount
	// Unmount is shortcut of syscall.Unmount
	Unmount = syscall.Unmount
)

// FindMountPoint gets all wwcdocker mount point
// like /var/run/wwcdocker/mnt/balabalabala
func FindMountPoint() ([]string, error) {
	v := Must2(Exec("mount"))
	switch tp := v.(type) {
	case string:
		return parseMountInfo(tp),nil
	default:
		// TODO:
		// replace fmt.Errorf by LogAndErrorf
		return nil, fmt.Errorf("Unexpected type: %T", tp)
	}
}

func parseMountInfo(info string) (result []string) {
	arrays := strings.Split(info, "\n")
	root := WwcdockerRoot

	for _, value := range arrays {
		point := strings.Split(value, " ")[2]
		if index := strings.Index(point, root); index != -1 {
			result = append(result, point)
		}
	}
	return result
}
