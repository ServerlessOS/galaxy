#!/bin/bash


#!/bin/bash

# 使用 find 命令查找并删除所有以 .txt 结尾的文件
find . -type f -name "*.txt" -exec rm -f {} \;

echo "Deleted all .txt files."


# Define the prefixes
virtual_node_prefix="virtual-node"
scheduler_prefix="scheduler"
dispatcher_prefix="dispatcher"

# Get all Pod names with the specified prefixes
virtual_node_pods=$(sudo kubectl get pods --selector=app=$virtual_node_prefix -o jsonpath='{.items[*].metadata.name}')
scheduler_pods=$(sudo kubectl get pods --selector=app=$scheduler_prefix -o jsonpath='{.items[*].metadata.name}')
dispatcher_pods=$(sudo kubectl get pods --selector=app=$dispatcher_prefix -o jsonpath='{.items[*].metadata.name}')

# Function to get logs and write to file
get_logs_and_write_to_file() {
    local prefix=$1
    local pods=$2
    local filename=$3

    echo "Getting logs for pods with prefix $prefix..."

    # Loop through pods and get logs
    for pod in $pods; do
        echo "Getting logs for $pod..."
        sudo kubectl logs $pod >> $filename
    done

    echo "Logs for pods with prefix $prefix written to $filename"
}

# Get logs for each prefix and write to respective files
get_logs_and_write_to_file $virtual_node_prefix "$virtual_node_pods" "node.txt"
get_logs_and_write_to_file $scheduler_prefix "$scheduler_pods" "scheduler.txt"
get_logs_and_write_to_file $dispatcher_prefix "$dispatcher_pods" "dispatcher.txt"
