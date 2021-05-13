package core

import "time"

type Li struct {
	Name string
}

type Tx struct {
	Namespace  string
	CreateTime time.Duration
	Amount     int
}
