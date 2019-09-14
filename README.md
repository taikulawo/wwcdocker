# WWCDocker 

[![Build Status](https://travis-ci.com/iamwwc/wwcdocker.svg?branch=master)](https://travis-ci.com/iamwwc/wwcdocker)

<p align="center">
  <img src="images/logo.jpg" align="center" alt="logo" width="300" height="300" style="border-radius:50%;">
</p>

## [中文版本](./README-zh-CN.md)

A simple docker implemention

- namespace for resource isolation
- cgroup for resource limit
- aufs for file sytem

### How to install?

`source <(curl -s -L https://raw.githubusercontent.com/iamwwc/wwcdocker/master/install.sh)`

After installation

Try following for fun!

```
wwcdocker run -ti busybox sh
```

### Development

We develop project in dev branch. After the test done, we will merge it to master ( which means stable verison ), and deploy a release according to git tag

**Dev branches are not guaranteed to be compiled**

**If you want give a try, please build master branch by your own, or download latest stable version from `releases` 😜**

For now, we only support `busybox` image ( which has `sh` installed by default ).

You can find all images we currently supported.

https://github.com/iamwwc/imageshub

So

```
wwcdocker run -ti ubuntu bash
```

maybe not working.

`wwcdocker` don't implement `docker pull` mechanism. All images need to use `docker export` to export complete `rootfs` by hand.

We will take it into account.

You can find more related discussion from here

https://github.com/iamwwc/wwcdocker/issues/2

New features will come in future 😀

Keep a eye on it!

### Development Tools

1. Vscode Remote Development - SSH
2. VMWare

My host is `Windows`, using VMware connect into VM with SSH which running in backgroud

### TODO

- [ ] docker exec
- [ ] docker ps
- [ ] docker container stop
- [ ] docker container rm
- [ ] docker run --network
- [ ] docker network create | rm
