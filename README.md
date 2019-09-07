### Docker Impl


[![Build Status](https://travis-ci.com/iamwwc/wwcdocker.svg?branch=master)](https://travis-ci.com/iamwwc/wwcdocker)

这是一个简易的Docker实现

- namespace进行资源隔离
- cgroup 资源限制
- aufs作为底层文件系统与镜像的实现

### 如何安装
`bash <(curl -s https://raw.githubusercontent.com/iamwwc/wwcdocker/master/install.sh)`

### 开发
默认在 `dev` 分支开发，开发完成，测试通过之后会发布至 `master` 分支，并构建 `release`

**dev 分支不保证能通过 go build**

**如果想尝试，请编译 master 分支，或者在 releases 下载最新稳定版本 😜**

现在只支持 busybox 镜像，而 busybox 只配备了 `sh`，并没有 `bash`

由于 wwcdocker 并没有实现类似 `docker pull` 的机制
所以全部的镜像都需要 `docker export` 来获得完整的 `rootfs`

你可以在下面找到支持的镜像

`https://github.com/iamwwc/imageshub`

后续考虑实现 docker pull
更多细节你可以在这里看到

https://github.com/iamwwc/wwcdocker/issues/2

后续会逐渐添加新的功能😀

### 开发工具

1. VSCode Remote Development -SSH
2. VMWare

宿主机是Windows，本地后台运行 `Ubuntu 18.04.2 LTS`

SSH挂载目录远程开发

### TODO

- [ ] docker exec
- [ ] docker ps
- [ ] docker container stop
- [ ] docker container rm
- [ ] docker run --network
- [ ] docker network create | rm