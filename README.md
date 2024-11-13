# GoEasyAPI

## 介绍
GoEasyAPI是一套基于go语言开发的自动化服务端API生成平台, 用户通过设计自己的数据结构和产品逻辑, 即可生成适用于不同需求的服务端代码. 



## 背景
在软件应用的开发中, 后端服务是一个繁杂的过程, 这个过程中, 会遇到各种关键问题, 包括服务端的架构, 服务端的部署, 数据库的设计和开发, 服务端压力均衡, 系统稳定性, 接口可用性等等问题. 为了实现这些问题的需求, 往往会让后端开发的整个过程变得比较繁琐, 试错成本比较高. 那么, 我们的目标是实现一个通用的后台API, 不同的功能模块通过统一的平台进行管理, 使得用户只要关注核心逻辑和数据结构即可, 不用再为了一些重复劳动制造轮子. 



## 开发目标和进度
1. 一个基本可用的产品底层框架, 基于此框架可以运行生成一套完整的API
2. 模块化功能, 提供基础的模块功能, 可以实现基于模块定义快捷生成API
3. 用户登录和管理模块
4. 数据库支持, 支持多种数据库, 包括sqlite, mysql, postgresql等等
5. 安全性支持, 支持数据加密, 支持密钥管理
6. AI支持, 支持与外部AI接口的对接
7. 管理后台, 支持用户管理, 服务管理, 数据库管理等等
8. 文档生成, 支持自动生成API文档
9. 测试支持, 支持自动测试API接口

 
## 目录结构
    controller       //控制器层
    core            //核心
    libraries       //功能库
    helper          //辅助库
    model           //模块库
    structs        //各种结构定义
    router          //路由层
    static          //静态资源
    vue            //前端
    main.go         //启动文件
    config.yaml       //配置文件
    db.sql           //数据库文件
    GoEasyApi       linux下编译好的服务程序
    GoEasyApi.exe   windows下编译好的服务程序

## 编译和运行
### 编译
   1. git 下载后. go mod tidy 加载依赖
   2. 执行 go run main.go 运行程序
   3. go build -o GoEasyApi 生成可执行文件

### windiws下运行
   1. 拷贝文件 config.yml, db.sql, GoEasyApi.exe 到 指定目录下
   2. 双击 GoEasyApi.exe 运行程序

### linux下运行
   1. 拷贝文件 config.yml, db.sql, GoEasyApi, start.sh 到 指定目录下
   2. 运行程序
        ```shell
        sh start.sh  
        ```

觉得不错的话，别忘 **star** 👏