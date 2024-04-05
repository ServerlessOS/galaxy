#!/bin/bash

# 目标主机列表
target_hosts=("192.168.1.186" "192.168.1.188" "192.168.1.189" "192.168.1.190" "192.168.1.191")

# 远程脚本路径
remote_script_path="/home/tank/bsz/ryze/rpc/sh/"

# 循环迭代目标主机
for host in "${target_hosts[@]}"; do
    # 使用ssh远程执行脚本
    ssh "$host" "cd \"$remote_script_path\" && ./build_all_rpc.sh"
    
    # 检查ssh命令的返回值，如果不为0，则打印错误信息
    if [ $? -ne 0 ]; then
        echo "Error: Failed to execute script on $host."
    else
        echo "Script executed successfully on $host."
    fi
done

