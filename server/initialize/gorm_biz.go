package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gupage"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(gupage.CliUser{}, gupage.Tree{}, gupage.CliMain{}, gupage.CliWallet{}, gupage.CliNotice{})
	if err != nil {
		return err
	}
	return nil
}
