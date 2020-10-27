package blockchain

import "github.com/bolt-master"

//桶的名称，该桶用于装区块信息
var BUCKET_NAME  = "blocks"
//表示最新的区块的key名
var LAST_KEY = "lasthash"
/*
区块链结构体实例:用于描述或表示代表一条区块链
* 该条区块链包括以下功能:
*           ①将新产生的区块与已有的区块链接起来，并保存
            ②可以查询某个区块的信息
            ③可以将所有区块进行遍历，输出区块信息
 */
type BlockChain struct {
	LastHash []byte//最新区块的hash
	BoltDb *bolt.DB
}



func NewBlockChain() BlockChain  {
	//1.创建创世区块
	genesis :=CreateGenesisBlock()//创世区块
	//2.创建一个存储区块数据的文件
	db,err:=bolt.Open("chain.db",0600,nil)
	if err != nil {
		panic(err.Error())
	}
	bl := BlockChain{
		LastHash: genesis.Hash,
		BoltDb:   db,
	}
	//3.把新创建的创世区块存入到chain.db当中的一个桶中
	db.Update(func(tx *bolt.Tx) error {
		bucket,err :=tx.CreateBucket([]byte(BUCKET_NAME))
		if err != nil {
			panic(err.Error())
		}
		serialBlock,err := genesis.Serialize()
		if err != nil {
			panic(err.Error())
		}
		//把创世区块存入到桶中
		bucket.Put(genesis.Hash,serialBlock)
		//更新指向最新区块的Hash值
		bucket.Put([]byte(LAST_KEY),genesis.Hash)
		return nil
	})
	return bl
}

/*
调用BlockChain的该SaveBlock方法，该方法可以将一个生成的新区块块保存到chain.db文件中
 */
func (bc BlockChain) SaveData(data []byte)  {
	db := bc.BoltDb
	var lastBlock *Block
	//先查询chain.db中存储的最新的区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("boltdb未创建，请重试！")
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		lastBlockBytes := bucket.Get(lastHash)
		lastBlock,_ =DeSerialize(lastBlockBytes)
		return nil
	})
	//先生成一个区块，把data存入新生成的区块中
	newBlock := NewBlock(lastBlock.Height+1,data,lastBlock.Hash)
	//更新chain.db
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("boltdb未创建，请重试！")
		}
		blockBytes,err :=newBlock.Serialize()
		if err != nil {
			return nil
		}
		bucket.Put(data,blockBytes)
		return nil
	})
	return
}
