package ethereum

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/polymorph-metadata/app/domain/txreceipt"
)

type EthereumClient struct {
	Client *ethclient.Client
}

func (ec *EthereumClient) senderFromTransaction(transaction *types.Transaction) (*common.Address, error) {
	chainID, err := ec.Client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	msg, err := transaction.AsMessage(types.NewEIP155Signer(chainID))

	if err != nil {
		return nil, err
	}

	sender := msg.From()

	return &sender, nil
}

func (ec *EthereumClient) GetTransactionInfo(txHash string) (result *txreceipt.TxReceipt, err error) {

	hash := common.HexToHash(txHash)
	transaction, isPending, err := ec.Client.TransactionByHash(context.Background(), hash)

	if err != nil {
		if errors.Is(ethereum.NotFound, err) {
			e := errors.New("Transaction Not Found")
			return nil, e
		}
		return nil, err
	}

	if isPending {
		return nil, errors.New("Transaction is not mined yet")
	}

	receipt, err := ec.Client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, err
	}

	sender, err := ec.senderFromTransaction(transaction)
	if err != nil {
		return nil, err
	}

	result = &txreceipt.TxReceipt{
		Hash:              txHash,
		From:              sender.Hex(),
		To:                transaction.To().Hex(),
		Status:            receipt.Status,
		Data:              common.Bytes2Hex(transaction.Data()),
		CumulativeGasUsed: receipt.CumulativeGasUsed,
		GasLimit:          transaction.Gas(),
		GasPrice:          transaction.GasPrice().Uint64(),
		Value:             transaction.Value().Uint64(),
		Nonce:             transaction.Nonce(),
	}

	return result, nil
}

func NewEthereumClient(nodeURL string) (*EthereumClient, error) {
	client, err := ethclient.Dial(nodeURL)

	if err != nil {
		return nil, err
	}

	return &EthereumClient{
		Client: client,
	}, nil
}
