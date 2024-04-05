cd /home/tank/bsz/ryze/rpc/virtualNode-rpc

docker build -t virtual-node-rpc:latest . &&  \
docker tag virtual-node-rpc:latest bszpe/virtual-node-rpc:latest

cd /home/tank/bsz/ryze/rpc/dispatcher-rpc

docker build -t dispatcher-rpc:latest . &&  \
docker tag dispatcher-rpc:latest bszpe/dispatcher-rpc:latest

cd /home/tank/bsz/ryze/rpc/scheduler-rpc
docker build -t scheduler-rpc:latest . &&  \
docker tag scheduler-rpc:latest bszpe/scheduler-rpc:latest
