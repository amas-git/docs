package core

import "time"

type Stat struct {
	CreateTime time.Duration
	ModifyTime time.Duration
	Amount     int
	AmountIn   int
	AmountOut  int
}
type Li struct {
	Stat
	Name string
	Txs  []Tx
}

// 遍历所有交易
func ForeachTxs(f func(li *Li) int) int {
	return -1
}

// 获取总数量
func (r *Li) Amount() int {
	return 0
}

func (r *Li) AmountByTag() int {
	return 0
}
