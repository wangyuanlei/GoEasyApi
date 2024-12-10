# GoEasyAPI

## 介绍
GoEasyAPI是一套基于go语言开发的自动化服务端API生成平台, 用户通过设计sql语句自己的数据api接口, 即可生成适用于不同需求的服务端代码. 



## 背景
在软件应用的开发中, 后端服务是一个繁杂的过程, 这个过程中, 会遇到各种关键问题, 包括服务端的架构, 服务端的部署, 数据库的设计和开发, 服务端压力均衡, 系统稳定性, 接口可用性等等问题. 为了实现这些问题的需求, 往往会让后端开发的整个过程变得比较繁琐, 试错成本比较高. 那么, 我们的目标是实现一个通用的后台API, 不同的功能模块通过统一的平台进行管理, 使得用户只要关注核心逻辑和数据结构即可, 不用再为了一些重复劳动制造轮子. 

 
## 目录结构
    controller             //控制器层
    core                   //核心
    libraries              //功能库
    helper                 //辅助库
    model                  //模块库
    structs                //各种结构定义
    router                 //路由层
    static                 //静态资源
    vue                    //前端代码
    main.go                //启动文件
    config.yaml            //配置文件
    db.sql                 //数据库文件
    GoEasyApi_linux        //linux下编译好的服务程序
    GoEasyApi_win.exe      //windows下编译好的服务程序

## 编译和运行
### 编译
   1. git 下载后. go mod tidy 加载依赖
   2. 执行 go run main.go 运行程序
   3. go build -o GoEasyApi 生成可执行文件

### windiws下运行
   1. 拷贝文件 config.yml, db.sql, GoEasyApi_win.exe 到 指定目录下
   2. 双击 GoEasyApi_win.exe 运行程序

### linux下运行
   1. 拷贝文件 config.yml, db.sql, GoEasyApi_linux, start.sh 到 指定目录下
   2. 运行程序
        ```shell
        sh start.sh  
        ```
### 编译运行
   docker build -t go_easy_api .
   docker run -d -p 8008:8008 go_easy_api

## 使用说明
1. 打开浏览器, 输入 http://localhost:8008 
2. 点击登录, 输入用户名和密码, 默认用户名和密码是admin  Qwert!@#456

## 接口说明
1. POST /user/register 注册接口
2. POST /user/login 登录接口
3. GET /user/user_info 获取用户信息接口
4. POST /api/xx/xxx 请求接口
5. GET /api/xx/xxx 请求接口


觉得不错的话，别忘 **star** 👏