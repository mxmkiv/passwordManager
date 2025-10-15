package files

import (
	"encoding/json"

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
