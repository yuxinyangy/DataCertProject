package util

import (
	"DataCertProject/blockchain"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

func MD5HashString(data string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	passwordBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(passwordBytes)
}

func MD5HashReader(reader io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	md5Hash := md5.New()
	md5Hash.Write(bytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

func SHA256HashBlock(block blockchain.Block) ([]byte) {
	//1、对block字段进行拼接

	//2、对拼接后的数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(""))
	return sha256Hash.Sum(nil)
}
