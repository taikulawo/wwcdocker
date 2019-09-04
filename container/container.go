package container

import (
	"fmt"

	sub "github.com/iamwwc/wwcdocker/cgroups/subsystems"
)

const (
	ContainerMountRoot          = "/var/run/wwcdocker/mount/%s"
	ContainerWriteLayerRoot     = "/var/run/wwcdocker/writelayer/%s"
	ContainerReadLayerRoot      = "/var/run/wwcdocker/readlayer/%s"
	DefaultContainerLogLocation = "/var/run/wwcdocker/log/%s"
	DefaultContainerInfoDir     = "/var/run/wwcdocker/info/"
)

// ContainerInfo has all container information
type ContainerInfo struct {
	Name          string             `json:"name"`
	ID            string             `json:"id"`
	Pid           int                `json:"pid"`
	ImageName     string             `json:"imageName"`
	Env           []string           `json:"env"`
	VolumePoints  map[string]string  `json:"volumePoints"`
	InitCmd       []string           `json:"initCmd"`
	CreateTime    string             `json:"createTime"`
	ResourceLimit *sub.ResourceConfig `json:"resourceLimit"`
	EnableTTY     bool               `json:"enableTty"`
	Detach        bool               `json:"detach"`
	FilePath      map[string]string  `json:"filePath"`
}

func getCwdFromID (id string) string {
	return fmt.Sprintf(ContainerMountRoot, id)
}
