package database

import (
	"github.com/anhminh10a2hoa/bunny-social-media/helpers"
	"github.com/anhminh10a2hoa/bunny-social-media/utils"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	config, err := utils.LoadConfig(".")
	database, err := gorm.Open(config.DBDRIVER, "host="+config.DBADDRESS+" port="+config.DBPORT+" user="+config.DBUSER+" dbname="+config.DBNAME+" password="+config.DBPASSWORD+" sslmode=disable")

	helpers.HandleErr(err)

	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
