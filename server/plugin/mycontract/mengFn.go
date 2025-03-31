package mycontract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

// 0xcF888cBc43b7976f179B7a0f7542eB1202659714
const contract = "0x58f6db1ca9Dc6E0F9b42D72491bfC9Bf940116b1"

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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = big.NewInt(5e9) // 3 Gwei in Wei
	return auth
}

// ServerRigisterNew
func (mengWeb3 *MengWeb3) SeverRigisterNew(list []MengOperationRecord) (string, error) {
	auth := mengWeb3.Auth()
	tx, err := mengWeb3.ContractInstance.ServerRegisterNew(auth, list)
	if err != nil {
		return "", err
	}
	fmt.Println("激活地址", tx.Hash().Hex())
	receipt, err := WaitForTxConfirmation(mengWeb3.client, tx.Hash())
	if err != nil {
		return "", err
	}
	fmt.Println("激活结果", receipt)
	return tx.Hash().Hex(), nil
}

// 升级
func (mengWeb3 *MengWeb3) CliUpdate(list []MengOperationRecord) (string, error) {
	auth := mengWeb3.Auth()
	tx, err := mengWeb3.ContractInstance.CliUpdate(auth, list)
	if err != nil {
		return "", err
	}
	receipt, err := WaitForTxConfirmation(mengWeb3.client, tx.Hash())
	if err != nil {
		return "", err
	}
	fmt.Println("升级结果", receipt)
	return tx.Hash().Hex(), nil
}

type WithdrawParams struct {
	FromAddress string   // 提取来源地址
	Amount      *big.Int // 提取金额
	GasLimit    uint64   // 自定义 Gas 限制
	GasPrice    *big.Int // 自定义 Gas 价格
}

// 提取某个地址金额
func (mengWeb3 *MengWeb3) Withdraw(params WithdrawParams) (map[string]interface{}, error) {

	// 转换地址格式
	fromAddress := common.HexToAddress(params.FromAddress)

	// 生成验证码（与合约相同的算法）
	currentTimestamp := time.Now().Unix()
	truncated := currentTimestamp / 1000
	verificationCode := new(big.Int).Mul(
		big.NewInt(truncated),
		big.NewInt(truncated),
	)

	// 发送交易
	tx, err := mengWeb3.ContractInstance.WithdrawFrom(
		mengWeb3.Auth(),
		fromAddress,
		params.Amount,
		verificationCode,
	)
	if err != nil {
		return nil, fmt.Errorf("transaction failed: %v", err)
	}

	// 等待确认
	receipt, err := WaitForTxConfirmation(mengWeb3.client, tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("confirmation failed: %v", err)
	}

	fmt.Println("Transaction successful!", receipt)

	// 返回完整结果
	return map[string]interface{}{
		"txHash":      tx.Hash().Hex(),
		"gasUsed":     receipt.GasUsed,
		"blockNumber": receipt.BlockNumber,
		"status":      receipt.Status,
	}, nil
}

// 添加这个函数到你的代码中以检查交易是否成功
//func WaitForTxConfirmation(client *ethclient.Client, txHash common.Hash) error {
//	ctx := context.Background()
//
//	// 循环检查交易是否被打包
//	for {
//		fmt.Println("Transaction  Waiting...")
//		time.Sleep(3 * time.Second)
//		receipt, err := client.TransactionReceipt(ctx, txHash)
//		if err == nil {
//			// 检查交易状态
//			if receipt.Status == types.ReceiptStatusSuccessful {
//				fmt.Println("Transaction successful!")
//				return nil
//			} else {
//				fmt.Println("Transaction failed!")
//				return errors.New("Transaction failed")
//			}
//
//		}
//		if err != nil {
//			if err.Error() == "not found" {
//				// 交易可能还未被打包，稍后再试
//				fmt.Println("Transaction not found yet. Waiting...")
//				time.Sleep(3 * time.Second)
//				continue
//			}
//			// 其他错误，返回
//			return err
//		}
//		// 如果交易还没被打包，稍后再试
//		fmt.Println("Waiting for transaction to be mined...")
//	}
//}

// 增强版交易等待函数
func WaitForTxConfirmation(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			receipt, err := client.TransactionReceipt(ctx, txHash)
			if err == nil {
				return receipt, nil
			}
			if err != ethereum.NotFound {
				return nil, err
			}
		}
	}
}
