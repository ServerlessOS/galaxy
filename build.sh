#!/bin/bash

# 切换到当前路径下
cd "$(dirname "$0")"

# 遍历每个文件夹并执行 go build
for dir in coordinator-rpc dispatcher-rpc func-manager gateway scheduler-rpc virtualNode-rpc; do
    if [ -d "$dir" ]; then
        echo "Building $dir..."
        cd "$dir"

        # 更新go mod
        go mod tidy

        # 编译并将生成的文件移动到当前路径下的bin文件夹内
        go build -o "../bin/$(basename "$dir")"

        cd ..
    else
        echo "Directory $dir does not exist."
    fi
done

echo "Build completed."
