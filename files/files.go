package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func WriteData(content []byte, name string) {
	dirName := "worktree"
	filepath := dirName + "/" + name
	os.Mkdir(dirName, 0755)
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	color.Green("Запись успешна")
}

func ReadData(name string) ([]byte, error) {
	filepath := "worktree/" + name
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
