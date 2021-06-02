package migrations

import (
	"github.com/anhminh10a2hoa/bunny-social-media/helpers"
	"github.com/anhminh10a2hoa/bunny-social-media/interfaces"
)

func createAccounts() {
	db := helpers.ConnectDB()

	users := &[2]interfaces.User{
		{
			Username: "admin", Email: "admin@gmail.com",
		},
		{
			Username: "anhminh10a2hoa", Email: "anhminh10a2hoa@gmail",
		},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := interfaces.Account{Type: "admin", Name: string(users[i].Username + "'s account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}

		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(User, Account)
	defer db.Close()

	createAccounts()
}
