package useraccounts

import (
	"github.com/anhminh10a2hoa/bunny-social-media/helpers"
	"github.com/anhminh10a2hoa/bunny-social-media/interfaces"
)

// UpdateAccount Create function update account
func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)
	defer db.Close()
}
