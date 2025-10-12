package main

import (
	"encoding/json"
	"fmt"
	"passwordLoger/files"
)

type account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

// //// account methods
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

////////

func main() {

	login := getData("Введите логин: ")
	password := getData("Введите пароль: ")
	url := getData("Введите url: ")

	acc := account{
		Login:    login,
		Password: password,
		Url:      url,
	}

	file, err := acc.toBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	acc.showAll()
	files.WriteData(file, "data.json")
}

func getData(txt string) string {
	fmt.Print(txt)
	var data string
	fmt.Scan(&data)
	return data
}
