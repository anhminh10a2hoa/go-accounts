package main

import (
	"github.com/anhminh10a2hoa/bunny-social-media/api"
	"github.com/anhminh10a2hoa/bunny-social-media/database"
)

func main() {
	// migrations.Migrate()
	// migrations.MigrateTransactions()
	database.InitDatabase()
	api.StartApi()
}
