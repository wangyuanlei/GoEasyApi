# GoEasyApi

#### 介绍
用go 实现, 数据库快速配置出对应的api接口. 减少后端繁琐的开发逻辑

#### 软件架构
软件架构说明


#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 目录说明
    libraries       //功能库
    model           //模块库
    database        //各种数据库支持
    main.go         //启动文件


#### 结构逻辑说明
    1. 接口 层级为  服务 -> 分组 -> 接口
        1. 一个服务对应一个客户端. 
        2. 分组 为1级 分类, 作为接口的归类. 
        3. 接口 外部参数 组合成sql 然后获得结果.

    1. 管理员管理信息. 
        1. 管理用户 分为 超级管理员 和 普通管理员
        2. 超级管理员能管理  普通管理员 信息.
        3. 普通管理员 自能管理自己对应的 服务.
    
    2. 


#### 环境搭建命令记录
    go mod init GoEasyApi  #初始化项目
    go get -u github.com/gin-gonic/gin  #当前项目加载 gin 框架
    go run main.go  #启动go
    go build  #打包项目
    go test -v  #测试项目
    go mod tidy  #整理依赖
    go version  #查看go版本
    go env  #查看go环境
 
    go env -w CGO_ENABLED=1  #编译sqlite数据库 需要这个设置
    