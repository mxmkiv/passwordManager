package main

import (
	"bufio"
	"fmt"
	"os"
	"passwordLoger/files"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

Menu:
	for {
		variant := getMenu(scanner)
		switch variant {
		case 1:
			files.CreatAccount(scanner)
		case 2:
			files.FindAccount(scanner)
		case 3:
		case 4:
			break Menu
		default:
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
		n, _ := strconv.Atoi(inp.Text())
		return n
	}

	return 0

}
