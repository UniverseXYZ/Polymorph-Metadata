package txreceipt

type TxReceipt struct {
	Hash              string `json:"hash" bson:"hash"`
	From              string `json:"from" bson:"from"`
	To                string `json:"to" bson:"to"`
	Status            uint64 `json:"status" bson:"status"`
	Data              string `json:"data" bson:"data"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	GasLimit          uint64 `json:"gasLimit" bson:"gasLimit"`
	GasPrice          uint64 `json:"gasPrice" bson:"gasPrice"`
	Value             uint64 `json:"value" bson:"value"`
	Nonce             uint64 `json:"nonce" bson:"nonce"`
}
