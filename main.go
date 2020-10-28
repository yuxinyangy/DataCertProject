package main

import (
	"DataCertProject/blockchain"
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//实例化一个区块链实例
	bc := blockchain.NewBlockChain()
	fmt.Printf("创世区块的Hash值:%x\n",bc.LastHash)
	block,err := bc.SaveData([]byte("这里存储上链的数据信息"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("区块的高度：%d\n",block.Height)
	fmt.Printf("区块的PrevHash：%x\n",block.PrevHash)
	return
	db_mysql.ConnectDB()
	//设置静态资源文件
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")
	beego.Run()
}

