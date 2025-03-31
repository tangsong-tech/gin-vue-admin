package online

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"math/big"
)

// // BnbBalance 获取账户Bnb余额
func BnbBalance(client *ethclient.Client, address string) (balance *decimal.Decimal, err error) {
	//获取余额
	res, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return nil, err
	}
	//res, err := client.PendingBalanceAt(context.Background(), common.HexToAddress(address))
	//if err != nil {
	//	fmt.Println(err)
	//	//return nil, err
	//}

	balance = WeiToDecimal(res)
	return balance, nil
}

// 获取usdt 金额
func Bnb20Balance(client *ethclient.Client, contract string, address string) (balance *decimal.Decimal, err error) {
	//获取余额
	instance, err := NewGoethereum(common.HexToAddress(contract), client)
	if err != nil {
		return nil, err
	}
	res, err := instance.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	balance = WeiToDecimal(res)
	return balance, nil
}

// 代币授权
func Bsc20IncreaseAllowance(client *ethclient.Client, contract string, keystore *keystore.Key, spender string, amount *decimal.Decimal) error {
	instance, err := NewGoethereum(common.HexToAddress(contract), client)
	if err != nil {
		return err
	}
	res, err := instance.Allowance(nil, keystore.Address, common.HexToAddress(spender))
	if err != nil {
		return err
	}
	if WeiToDecimal(res).Cmp(decimal.NewFromInt(10000)) > 0 {
		fmt.Println("已经授权成功", keystore.Address)
	}

	//获取nonce
	nonce, _ := client.PendingNonceAt(context.Background(), keystore.Address)
	//获取gasPrice
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	chainID, _ := client.NetworkID(context.Background())
	//构造auth
	auth, _ := bind.NewKeyedTransactorWithChainID(keystore.PrivateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(70000) // in units
	auth.GasPrice = gasPrice
	allowance, err := instance.Approve(auth, common.HexToAddress(spender), DecimalToWei(amount))
	if err != nil {
		return err
	}
	fmt.Println("授权成功", allowance.Hash().Hex())
	return nil
}

// 查询授权额度
func GetApproval(client *ethclient.Client, from, spender string) error {
	instance, err := NewGoethereum(common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), client)
	if err != nil {
		return err
	}
	res, err := instance.Allowance(nil, common.HexToAddress(from), common.HexToAddress(spender))
	if err != nil {
		return err
	}
	fmt.Println("授权额度", WeiToDecimal(res))
	if WeiToDecimal(res).Cmp(decimal.NewFromInt(1)) > 0 {
		fmt.Println("已经认证成功", from)
		return nil
	}
	return errors.New("认证额度不足")
}
