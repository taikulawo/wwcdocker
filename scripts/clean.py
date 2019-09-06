#/usr/bin/python3

# 清理脚本，每次删除之前留下的cgroup，以及挂载点

import re
import subprocess


def call(args):
  return subprocess.run(args, stdout=subprocess.PIPE, stderr=subprocess.STDOUT).stdout

reg = "/var/run/wwcdocker/"

points = []
if __name__ == "__main__":
    lines:bytes = call(["mount"])
    for line in lines.decode("utf-8").split("\n"):
      mountpoint = line.split(" ")[2]
      if line == '':
          continue
      r = re.match(reg,mountpoint)
      if r is not None:
        points.append(mountpoint)