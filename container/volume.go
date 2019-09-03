package container

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/iamwwc/wwcdocker/common"
	log "github.com/sirupsen/logrus"
)

/**
NewWorkSpace -> CreateWriteLayer -> CreateReadLayer
-> CreateMountPoint -> MountVolume
// mount volume into container special path
*/

// MountVolume mounts volume into container
func MountVolume(parentPath, containerPath, containerName string) {
	// 将 parent 挂载到 container 的工作目录下的 url中
	if err := os.MkdirAll(parentPath, 0777); err != nil {

	}
	// mntUrl :=
}

// NewWorkSpace create new container working directory
func NewWorkSpace(info *ContainerInfo) string {
	containerName := info.Id
	wlayerPath := path.Join(ContainerWriteLayerRoot, containerName)
	rlayerPath := path.Join(ContainerReadLayerRoot, containerName)
	info.FilePath["readLayerPath"] = rlayerPath
	info.FilePath["writeLayerPath"] = wlayerPath
	if ok, err := common.NameExists(wlayerPath); ok {
		log.Warnf("Found same container write layer exists in %s, Layer will be remove", layerPath)
	}
	createNewWriteLayer(wlayerPath)
	createNewReadLayer(rlayerPath)

	// 将 write layer 与 read layer 组合挂载成aufs文件系统
	createMountPoint(wlayerPath, rlayerPath,info)
	if len(info.VolumePoints) > 0 {
		for k, v := range info.VolumePoints {
			if k != "" && v != "" {
				MountVolume(k, v, info.Id)
			}
			continue
			log.Errorf("Invalid mount path %s:%s", k, v)
		}
	}
}

func createNewWriteLayer(name string) error {
	if err := os.Mkdir(name, 0777); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// createNewReadLayer create working folder from the given image.
func createNewReadLayer(imageLayer string) {

}
func createMountPoint(wlayerpath, rlayerpath string, info *ContainerInfo) {
	mountpath := fmt.Sprintf(ContainerMountRoot,info.Id)
	info.FilePath["mountpath"] = mountpath
	// rlayerpath 就是 镜像的 只读文件夹 位置
	dirs := fmt.Sprintf("%s:%s", wlayerpath, rlayerpath)
	if err := exec.Command("mount", "-t", "aufs", "-o", dirs, "none",mountpath).Start(); err != nil {

	}
}
