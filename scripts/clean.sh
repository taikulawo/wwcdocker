#!/bin/bash

# 每次测试之前将之前的mountpoint全部卸载
# 删除 wwcdocker 的cgroup目录

# set -x

root="/var/lib/wwcdocker/mnt"

info=$(mount)

# 以 wwcdocker 开头的挂载点
mountpoints=()

while read -r line;do
  point=$(echo $line | cut -d' ' -f3 -)
  if [[ $point =~ $root ]]; then
    mountpoints+=($point)
  fi
done <<< $info

# 解挂文件系统
for i in "${mountpoints[@]}"; do
  umount $i
done

# 数组空格分离
# 还是Python写起来方便...
# 系统自带而且写起来还爽,适合当脚本写

cgroups=()

mountinfo=$(mount)

# 获取全部的 cgroup
while read -r line; do
  # here string
  # https://unix.stackexchange.com/questions/80362/what-does-mean
  type=$(cut -d ' ' -f1 <<< $line)
  path=$(cut -d ' ' -f3 <<< $line)
  if [[ $type == "cgroup" ]]; then
    sys=$(echo $path | rev | cut -d '/' -f1 | rev)
    cgroups+=($sys)
  fi
done <<< $mountinfo


# 删除每一个cgroup中的 wwcdocker
for system in "${cgroups[@]}"; do
  target="/sys/fs/cgroup/${system}/wwcdocker"
  if [[ -d $target ]]; then
    printf "Removing ${system} system cgroup\n"
    cgdelete -r "${system}:wwcdocker"
  fi
done