package main

import (
	"DataCertProject/blockchain"
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//1.生成第一个区块
	block := blockchain.NewBlock(0,[]byte{},[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

	fmt.Println(block)
	fmt.Printf("区块Hash：%x",block.Hash)
	return
	//1.链接数据库
	db_mysql.ConnectDB()
	//设置静态资源文件
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")
	beego.Run()
}

