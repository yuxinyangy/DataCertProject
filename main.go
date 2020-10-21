package main

import (
	"DataCertProject/blockchain"
	"DataCertProject/db_mysql"
	"DataCertProject/models"
	_ "DataCertProject/routers"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	user1 := models.User{
		Id:       1,
		Phone:    "19970349328",
		Password: "111111",
	}
	fmt.Println("内存中的数据User1",user1)

	_, _ = json.Marshal(user1)
	xmlBytes, _ := xml.Marshal(user1)
	fmt.Println(string(xmlBytes))

	var user2 models.User
	xml.Unmarshal(xmlBytes, &user2)
	fmt.Println("反序列化的User2:", user2)
	return



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

