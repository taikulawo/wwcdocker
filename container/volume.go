package container

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/iamwwc/wwcdocker/common"
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
	if err := syscall.Mount(parentPath, containerPath, "aufs", 0, dirs); err != nil {
		return err
	}
	return nil
}

// NewWorkspace create new container working directory
func NewWorkspace(root, containerID,imageName string, volumes map[string]string) string {
	wlayerPath := path.Join(common.ContainerWriteLayerRoot, containerID)
	rlayerPath := path.Join(common.ContainerReadLayerRoot, containerID)
	createNewWriteLayer(wlayerPath)
	createNewReadLayer(rlayerPath,imageName)

	mountpath := getCwdFromID(containerID)
	// 将 write layer 与 read layer 组合挂载成aufs文件系统
	if err := createMountPoint(mountpath, wlayerPath, rlayerPath); err != nil {
		log.Errorf("Fail to mount writelayer and read layer. Error: %v", err)
	}
	if len(volumes) > 0 {
		for k, v := range volumes {
			if k != "" && v != "" {
				MountVolume(k, v, containerID)
			}
			log.Errorf("Invalid mount path %s:%s", k, v)
			continue
		}
	}
	return mountpath
}

func createNewWriteLayer(name string) error {
	if err := os.MkdirAll(name, 0777); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

const (
	imagesHub = "https://raw.githubusercontent.com/iamwwc/imageshub/master/"
)

func findImages(imagePath, image string) string {
	if !common.NameExists(imagePath) {
		common.Must(os.MkdirAll(imagePath,0644))
	}
	name := image + ".tar"
	absolutePath := path.Join(imagePath,name)
	if !common.NameExists(absolutePath) {
		common.DownloadFromUrl(imagesHub + name, imagePath)
	}
	return absolutePath
}

// createNewReadLayer create working folder from the given image.
// root is container read layer folder
// such as /var/lib/wwcdocker/readlayer/213kjassdqw/
func createNewReadLayer(root, imageLayer string) error {
	log.Debugf("Image: %s",imageLayer)
	imagePath := path.Join(common.WwcdockerRoot, "images")
	tarURL := findImages(imagePath,imageLayer)
	_, err := os.Stat(tarURL)
	if os.IsNotExist(err) {
		return fmt.Errorf("busybox.tar don't exist in %s",tarURL)
	}
	if err := os.MkdirAll(root,0644); err != nil {
		return err
	}
	if _, err := exec.Command("tar","-xvf",tarURL,"-C",root).CombinedOutput(); err != nil {
		return fmt.Errorf("untar error. %v", err)
	}
	return nil
}
func createMountPoint(mountpath, wlayerpath, rlayerpath string) error {
	// rlayerpath 就是 镜像的 只读文件夹 位置
	if err := os.MkdirAll(mountpath,0777); err != nil {
		return err
	}
	dirs := fmt.Sprintf("dirs=%s:%s", wlayerpath, rlayerpath)
	if _, err := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", mountpath).CombinedOutput(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// DeleteWorkSpace deletes write layer, unmounts mountpoint
func DeleteWorkSpace(containerID string) {
	deleteMountPoint(path.Join(common.ContainerMountRoot, containerID))
	deleteWriteLayer(path.Join(common.ContainerWriteLayerRoot, containerID))
}

func deleteMountPoint(mntpoint string) error {
	if err := syscall.Unmount(mntpoint,syscall.MNT_DETACH); err != nil {
		return err
	}
	_, err := os.Stat(mntpoint)
	if os.IsExist(err) {
		return os.RemoveAll(mntpoint)
	}
	return nil
}

// deleteWriteLayer deletes container write layer located on writelayer folder
func deleteWriteLayer(path string) error {
	return os.RemoveAll(path)
}