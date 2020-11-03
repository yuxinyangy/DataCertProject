package blockchain

import (
	"DataCertProject/models"
	"errors"
	"github.com/bolt-master"
	"math/big"
)

//桶的名称，该桶用于装区块信息
var BUCKET_NAME = "blocks"

//表示最新的区块的key名
var LAST_KEY = "lasthash"

var CHAINDB = "chain.db"

var CHAIN BlockChain

/*
区块链结构体实例:用于描述或表示代表一条区块链
* 该条区块链包括以下功能:
*           ①将新产生的区块与已有的区块链接起来，并保存
            ②可以查询某个区块的信息
            ③可以将所有区块进行遍历，输出区块信息
*/
type BlockChain struct {
	LastHash []byte //最新区块的hash
	BoltDb   *bolt.DB
}

/*
查询所有的区块信息，并返回。将所有的区块放入到切片中
*/

func (bc BlockChain) QueryAllBlocks() []*Block {
	blocks := make([]*Block, 0)
	db := bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据出错")
		}
		eachKey := bc.LastHash
		preHashBig := new(big.Int)
		zeroBig := big.NewInt(0) //0的大整数
		for {
			eachBlockBytes := bucket.Get(eachKey)
			//反序列化后得到的每一个区块
			eachBlock, _ := DeSerialize(eachBlockBytes)
			//将遍历到的每一个区块链结构体指针放入到[]byte容器中
			blocks = append(blocks, eachBlock)

			preHashBig.SetBytes(eachBlock.PrevHash)
			if preHashBig.Cmp(zeroBig) == 0 { //通过if条件语句判断区块链遍历是否已到创世区块，如果到创世区块，跳出循环
				break
			}
			//否则，继续向前遍历
			eachKey = eachBlock.PrevHash
		}
		return nil
	})
	return blocks
}

/*
通过区块的高度查询某个具体的区块，返回区块实例
*/
func (bc BlockChain) QueryBlockHeight(height int64) *Block {
	if height < 0 { //如果目标高度小于0，则说明参数不合法
		return nil
	}
	var block *Block
	db := bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据失败")
		}
		hashKey := bc.LastHash
		for {
			lastBlockBytes := bucket.Get(hashKey)
			eachBlock, _ := DeSerialize(lastBlockBytes)
			if eachBlock.Height < height { //给定的数字超出区块的高度
				break
			}
			if eachBlock.Height == height { //高度和目标一致，已经找到目标区块，结束循环
				block = eachBlock
				break
			}
			//遍历的当前的区块的高度与目标高度不一致，继续往前遍历
			//以eachBlock.PrevHash为key，使用Get获取上一个区块的数据
			hashKey = eachBlock.PrevHash
		}
		return nil
	})
	return block
}

func NewBlockChain() BlockChain {
	//0.打开存储区块数据的chain.db文件
	db, err := bolt.Open(CHAINDB, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	var bl BlockChain
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		if len(lastHash) == 0 { //无创世区块
			//1.创建创世区块
			genesis := CreateGenesisBlock() //创世区块
			//2.创建一个存储区块数据的文件
			bl = BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
			genesisBytes, _ := genesis.Serialize()
			bucket.Put(genesis.Hash, genesisBytes)
			bucket.Put([]byte(LAST_KEY), genesis.Hash)
		} else { //有创世区块
			lastHash := bucket.Get([]byte(LAST_KEY))
			lastBlockBytes := bucket.Get(lastHash)
			lastBlock, err := DeSerialize(lastBlockBytes)
			if err != nil {
				panic("读取区块链数据失败")
			}
			bl = BlockChain{
				LastHash: lastBlock.Hash,
				BoltDb:   db,
			}
		}
		return nil
	})
	//为全局变量赋值
	CHAIN = bl
	return bl
}

func (bc BlockChain) QueryBlockByCertId(cert_id []byte) (*Block, error) {
	var block *Block
	db := bc.BoltDb
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("查询区块数据遇到错误")
		}
		//桶存在
		eachHash := bucket.Get([]byte(LAST_KEY))
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0)
		var certRecord *models.CertRecord
		for {
			eachBlockBytes := bucket.Get(eachHash)
			eachBlock, _ := DeSerialize(eachBlockBytes)
			//序列化以后的结构体数据certRecord 类型:eachBlock.Data
			certRecord,_ =models.DeSerializeRecord(eachBlock.Data)
			//找到的情况
			if string(certRecord.CertId) == string(cert_id) {
				block = eachBlock
				break
			}
			//找不到的情况
			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig) == 0 {
				break
			}
		}
		return nil
	})

	return block, err
}

/*
调用BlockChain的该SaveBlock方法，该方法可以将一个生成的新区块块保存到chain.db文件中
*/
func (bc BlockChain) SaveData(data []byte) (Block, error) {
	db := bc.BoltDb
	var e error
	var lastBlock *Block
	//先查询chain.db中存储的最新的区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			e := errors.New("boltdb未创建，请重试！")
			return e
		}
		lastBlockBytes := bucket.Get(bc.LastHash)
		lastBlock, _ = DeSerialize(lastBlockBytes)
		return nil
	})
	//先生成一个区块，把data存入新生成的区块中
	newBlock := NewBlock(lastBlock.Height+1, data, lastBlock.Hash)
	//更新chain.db
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//区块序列化
		newBlockBytes, _ := newBlock.Serialize()
		//把区块信息保存到boltdb中
		bucket.Put(newBlock.Hash, newBlockBytes)
		//更新代表最后一个区块hash值的记录
		bucket.Put([]byte(LAST_KEY), newBlock.Hash)
		bc.LastHash = newBlock.Hash
		return nil
	})
	return newBlock, e
}
