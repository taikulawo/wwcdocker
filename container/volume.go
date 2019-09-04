package container

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"syscall"

	"github.com/iamwwc/wwcdocker/common"
	log "github.com/sirupsen/logrus"
)

/**
NewWorkSpace -> CreateWriteLayer -> CreateReadLayer
-> CreateMountPoint -> MountVolume
// mount volume into container special path
*/

// MountVolume mounts volume into container
func MountVolume(parentPath, containerPath, id string) error {
	// 将 parent 挂载到 container 的工作目录下的 url中
	if err := os.MkdirAll(parentPath, 0777); err != nil {
		return err
	}
	cwd := getCwdFromID(id)
	mountPointInContainer := path.Join(cwd, containerPath)
	if err := os.MkdirAll(mountPointInContainer, 0700); err != nil {
		return fmt.Errorf("Failed to create volume in container. Error: %v", err)
	}
	dirs := "dirs=" + parentPath
	if err := syscall.Mount(parentPath, containerPath, "aufs",0,dirs); err != nil {
		return err
	}
	return nil
}

// NewWorkSpace create new container working directory
func NewWorkSpace(info *ContainerInfo) string {
	containerName := info.ID
	wlayerPath := path.Join(ContainerWriteLayerRoot, containerName)
	rlayerPath := path.Join(ContainerReadLayerRoot, containerName)
	info.FilePath["readLayerPath"] = rlayerPath
	info.FilePath["writeLayerPath"] = wlayerPath
	if ok, err := common.NameExists(wlayerPath); ok {
		log.Warnf("Found same container write layer exists in %s, Layer will be removed. Error: %s", wlayerPath, err)
	}

	if ok, err := common.NameExists(rlayerPath); ok {
		log.Warnf("Found same container read layer exists in %s, Layer will be removed. Error: %s", rlayerPath, err)
	}
	createNewWriteLayer(wlayerPath)
	createNewReadLayer(rlayerPath)


	mountpath := getCwdFromID(info.ID)
	// 将 write layer 与 read layer 组合挂载成aufs文件系统
	if err := createMountPoint(mountpath, wlayerPath, rlayerPath); err != nil {
		log.Errorf("Fail to mount writelayer and read layer. Error: %v", err)
	}
	if len(info.VolumePoints) > 0 {
		for k, v := range info.VolumePoints {
			if k != "" && v != "" {
				MountVolume(k, v, info.ID)
			}
			continue
			log.Errorf("Invalid mount path %s:%s", k, v)
		}
	}
	return mountpath
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
func createMountPoint(mountpath, wlayerpath, rlayerpath string) error {
	// rlayerpath 就是 镜像的 只读文件夹 位置
	dirs := fmt.Sprintf("%s:%s", wlayerpath, rlayerpath)
	if err := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", mountpath).Start(); err != nil {
		return err
	}
	return nil
}
