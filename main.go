package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"passwordLoger/files"
)

type account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

// ////////// account methods ////////////////////////
func (acc *account) showAll() {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *account) toBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

/////////////////////////////////////////////////////

func main() {

	login := getData("Введите логин: ")
	password := getData("Введите пароль: ")
	url := getData("Введите url: ")

	acc, err := creatAccount(login, password, url)
	if err != nil {
		fmt.Print(err)
	}

	file, err := acc.toBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	acc.showAll()
	files.WriteData(file, "data.json")
}

func creatAccount(login string, password string, url string) (*account, error) {

	////check empty login

	if password == "" {
		password = files.GeneratePassword()
	}

	acc := account{
		Login:    login,
		Password: password,
		Url:      url,
	}

	return &acc, nil

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
