go-gin-example/
├── configs 配置文件
├── docs    文档集合
├── global  全局变量
├── internal内部模块
├       |-----dao   数据访问层
├       |-----middleware    HTTP中间件
├       |-----model 模型层，用于存放model对象
├       |-----routers   路由相关逻辑
├       |-----service   项目核心业务逻辑
├── pkg 项目相关的模块包
└── storage 项目生成的临时文件
├── scripts 各类构建,安装,分析等操作的脚本
└── third_party 第三方的资源工具，如Swagger UI

 --- 

    conf：用于存储配置文件
    middleware：应用中间件
    models：应用数据库模型
    pkg：第三方包
    routers 路由逻辑处理
    runtime：应用运行时数据
