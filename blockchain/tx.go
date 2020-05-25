package blockchain

type TxOutput struct {
	Value  int
	PubKey string // user public key
}

type TxInput struct {
	ID  []byte
	Out int    // associated TxOutput
	Sig string // user signature
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
