package main

import (
	"bufio"
	"fmt"
	"os"
	"passwordLoger/files"
	"strconv"

	"github.com/fatih/color"
)

func main() {

	vault := files.NewVault()
	scanner := bufio.NewScanner(os.Stdin)

	const (
		CreateAccount = 1
		FindAccount   = 2
		DeleteAccount = 3
		Exit          = 4
	)

Menu:
	for {
		variant := getMenu(scanner)
		switch variant {
		case CreateAccount:
			files.CreateAccount(vault, scanner)
		case FindAccount:
			text := "Введите url: "
			url := files.GetData(scanner, text)
			result := files.FindAccount(vault, url)
			files.ShowData(result, url)
		case DeleteAccount:
			text := "Введите url: "
			url := files.GetData(scanner, text)
			result, msg := vault.DeleteAccount(scanner, url)
			if result {
				color.Green(msg)
			} else {
				color.Red(msg)
			}
		case Exit:
			break Menu
		default:
			color.Blue("Введите число от 1 до 4 ")
			continue Menu
		}
	}

}

func getMenu(inp *bufio.Scanner) int {
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")

	if inp.Scan() {
		n, err := strconv.Atoi(inp.Text())
		if err != nil {
			color.Red("Ошибка ввода")
		}
		return n
	}

	return 0

}
