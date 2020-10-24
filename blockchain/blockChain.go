package blockchain

import "github.com/bolt-master"

//桶的名称，该桶用于装区块信息
var BUCKET_NAME  = "blocks"
//表示最新的区块的key名
var LAST_KEY = "lasthash"
/*
区块链结构体实例:用于描述或表示代表一条区块链
* 该条区块链包括以下功能:
*           ①将新产生的区块与已有的区块链接起来
            ②可以查询某个区块的信息
            ③可以将所有区块进行遍历，输出区块信息
 */
type BlockChain struct {
	LastHash []byte//最新区块的hash
	BoltDb *bolt.DB
}



func NewBlockChain() BlockChain  {
	genesis :=CreateGenesisBlock()//创世区块
	db,err:=bolt.Open("chain.db",0600,nil)
	if err != nil {
		panic(err.Error())
	}
	bl := BlockChain{
		LastHash: genesis.Hash,
		BoltDb:   db,
	}
	return bl
}

/*
调用BlockChain的该SaveBlock方法，该方法可以将一个生成的新区块块保存到chain.db文件中
 */
func (bc BlockChain) SaveBlock(block Block)  {
	db := bc.BoltDb
	//操作chain.db文件
	db.Update(func(tx *bolt.Tx) error {
		var tong  *bolt.Bucket
		tong = tx.Bucket([]byte(BUCKET_NAME))
		if tong == nil {
			tong,err := tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				return err
			}
			//先查看获取看桶中是否已包含要保存的区块
			lastBlock := tong.Get([]byte(LAST_KEY))
			blockHash,err := block.Serialize()
			if err != nil {
				return nil
			}
			if lastBlock == nil {
				tong.Put(block.Hash,blockHash)
				tong.Put([]byte(LAST_KEY),blockHash)
			}
		}
		//桶为空。表示要新建
		tong,err := tx.CreateBucket([]byte(BUCKET_NAME))
		if err != nil {
			return err
		}
		//把区块存到桶中,并更新最新区块的信息
		blockBytes,err :=block.Serialize()
		tong.Put(block.Hash,blockBytes)

		return nil
	})
	return
}
