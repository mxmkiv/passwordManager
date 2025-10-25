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

func (acc *Account) DeleteAccount() {

}

func CreatAccount(inp *bufio.Scanner) {

	login := getData(inp, "Введите логин: ")
	password := getData(inp, "Введите пароль: ")
	url := getData(inp, "Введите url: ")

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

	vault := NewVault()
	vault.AddAccount(&acc)
}

func getData(inp *bufio.Scanner, txt string) string {

	fmt.Print(txt)

	if inp.Scan() {
		return inp.Text()
	}

	return ""
}
