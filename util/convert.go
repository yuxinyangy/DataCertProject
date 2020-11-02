package util

import (
	"bytes"
	"encoding/binary"
)

/*
*int64转换为[]byte
 */
func IntToBytes(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

/*
*string字符串转换为[]byte
 */
func StringToBytes(st string) []byte {
	return []byte(st)
}
