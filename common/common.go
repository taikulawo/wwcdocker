package common

const (
	// WwcdockerRoot is container root
	WwcdockerRoot               = "/var/lib/wwcdocker/"
	ContainerMountRoot          = WwcdockerRoot + "mnt"
	ContainerWriteLayerRoot     = WwcdockerRoot + "writelayers"
	ContainerReadLayerRoot      = WwcdockerRoot + "readlayers"
	DefaultContainerLogLocation = WwcdockerRoot + "log"
	DefaultContainerInfoDir     = WwcdockerRoot + "info"
)