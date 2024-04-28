#!/bin/bash

# 检查是否存在与脚本同级目录下的 bin 文件夹，如果不存在则创建
if [ ! -d "./bin" ]; then
    mkdir bin
fi

# 遍历子目录，执行 go build 并将生成的可执行文件移动到 bin 目录下
for dir in gateway coordinator-rpc dispatcher-rpc func-manager scheduler-rpc virtualNode-rpc; do
    echo "Building $dir..."
    cd "$dir" || exit
    go build -o "../bin/$dir"
    cd ..
done

echo "Build completed!"