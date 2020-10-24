package main

import (
	"DataCertProject/blockchain"
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//user1 := models.User{
	//	Id:       1,
	//	Phone:    "19970349328",
	//	Password: "111111",
	//}
	//fmt.Println("内存中的数据User1",user1)
	//
	//_,_ = json.Marshal(user1)
	//xmlBytes, _ := xml.Marshal(user1)
	//fmt.Println(string(xmlBytes))
	//
	//var user2 models.User
	//xml.Unmarshal(xmlBytes, &user2)
	//fmt.Println("反序列化的User2:", user2)
	//return



	//1.生成第一个区块
	block := blockchain.CreateGenesisBlock()
	fmt.Println(block)
	fmt.Printf("区块Hash：%x",block.Hash)

	//db,err:=bolt.Open("chain.db",0600,nil)
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer db.Close()
	////操作chain.db文件
	//db.Update(func(tx *bolt.Tx) error {
	//	var tong  *bolt.Bucket
	//	tong = tx.Bucket([]byte(BUCKET_NAME))
	//   if tong == nil {
	//		tong,err = tx.CreateBucket([]byte(BUCKET_NAME))
	//		if err != nil {
	//			return err
	//		}
	//
	//   }
	//   //先查看获取看桶中是否已包含要保存的区块
	//	lastBlock := tong.Get([]byte("lastHash"))
	//	blockHash,err := block.Serialize()
	//	if err != nil {
	//		return nil
	//	}
	//	if lastBlock == nil {
	//		tong.Put(block.Hash,blockHash)
	//		tong.Put([]byte("lastHash"),blockHash)
	//	}
	//   return nil
	//})
	//return
	//1.链接数据库
	db_mysql.ConnectDB()
	//设置静态资源文件
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")
	beego.Run()
}

