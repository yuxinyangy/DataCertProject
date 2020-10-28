package blockchain

import (
	"DataCertProject/util"
	"bytes"
	"crypto/sha256"
	"math/big"
)

const DIFFFICULTY  = 16

/**
 *工作量证明结构体
 */

type ProofOfWork struct {
	//目标值
	Target *big.Int
	Block Block
}

/*
实例化一个pow算法实例
 */
func NewPoW(block Block) ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target,255-DIFFFICULTY)
	pow := ProofOfWork{
		Target:target,
		Block:block,
	}
	return pow
}

func (p ProofOfWork) Run()( []byte,int64){
	var nonce int64
	bigBlock := new(big.Int)
	var block256Hash []byte
	for {
		block :=p.Block
		heightBytes,_ :=util.IntToBytes(block.Height)
		timeBytes,_ := util.IntToBytes(block.TimeStamp)
		versionBytes := util.StringToBytes(block.Version)
		nonceBytes,_ := util.IntToBytes(nonce)
		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
		},[]byte{})
		sha256Hash := sha256.New()
		sha256Hash.Write(blockBytes)
		block256Hash = sha256Hash.Sum(nil)

		//fmt.Printf("挖矿中，当前尝试nonce值:%d\n",nonce)
		bigBlock = bigBlock.SetBytes(block256Hash)
		//fmt.Printf("目标值: %x\n",p.Target)
		//fmt.Printf("hash值:%x\n",bigBlock)
		if p.Target.Cmp(bigBlock)==1{
			break
		}
		nonce++
	}
	return block256Hash,nonce
}
