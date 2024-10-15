package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, secret string) (string, error) {
	// รวมรหัสผ่านกับรหัสลับ
	combined := password + secret
	// log.Print(combined)
	// ใช้ bcrypt ในการแฮชรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(combined), bcrypt.DefaultCost)
	// log.Print(hashedPassword)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}
