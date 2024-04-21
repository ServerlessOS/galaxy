1. node与func-manager互动并正确打开镜像
      1. func-manager
           1. 增删查yaml配置
           2. 向顶层控制器注册，同步函数信息
           3. 启动时掌握通用的gateway地址
      2. gateway
           1. 做func-manager列表管理，为了效率，将node分配给固定func-manger管辖
           2. 接受func-manager注册/node注册时返回一个func-manager地址
           3. gateway调用dispatcher后需要改成拿到一个地址，需要访问这个http地址获取函数计算结果
           4. node调用gateway需要直接转发消息给对应的dispatcher地址 
      3. 重构node
          1. 检测docker环境
          2. 与func-manager互动能力
          3. 本地发起docker指令并主动联系gateway向dispatcher注册函数实例
          4. 启动时掌握gateway的通用地址
2. 注册机制大改动与gateway信息
   1. 地址支持由接收方解析
   2. 顶层控制器总是最先启动的，gateway仅掌握顶层控制器的地址，所有组件的注册均先访问gateway，然后被转发到顶层控制器
   3. 由顶层控制器负责告诉新gateway它需要负责的dispatcher地址，而不是gateway转发时自动缓存
   4. gateway需要保存的信息为1个顶层控制器地址、n个dispatcher、n个func-manager


已知bug：
    func-manager应该独立管理函数数据不互通，所以多func-manager可能出现找不到函数yaml文件的情况，考虑用数据库解决？

将getlocalIP列装到全局