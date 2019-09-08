#!/bin/bash

# Raw location
# https://raw.githubusercontent.com/iamwwc/wwcdocker/master/install.sh

# TO INSTALL this script, run following
# source <(curl -s -L https://raw.githubusercontent.com/iamwwc/wwcdocker/master/install.sh)

red='\033[0;31m'
plain='\033[0m'

[[ $EUID -ne 0 ]] && echo -e "[${red}Error${plain}] This script must be run as root!" && exit 1

beforepath=$(pwd)

tempdir="/tmp/wwcdockertempdir${RANDOM}"

mkdir $tempdir -p

cd $tempdir

# download wwcdocker binary
curl -s https://api.github.com/repos/iamwwc/wwcdocker/releases/latest \
  | grep browser_download_url \
  | cut -d '"' -f 4 \
  | wget -i -

cwd=$(pwd)
chmod u+x wwcdocker
#export to path
export PATH=$PATH:${cwd}

cd $beforepath

green='\033[0;32m'
echo -e "${green}Type wwcdocker for help :)${plain}"
