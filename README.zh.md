##  事件驱动的订单服务

其他语言：

### **[English](README.md)**

这个项目展示了如何在Go中进行事件驱动的微服务。 它包括两个微服务，一个是[“Payment”](https://github.com/jfeng45/payment)服务，另一个是这个。 订单服务调用付款服务进行付款。 你需要同时运行两个项目才能使其正常运行。

## 如何开始

### 安装

#### 下载代码

```
go get github.com/jfeng45/order
```

#### 设置 MySQL

```
Install MySQL
run SQL script in script folder to create database and table
```
#### 安装 NATS

[安装 NATS](https://docs.nats.io/nats-server/installation)

### 启动应用程序

#### 启动 MySQL Server
```
cd [MySQLroot]/bin
mysqld
```
#### 运行主程序
```
cd [rootOfProject]/cmd
go run main.go
```
#### 运行 "Payment" 服务

["Payment"](https://github.com/jfeng45/payment)

## 授权

[MIT](LICENSE.txt) 授权


