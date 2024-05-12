# 简介
该项目为一个为了解决k8s中心化瓶颈、提高k8s负载上限、合理分配性能资源的新控制平面。以微服务的思想，通过对功能做拆分细化，实现了负载性能快速扩容，组件资源灵活分配的特性。
# 快速使用
## 编译安装
可以参考根目录下的[build.sh](https://github.com/ServerlessOS/galaxy/blob/main/build.sh)，或在各组件对应的go.mod所在目录下，使用以下命令：
```
 go mod tidy
 go build
```
## 启动
启动与停止各组件可参考./bin/run.sh和./bin/kill.sh文件，具体来说：

- coordinator-rpc为顶层控制器，无需启动参数
- gateway网关，要求使用`-c` 指明顶层控制器的地址，可以使用`-r`和`-p`分别指定本地rpc和http绑定的地址也可以不指定
- 其它组件，要求使用`-g`指明网关地址，可以使用`-r`指定本地rpc绑定的地址
# 使用
目前仅实装了`getGatewayList`、`create`、`createFile`api接口

- `http://gatewayAddress:port/getGatewaylist`其中`gatewayAddress:port`需要替换为对应地址，用于上游DDNS服务实现网址网关一对多负载均衡
- `http://gatewayAddress:port/create?funcName=ubuntu&requireCpu=4&requireMem=4`用于发起新实例创建请求
- `http://gatewayAddress:port/createFile?funcName=ubuntu`用于创造配置文件，配置文件中心还未实装所以该功能目前没用

