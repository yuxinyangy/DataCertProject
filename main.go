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
	fmt.Printf("最新区块的Hash值:%x\n",bc.LastHash)
	//block,err := bc.SaveData([]byte("这里存储上链的数据信息"))
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Printf("区块的高度：%d\n",block.Height)
	//fmt.Printf("区块的PrevHash：%x\n",block.PrevHash)
	//block1 :=bc.QueryBlockHeight(1)
	//if block1 == nil {
	//	fmt.Println("抱歉，输入有误")
	//	return
	//}
	//fmt.Println("区块的高度是:",block1.Height)
	//fmt.Println("区块存的信息是:",string(block1.Data))
	blocks := bc.QueryAllBlocks()
	if len(blocks) == 0 {
		fmt.Println("暂未查询到区块数据")
		return
	}
	for _,block :=range blocks{
		fmt.Printf("高度:%d,哈希:%x,Prev哈希:%x\n",block.Height,block.Hash,block.PrevHash)
	}
	return
	db_mysql.ConnectDB()
	//设置静态资源文件
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")
	beego.Run()
}

