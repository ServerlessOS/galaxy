#!/bin/bash

# 启动 coordinator-rpc，并显示日志内容
./coordinator-rpc > coordinator-rpc.log &
echo "coordinator-rpc started, waiting for 3 seconds..."
sleep 3

# 启动 gateway，并显示日志内容
./gateway -c "127.0.0.1" -r "127.0.0.1:16448" > gateway.log &
echo "gateway started, waiting for 3 seconds..."
sleep 3

# 启动 func-manager，并显示日志内容
./func-manager -g "127.0.0.1" -r "127.0.0.1:16449" > func-manager.log &
echo "func-manager started, waiting for 3 seconds..."
sleep 3

# 启动 dispatcher-rpc，并显示日志内容
./dispatcher-rpc -g "127.0.0.1" -r "127.0.0.1:16444" > dispatcher-rpc.log &
echo "dispatcher-rpc started, waiting for 3 seconds..."
sleep 3

# 启动 scheduler-rpc，并显示日志内容
./scheduler-rpc -g "127.0.0.1" -r "127.0.0.1:16445" > scheduler-rpc.log &
echo "scheduler-rpc started, waiting for 3 seconds..."
sleep 3

# 启动 virtualNode-rpc，并显示日志内容
./virtualNode-rpc -g "127.0.0.1" -r "127.0.0.1:16446" > virtualNode-rpc.log &
echo "virtualNode-rpc started, waiting for 3 seconds..."
sleep 3

echo "All components started."
