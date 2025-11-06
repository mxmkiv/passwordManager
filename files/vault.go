package files

import (
	"bufio"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
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
	if index == -1 {
		return text
	} else {
		result.WriteString(text[:index])
		line := text[index : index+len(substr)]
		result.WriteString(green(line))
		result.WriteString(text[index+len(substr):])
	}

	return result.String()
}

func deleteConfirm(inp *bufio.Scanner) int {

Choose:
	for {

		fmt.Print("Удалить ([y]/n)?: ")
		inp.Scan()
		choose := inp.Text()

		switch choose {
		case "y", "":
			return 1
		case "n":
			return 0
		default:
			continue Choose
		}
	}
}

func chooseForDelete(inp *bufio.Scanner, acc []Account, url string) int {

	ShowData(acc, url)

	var res int

	flag := false
	for flag == false {
		fmt.Print("Выберите аккаунт для удаления (отмена - n): ")
		inp.Scan()
		line := inp.Text()

		if line == "n" {
			res = -1
			flag = true
		} else {
			numb, err := strconv.Atoi(line)
			if err != nil {
				color.Red("Некорректный выбор")
			}

			if numb > 0 && numb < len(acc)+1 {
				res = numb - 1
				flag = true
			} else {
				color.Red("Некорректный выбор")
			}
		}
	}

	return res
}

func (vault *Vault) DeleteAccount(inp *bufio.Scanner, url string) (bool, string) {

	var notMatched []Account
	var matched []Account

	deleteFlag := false
	msg := "Не найдено"

	// поиск аккаунтов
	for _, elem := range vault.Accounts {
		if strings.Contains(strings.ToLower(elem.Url), strings.ToLower(url)) {
			matched = append(matched, elem)
			deleteFlag = true
		} else {
			notMatched = append(notMatched, elem)
		}
	}

	if len(matched) > 1 { // если подходящих аккаунтов несколько
		indexToDelete := chooseForDelete(inp, matched, url)
		if indexToDelete == -1 { // отмена удаления
			msg = "Отмена"
			deleteFlag = false
		} else { // удаление
			afterDelete := slices.Delete(matched, indexToDelete, indexToDelete+1)
			vault.Accounts = append(notMatched, afterDelete...)
			msg = "Успешно удалено"
		}
	} else if len(matched) == 1 { // если один
		ShowData(matched, url)
		result := deleteConfirm(inp)
		if result == 1 { // подтвержденое удаления
			vault.Accounts = notMatched
			msg = "Успешно удалено"
		} else { // отмена удаления
			deleteFlag = false
			msg = "Отмена"
		}
	}

	if deleteFlag { // если удаление произошло перезаписываем
		vault.saveData()
	}

	return deleteFlag, msg

}

func FindAccount(vault *Vault, url string) []Account {

	var searchResults []Account
	for _, elem := range vault.Accounts {
		if strings.Contains(strings.ToLower(elem.Url), strings.ToLower(url)) {
			searchResults = append(searchResults, elem)
		}
	}

	return searchResults

}

func ShowData(searchResults []Account, url string) {

	if len(searchResults) != 0 {

		fmt.Printf("Результаты посика по %q: \n\n", url)
		const separator = "──────────────────────────────"

		for i, acc := range searchResults {
			fmt.Println(separator)
			colorString := highlight(acc.Url, url)

			fmt.Printf("%v URL:    %s\n", i+1, colorString)
			fmt.Printf("  Логин:  %s\n", acc.Login)
			fmt.Printf("  Пароль: %s\n", acc.Password)
		}

		fmt.Printf("%s\n\n", separator)

	} else {
		color.Red("Нет записей для данного сервиса")
	}

}

func (vault *Vault) AddAccount(acc *Account) {
	vault.Accounts = append(vault.Accounts, *acc)

	vault.saveData()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) saveData() {

	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	WriteData(data, "data.json")

}
