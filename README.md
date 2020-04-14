# A Event Driven Go Order Microservice

Other language: 
### **[中文](README.zh.md)**
 
This is a project to show case how to do event driven Microservice in Go. It includes two Microservices, one is  ["Payment"](https://github.com/jfeng45/payment) service, the other is this one. The Order service calls the Payment service to make a payment. You need to run both projects together to make it work. 

## Getting Started

### Installation

#### Download Code

```
go get github.com/jfeng45/order
```

#### Set Up MySQL

```
Install MySQL
run SQL script in script folder to create database and table
```
#### Install NATS

[Install NATS](https://docs.nats.io/nats-server/installation)

### Start Application

#### Start MySQL Server
```
cd [MySQLroot]/bin
mysqld
```
#### Run main
```
cd [rootOfProject]/cmd
go run main.go
```
#### Run "Payment" Service

["Payment"](https://github.com/jfeng45/payment)

## License

[MIT](LICENSE.txt) License



