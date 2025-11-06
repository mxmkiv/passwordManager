package files

import (
	"bufio"
	"fmt"
)

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

func CreateAccount(vault *Vault, inp *bufio.Scanner) {

	login := GetData(inp, "Введите логин: ")
	password := GetData(inp, "Введите пароль: ")
	url := GetData(inp, "Введите url: ")

	if password == "" {
		password = GeneratePassword()
	} else {
		// ("оценка пароля")
	}

	acc := Account{
		Login:    login,
		Password: password,
		Url:      url,
	}

	vault.AddAccount(&acc)
}

func GetData(inp *bufio.Scanner, txt string) string {

	fmt.Print(txt)

	if inp.Scan() {
		return inp.Text()
	}

	return ""
}
