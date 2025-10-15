package files

import (
	"math/rand"
	"strings"
)

/*
	parameters of the generated password
	length = 12
	contains lowercase and uppercase latin characters
	contains numbers
	contains special characters
*/

func GeneratePassword() string {

	var validCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHLJKLMNOPQRSTUVWXYZ1234567890!*-#@"
	const pass_len = 12
	var pass strings.Builder
	for i := 0; i < pass_len; i++ {
		numb := rand.Intn(len(validCharacters))
		pass.WriteString(string(validCharacters[numb]))
	}

	return pass.String()
}

// func СheckEntropy(password string) float64 {

// }
