package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"golang.org/x/crypto/ssh"
)

func main() {
	// SSH连接配置
	sshConfig := &ssh.ClientConfig{
		User: "your_username",
		Auth: []ssh.AuthMethod{
			ssh.Password("your_password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// SSH连接
	conn, err := ssh.Dial("tcp", "your_ssh_host:22", sshConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 创建一个新的SSH隧道
	tunnel := "127.0.0.1:3306"
	listener, err := conn.Listen("tcp", tunnel)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	// 使用隧道连接数据库
	dsn := "your_username:your_password@tcp(127.0.0.1:3306)/your_database"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("数据库连接成功")
}
