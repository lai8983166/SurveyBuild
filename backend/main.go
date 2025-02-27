package main

import (
	"backend/config" // 替换为你自己的项目路径
	"fmt"
	"log"
)

func main() {
	// 初始化数据库
	config.InitDatabase()

	// 检查数据库连接
	if config.DB != nil {
		fmt.Println("数据库连接成功")
	} else {
		log.Fatal("数据库连接失败")
	}

	// 其他应用逻辑（例如启动HTTP服务等）
}
