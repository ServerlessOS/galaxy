
1. 应该支持func-manager配置文件的内容，目前仅支持按照镜像名启动镜像
2. node会直接以镜像名启动镜像，一方面是需要配置启动参数由用户指定，另一方面是需要检查机器是否安装了docker，如果没装应该提示
已知bug：
    1. func-manager应该独立管理函数数据不互通，所以多func-manager可能出现找不到函数yaml文件的情况，考虑用数据库解决？
    2. gateway更新目录后应该同步更新client/目前没做真正删除client连接