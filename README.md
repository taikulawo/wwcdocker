### Docker Impl

这是一个简易的Docker实现

- namespace进行资源隔离
- cgroup 资源限制
- aufs作为底层文件系统与镜像的实现

后续会逐渐添加新的功能

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