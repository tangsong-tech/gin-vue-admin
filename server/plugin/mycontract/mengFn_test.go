package mycontract

import (
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
	keystore, err := wallet.ImportKeyStore("0x88EA65Ce12BB49C4385424Eb0324F18AbCbC126F")
	if err != nil {
		t.Error(err)
		return
	}
	mengWeb3, err := NewMengWeb3(keystore, client)
	if err != nil {
		t.Error(err)
		return
	}
	rigister, err := mengWeb3.SeverRigister("0x9Ba5A5bB563B534BC671Dafe9BbEd80333008d37", "0xAbDc60C4d049c66CdcEa7D7b5F0e371744356946")
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
	keystore, err := wallet.ImportKeyStore("0x88EA65Ce12BB49C4385424Eb0324F18AbCbC126F")
	if err != nil {
		t.Error(err)
		return
	}
	mengWeb3, err := NewMengWeb3(keystore, client)
	if err != nil {
		t.Error(err)
		return
	}
	am := decimal.NewFromFloat(0.1)
	withdraw, err := mengWeb3.SeverWithdraw("0xC9a5ee343aF59812A4B743eD32611223efB509e4", "0xAbDc60C4d049c66CdcEa7D7b5F0e371744356946", &am)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(withdraw)
}
