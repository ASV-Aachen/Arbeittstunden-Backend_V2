package setup

import (
	"time"

	gocloak "github.com/ASV-Aachen/ArbeitsstundenDB/modules/gocloak"
	secur "github.com/ASV-Aachen/ArbeitsstundenDB/utils"
	"github.com/gofiber/fiber/v2"
)

func CheckCookie(cookie fiber.Cookie) (bool, gocloak.User) {
	cookieValue := cookie.Value
	userdata_decrypted, _ := secur.Decrypt(cookieValue)
	userdata, err := secur.DecodeUser(userdata_decrypted)

	if err != nil {
		return false, gocloak.User{}
	}

	// DONE: ABLAUFDATUM EINFÜHREN
	if time.Now().After(userdata.Expires) {
		return false, gocloak.User{}
	}

	return true, userdata
}

func CreateCookie(User gocloak.User) fiber.Cookie {
	User.AddExpireTime()

	encoded, _ := secur.EncodeUser(User)
	encryptedValue, _ := secur.Encrypt(encoded)

	cookie := &fiber.Cookie{
		Name:    "ArbeitsstundenCookie",
		Value:   encryptedValue,
		Secure:  true,
		Expires: time.Now().Add(48 * time.Hour),
	}

	return *cookie
}
