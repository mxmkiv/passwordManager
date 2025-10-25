package files

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts []Account `json:"Accounts"`
}

func NewVault() *Vault {
	file, err := ReadData("data.json")
	if err != nil {
		return &Vault{
			Accounts: []Account{},
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts: []Account{},
		}
	}
	return &vault
}

func highlight(text, substr string) string {

	textLower := strings.ToLower(text)
	substrLower := strings.ToLower(substr)
	green := color.New(color.FgGreen).SprintFunc()

	var result strings.Builder

	index := strings.Index(textLower, substrLower)
	result.WriteString(textLower[:index])
	result.WriteString(green(substrLower))
	result.WriteString(textLower[index+len(substr):])

	return result.String()
}

func FindAccount(inp *bufio.Scanner) {
	vault := NewVault()
	text := "Введите url: "
	url := getData(inp, text)

	var searchReaults []Account
	for _, elem := range vault.Accounts {
		if strings.Contains(elem.Url, url) {
			searchReaults = append(searchReaults, elem)
		}
	}

	if len(searchReaults) != 0 {

		fmt.Printf("Результаты посика по %q: \n\n", url)
		const separator = "──────────────────────────────"

		for _, acc := range searchReaults {
			fmt.Println(separator)
			fmt.Printf("URL:    %s\n", highlight(acc.Url, url))
			fmt.Printf("Логин:  %s\n", acc.Login)
			fmt.Printf("Пароль: %s\n", acc.Password)
		}

		fmt.Println(separator)

	} else {
		color.Red("Нет записей для данного сервиса")
	}

}

func (vault *Vault) AddAccount(acc *Account) {
	vault.Accounts = append(vault.Accounts, *acc)
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	WriteData(data, "data.json")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
