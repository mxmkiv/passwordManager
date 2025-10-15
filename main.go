package main

import (
	"fmt"
	"passwordLoger/files"
)

func main() {

Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			files.CreatAccount()
		case 2:
		case 3:
		default:
			break Menu
		}
	}

}

func getMenu() int {
	var variant int
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)

	return variant
}
