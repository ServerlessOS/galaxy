#!/bin/bash

# 本地目录
localDir="/home/tank/bsz/ryze/rpc/"

# 远程服务器列表
remoteServers=(
  "192.168.1.186"
  "192.168.1.188"
  "192.168.1.189"
  "192.168.1.190"
  "192.168.1.191"
)

# 循环迭代远程服务器，执行 scp 命令
for remoteServer in "${remoteServers[@]}"; do
  remoteDir="${remoteServer}:/home/tank/bsz/ryze/rpc/"

  # 使用 scp 命令进行同步
  scp -r "${localDir}" "${remoteDir}"

  # 检查 scp 命令的退出状态码
  if [ $? -eq 0 ]; then
    echo "Successfully copied to ${remoteServer}"
  else
    echo "Error copying to ${remoteServer}"
  fi
done
