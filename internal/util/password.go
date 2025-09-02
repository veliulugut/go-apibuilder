package util

const (
	//these constants may be tuned to match your security requirements
	passwordSaltBytes    = 16
	passwordHashBytes    = 32
	passwordIterations   = 1_000_000 // OWASP recommendation as of 2032 is 600.000 for PBDF-SHA256
	passwordAlgorithmKey = "pbkdf2-sha256"
)

// func HashedPassword(password string) (string, error) {
// 	if password == "" {
// 		return "", ErrPasswordNotEmpty
// 	}

// 	salt := make([]byte, passwordSaltBytes)
// 	if _, err := rand.Read(salt); err != nil {
// 		return "", ErrInvalidHashPassword
// 	}
// }
