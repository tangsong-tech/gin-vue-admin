package mycontract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chain/online"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/chain/wallet"
	"github.com/shopspring/decimal"
	"testing"
)

//0x88EA65Ce12BB49C4385424Eb0324F18AbCbC126F  //部署地址

func TestClient(t *testing.T) {
	client, err := online.NewClient()
	if err != nil {
		t.Error(err)
		return
	}
	keystore, err := wallet.ImportKeyStore("0x39999756E2FF1D9b7d6592510b4E9bD08851D8B1")
	if err != nil {
		t.Error(err)
		return
	}
	mengWeb3, err := NewMengWeb3(keystore, client)
	if err != nil {
		t.Error(err)
		return
	}
	balance, _ := client.BalanceAt(context.Background(), common.HexToAddress("0x27B31CdFFF4A49397E39C5F392a8F477A2aE1f87"), nil)
	fmt.Println(online.WeiToDecimal(balance))
	bnb20Balance, err := online.Bnb20Balance(client, "0x55d398326f99059fF775485246999027B3197955", "0x27B31CdFFF4A49397E39C5F392a8F477A2aE1f87")
	if err != nil {
		return
	}
	fmt.Println(bnb20Balance)
	err = online.GetApproval(client, "0x27B31CdFFF4A49397E39C5F392a8F477A2aE1f87", "0x58f6db1ca9Dc6E0F9b42D72491bfC9Bf940116b1")
	if err != nil {
		return
	}
	var list []MengOperationRecord
	num := decimal.NewFromFloat(1)
	list = append(list, MengOperationRecord{
		RelatedAddress: common.HexToAddress("0x27B31CdFFF4A49397E39C5F392a8F477A2aE1f87"),
		Amount:         online.DecimalToWei(&num),
	})
	num1 := decimal.NewFromFloat(0.0002)
	list = append(list, MengOperationRecord{
		RelatedAddress: common.HexToAddress("0x35DdBaeA8a78d37329Ca4bb574C410D5D2f464f6"),
		Amount:         online.DecimalToWei(&num1),
	})
	num2 := decimal.NewFromFloat(0.0001)
	list = append(list, MengOperationRecord{
		RelatedAddress: common.HexToAddress("0x154b8BB871b72C501aE45765d945A16b8659F417"),
		Amount:         online.DecimalToWei(&num2),
	})
	fmt.Println(list)

	rigister, err := mengWeb3.SeverRigisterNew(list)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rigister)

}

// 提取测试
func TestWithdraw(t *testing.T) {
	client, err := online.NewClient()
	if err != nil {
		t.Error(err)
		return
	}
	keystore, err := wallet.ImportKeyStore("0x39999756E2FF1D9b7d6592510b4E9bD08851D8B1")
	if err != nil {
		t.Error(err)
		return
	}
	mengWeb3, err := NewMengWeb3(keystore, client)
	if err != nil {
		t.Error(err)
		return
	}
	var withdrawParams WithdrawParams
	num := decimal.NewFromFloat(0.01)
	withdrawParams.Amount = online.DecimalToWei(&num)
	withdrawParams.FromAddress = "0x27B31CdFFF4A49397E39C5F392a8F477A2aE1f87"
	withdraw, err := mengWeb3.Withdraw(withdrawParams)
	if err != nil {
		return
	}
	fmt.Println(withdraw)
}
