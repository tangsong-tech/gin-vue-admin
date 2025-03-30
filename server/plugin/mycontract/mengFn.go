package mycontract

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chain/online"
	"github.com/shopspring/decimal"
	"math/big"
	"time"
)

// 0xcF888cBc43b7976f179B7a0f7542eB1202659714
const contract = "0x57324aba2aaac95a03fc1657f4434741a636ff73"

type MengWeb3 struct {
	client           *ethclient.Client
	keystore         *keystore.Key
	ContractInstance *MycontractTransactor // 新增字段存储智能合约实例
}

func NewMengWeb3(keystore *keystore.Key, client *ethclient.Client) (*MengWeb3, error) {
	mengWeb3 := &MengWeb3{
		client:   client,
		keystore: keystore,
	}
	// 创建合约实例
	contractInstance, err := NewMycontractTransactor(common.HexToAddress(contract), client)
	if err != nil {
		return nil, err
	}
	mengWeb3.ContractInstance = contractInstance
	return mengWeb3, nil
}

// 构造TransactionOpts
func (mengWeb3 *MengWeb3) Auth() *bind.TransactOpts {
	//获取nonce
	nonce, _ := mengWeb3.client.PendingNonceAt(context.Background(), mengWeb3.keystore.Address)
	//获取gasPrice
	//gasPrice, _ := client.SuggestGasPrice(context.Background())

	chainID, _ := mengWeb3.client.NetworkID(context.Background())
	//构造auth
	auth, _ := bind.NewKeyedTransactorWithChainID(mengWeb3.keystore.PrivateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)             // in wei
	auth.GasLimit = uint64(1000000)        // in units
	auth.GasPrice = big.NewInt(3000000000) // 3 Gwei in Wei
	return auth
}

// ServerRigisterNew
func (mengWeb3 *MengWeb3) SeverRigisterNew(amount *big.Int, list []MengOprationRecord, isrigister bool) (string, error) {
	auth := mengWeb3.Auth()
	tx, err := mengWeb3.ContractInstance.ServerRigisterNew(auth, amount, list, isrigister)
	if err != nil {
		return "", err
	}
	err = WaitForTxConfirmation(mengWeb3.client, tx.Hash())
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// 帮助用户链上注册
func (mengWeb3 *MengWeb3) SeverRigister(from string, to string) (string, error) {
	auth := mengWeb3.Auth()
	tx, err := mengWeb3.ContractInstance.SeverRigister(auth, common.HexToAddress(from), common.HexToAddress(to))
	if err != nil {
		return "", err
	}
	fmt.Println(tx.Hash().Hex())
	err = WaitForTxConfirmation(mengWeb3.client, tx.Hash())
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// 服务端提取奖金
func (mengWeb3 *MengWeb3) SeverWithdraw(from string, to string, amount *decimal.Decimal) (string, error) {
	auth := mengWeb3.Auth()
	tx, err := mengWeb3.ContractInstance.SeverFromtowith(auth, common.HexToAddress(from), common.HexToAddress(to), online.DecimalToWei(amount))
	if err != nil {
		return "", err
	}
	err = WaitForTxConfirmation(mengWeb3.client, tx.Hash())
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// 添加这个函数到你的代码中以检查交易是否成功
func WaitForTxConfirmation(client *ethclient.Client, txHash common.Hash) error {
	ctx := context.Background()

	// 循环检查交易是否被打包
	for {
		fmt.Println("Transaction  Waiting...")
		time.Sleep(3 * time.Second)
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err == nil {
			// 检查交易状态
			if receipt.Status == types.ReceiptStatusSuccessful {
				fmt.Println("Transaction successful!")
				return nil
			} else {
				fmt.Println("Transaction failed!")
				return errors.New("Transaction failed")
			}

		}
		if err != nil {
			if err.Error() == "not found" {
				// 交易可能还未被打包，稍后再试
				fmt.Println("Transaction not found yet. Waiting...")
				time.Sleep(3 * time.Second)
				continue
			}
			// 其他错误，返回
			return err
		}
		// 如果交易还没被打包，稍后再试
		fmt.Println("Waiting for transaction to be mined...")
	}
}
