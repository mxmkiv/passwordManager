package files

import (
	"bufio"
	"fmt"
	"os"
)

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

func (acc *Account) DeleteAccount() {

}

func CreatAccount() {

	login := getData("Введите логин: ")
	password := getData("Введите пароль: ")
	url := getData("Введите url: ")

	if password == "" {
		password = GeneratePassword()
	} else {
		//fmt.Println("оценка пароля")
	}

	acc := Account{
		Login:    login,
		Password: password,
		Url:      url,
	}

	vault := NewVault()
	vault.AddAccount(&acc)
}

func getData(txt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var data string

	fmt.Print(txt)

	for {
		if scanner.Scan() {
			line := scanner.Text()
			data = line
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка", err)
			break
		}
	}

	return data
}
