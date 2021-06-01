package users

import (
	"github.com/anhminh10a2hoa/bunny-social-media/helpers"
	"github.com/anhminh10a2hoa/bunny-social-media/interfaces"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func Login(username string, password string) map[string]interfaces{} {
	db := helpers.ConnectDB()
	user := &interfaces.User{}

	if db.Where("username = ?", username).First(&user).RecordNotFound() {
		return map[string]interfaces{}{"message": "User not found"}
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interfaces{}{"message": "Wrong password"}
	}

	accounts := []interfaces.ResponseAccount{}
	db.Table("accounts").Select("id, name, balance").Where("user_id = ? ", user.ID).Scan(&accounts)

		responseUser := &interfaces.ResponseUser{
			ID: user.ID,
			Username: user.Username,
			Email: user.Email,
			Accounts: accounts,
	}

	defer db.Close()

	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry": time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	var response = map[string]interface{}{"message": "all is fine"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}