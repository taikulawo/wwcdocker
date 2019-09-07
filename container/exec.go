package container

import (
	"encoding/json"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/iamwwc/wwcdocker/cgroups"
	"github.com/iamwwc/wwcdocker/common"
	log "github.com/sirupsen/logrus"
)

// ExecProcess exec container process and run it
func ExecProcess(process *exec.Cmd, info *ContainerInfo) error {
	now := time.Now().Format("2019-09-03 16:36:05")
	info.CreateTime = now
	if err := process.Start(); err != nil {
		log.Errorf("Failed to Start Process, %v", err)
		return err
	}
	pid := process.Process.Pid
	info.Pid = pid

	return cgroups.CreateAndSetLimit(info.ID, pid, info.ResourceLimit)
}

func recordContainerInfo(info *ContainerInfo) error {
	base := path.Dir(common.DefaultContainerInfoDir)
	if err := os.MkdirAll(base, 0644); err != nil {
		return err
	}

	name := path.Base(common.DefaultContainerInfoDir)
	infoFile, err := os.Create(name)

	if err != nil {
		return err
	}

	i, err := json.Marshal(info)
	if err != nil {
		return err
	}
	content := string(i)
	n, err := infoFile.WriteString(content)
	if err != nil {
		return err
	}
	log.Debugf("%d characters has been written to %s", n, path.Join(common.DefaultContainerInfoDir, name))
	return nil
}
