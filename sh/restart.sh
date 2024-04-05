
cd /home/tank/bsz/ryze/rpc
sudo kubectl delete -f yaml/

while true; do
    # 使用 kubectl 命令检查 default 命名空间中是否有 Pod
    pod_count=$(kubectl get pods --namespace=default --no-headers | wc -l)

    if [ "$pod_count" -gt 0 ]; then
        echo "Pods found in default namespace. Sleeping for 2 seconds..."
        sleep 2
    else
        echo "OK: No pods found in default namespace."
        break
    fi
done

sudo kubectl apply -f yaml/

