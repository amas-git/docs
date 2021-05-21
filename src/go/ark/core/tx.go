package core

// li://bitrue/xrp
type Tx struct {
	Stat
	Namespace string
	Amount    int
	From      string
	To        string
}
