#!/bin/bash

# 查询并关闭 coordinator-rpc
if pgrep -x "coordinator-rpc" > /dev/null
then
    echo "Stopping coordinator-rpc..."
    pkill -x coordinator-rpc
else
    echo "coordinator-rpc is not running."
fi

# 查询并关闭 gateway
if pgrep -x "gateway" > /dev/null
then
    echo "Stopping gateway..."
    pkill -x gateway
else
    echo "gateway is not running."
fi

# 查询并关闭 func-manager
if pgrep -x "func-manager" > /dev/null
then
    echo "Stopping func-manager..."
    pkill -x func-manager
else
    echo "func-manager is not running."
fi

# 查询并关闭 dispatcher-rpc
if pgrep -x "dispatcher-rpc" > /dev/null
then
    echo "Stopping dispatcher-rpc..."
    pkill -x dispatcher-rpc
else
    echo "dispatcher-rpc is not running."
fi

# 查询并关闭 scheduler-rpc
if pgrep -x "scheduler-rpc" > /dev/null
then
    echo "Stopping scheduler-rpc..."
    pkill -x scheduler-rpc
else
    echo "scheduler-rpc is not running."
fi

# 查询并关闭 virtualNode-rpc
if pgrep -x "virtualNode-rpc" > /dev/null
then
    echo "Stopping virtualNode-rpc..."
    pkill -x virtualNode-rpc
else
    echo "virtualNode-rpc is not running."
fi

echo "All programs stopped."

