package cmd

import (
	"io/ioutil"
	"os"
	"path"
	"syscall"

	"github.com/iamwwc/wwcdocker/common"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var RemoveCommand = cli.Command{
	Name: "remove",
	Usage: "Remove container, mount points, read layers, write layers",
	Action: func(context *cli.Context) error {
		// TOOD
		return nil
	},
}

var funcs = map[string]removeFunc{
	"m": removeAllMounts,
	"r": removeAllWriteLayers,
	"w": removeAllWriteLayers,
}

type removeFunc func() error

func removeAll() error {
	return nil
}

func removeAllMounts() error {
	points, err := common.FindMountPoint()
	if err != nil {
		log.Error(err)
	}

	// 先解除挂载点
	for _, point := range points {
		if err := removeAMountPoint(point); err != nil {
			return err
		}
	}

	// 删除 mnt/ 文件夹
	root := common.ContainerMountRoot
	if common.NameExists(root) {
		return os.RemoveAll(root)
	}
	return nil
}

func removeAMountPoint(point string) error {
	if err := syscall.Unmount(point, syscall.MNT_DETACH); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func removeAllReadLayers() error {
	root := common.ContainerReadLayerRoot
	if !common.NameExists(root) {
		log.Debugf("Root read layers don't exist. %s", root)
		return nil
	}
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Error(err)
	}
	for _, f := range files {
		p := path.Join(root, f.Name())
		removeAMountPoint(p)
	}
	return nil
}

// 为什么哟了removeAll却出来一个removeA？
// 后面添加删除单个container，所以这里包装一下，后面代码复用
func removeAReadLayer(path string) error {
	return doRemove(path)
}

func removeAllWriteLayers() error {
	return doRemove(common.ContainerWriteLayerRoot)
}

func doRemove(n string) error {
	if common.NameExists(n) {
		return os.RemoveAll(n)
	}
	return nil
}
