package container

import (
	sub "github.com/iamwwc/wwcdocker/cgroups/subsystems"
)

// ContainerInfo has all container information
type ContainerInfo struct {
	Name          string             `json:"name"`
	Id            string             `json:"id"`
	Pid           string             `json:"pid"`
	ImageName     string             `json:"imageName"`
	Env           []string           `json:"env"`
	MountPoints   map[string]string  `json:"mountpoints"`
	InitCmd       string             `json:"initCmd"`
	CreateTime    string             `json:"createTime"`
	ResourceLimit sub.ResourceConfig `json:"resourceLimit"`
	EnableTTY     bool               `json:"enabletty"`
	Detach        bool               `json:"detach"`
}

