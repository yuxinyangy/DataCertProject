package blockchain

import "math/big"

/**
 *工作量证明结构体
 */

type ProofOfWork struct {
	//目标值
	Target big.Int
}
