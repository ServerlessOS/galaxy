# 简介
galaxy是一个为了解决k8s中心化瓶颈、提高k8s负载上限、合理分配性能资源的新kubernetes控制平面。它在依然提供了应用部署、维护和扩展的基本机制的前提下，提高了可用性和负载快速波动时的资源利用率，并提升了理论性能上限。
## 特征：

- 全组件可伸缩、自动伸缩，组件实例间支持负载均衡，提升服务容量上限，解决性能倾斜问题、提高资源利用率。
- 组件划分粒度更细，组件自身功能更明确准确，使得链路中性能瓶颈定位更准确，扩容效果更好，资源利用率更高。
- 不含中心化瓶颈，相比现有控制平面的理论性能上限更高，容灾能力更强。
- 实现了面向组件的监控告警系统，实时监控控制平面运行状态并自动处理。
- 提供了容器创建、执行、删除、重启等管理能力，支持配置文件管理与批量启动，支持按需负载均衡部署。
# 快速开发
**需要您有一个go环境（go version ≥ go1.21）**
```
git clone https://github.com/ServerlessOS/galaxy
cd galaxy
```
## 模块：

- [cluster-manager](https://github.com/ServerlessOS/galaxy/tree/main/cluster-manager) ：组件性能监控、历史性能数据管理与自动扩缩容请求发起
- [coordinator-rpc](https://github.com/ServerlessOS/galaxy/tree/main/coordinator-rpc) ： 顶层控制器，组件启动时向顶层控制器注册，获取相应的信息并全局通报
- [dispatcher-rpc](https://github.com/ServerlessOS/galaxy/tree/main/dispatcher-rpc) ：调度器，负责记录用户实例和内部模块地址，分发对应请求
- [func-manager](https://github.com/ServerlessOS/galaxy/tree/main/func-manager) ：函数管理，管理函数基本信息，作为实例启动时的配置信息
- [gateway](https://github.com/ServerlessOS/galaxy/tree/main/gateway) ：网关，向上配合外源的DDNS实现多IP解析单域名，向下缓存和分发请求信息，是唯一公开的请求入口
- [scheduler-rpc](https://github.com/ServerlessOS/galaxy/tree/main/scheduler-rpc) ：负责支持负载均衡功能，决定实例应该在哪台节点上部署
- [virtualNode-rpc](https://github.com/ServerlessOS/galaxy/tree/main/virtualNode-rpc) ：物理服务器上的节点程序，负责在服务器上执行具体指令，以及上传性能数据
## 仓库：

- [constant](https://github.com/ServerlessOS/galaxy/tree/main/constant)：全局统一使用的常量库，例如各组件默认端口等
- [proto](https://github.com/ServerlessOS/galaxy/tree/main/proto)：全局统一使用的rpc接口库，各组件在本地实现接口的对应功能供外界调用
# 快速使用
## 编译安装
可以参考根目录下的[build.sh](https://github.com/ServerlessOS/galaxy/blob/main/build.sh)，或在各组件对应的go.mod所在目录下，使用以下命令：
```
 go mod tidy
 go build
```
## 启动
启动与停止各组件可参考[run.sh](https://github.com/ServerlessOS/galaxy/blob/main/bin/run.sh)和[kill.sh](https://github.com/ServerlessOS/galaxy/blob/main/bin/kill.sh)文件，具体来说：

- coordinator-rpc为顶层控制器，无需启动参数
- gateway网关，要求使用`-c` 指明顶层控制器的地址，可以使用`-r`和`-p`分别指定本地rpc和http绑定的地址也可以不指定
- 其它组件，要求使用`-g`指明网关地址，可以使用`-r`指定本地rpc绑定的地址
## 使用
目前仅实装了`getGatewayList`、`create`、`createFile`api接口

- `http://gatewayAddress:port/getGatewaylist`其中`gatewayAddress:port`需要替换为对应地址，用于上游DDNS服务实现网址网关一对多负载均衡
- `http://gatewayAddress:port/create?funcName=ubuntu&requireCpu=4&requireMem=4`用于发起新实例创建请求
- `http://gatewayAddress:port/createFile?funcName=ubuntu`用于创造配置文件，配置文件中心还未实装所以该功能目前没用
# 
