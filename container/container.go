package container

import (
	"path"

	sub "github.com/iamwwc/wwcdocker/cgroups/subsystems"
)

const (
	// WwcdockerRoot is container root
	WwcdockerRoot               = "/var/lib/wwcdocker/"
	ContainerMountRoot          = WwcdockerRoot + "mnt"
	ContainerWriteLayerRoot     = WwcdockerRoot + "writelayers"
	ContainerReadLayerRoot      = WwcdockerRoot + "readlayers"
	DefaultContainerLogLocation = WwcdockerRoot + "log"
	DefaultContainerInfoDir     = WwcdockerRoot + "info"
)

// ContainerInfo has all container information
type ContainerInfo struct {
	Name          string              `json:"name"`
	ID            string              `json:"id"`
	Pid           int                 `json:"pid"`
	ImageName     string              `json:"imageName"`
	Rm            bool                `json:"rm"` // Remove container after container stopped
	Env           []string            `json:"env"`
	VolumePoints  map[string]string   `json:"volumePoints"`
	InitCmd       []string            `json:"initCmd"`
	CreateTime    string              `json:"createTime"`
	ResourceLimit *sub.ResourceConfig `json:"resourceLimit"`
	EnableTTY     bool                `json:"enableTty"`
	Detach        bool                `json:"detach"`
	FilePath      map[string]string   `json:"filePath"`
}

func getCwdFromID(id string) string {
	return path.Join(ContainerMountRoot, id)
}
