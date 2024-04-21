package bcrpyt

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plaintext password using bcrypt with the given cost factor.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// ValidatePassword compares a plaintext password with a stored hashed password.
func ValidatePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
