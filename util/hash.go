package util

import (
	"DataCertProject/blockchain"
	"crypto/sha256"
)

func SHA256HashBlock(block blockchain.Block) ([]byte) {
	//1、对block字段进行拼接

	//2、对拼接后的数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(""))
	return sha256Hash.Sum(nil)
}
