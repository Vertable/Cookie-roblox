package constants

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

var (
	Accounts    = []RobloxAccount{}
	accountFile = "accounts.txt"
	Working     = map[int64]bool{}
)

type jsonResponse struct {
	UserId int64  `json:"UserId"`
	Name   string `json:"Name"`
}

func UnmarshalResponse(response io.Reader) (jsonResponse, error) {
	var playerData jsonResponse
	jsonErr := json.NewDecoder(response).Decode(&playerData)

	return playerData, jsonErr
}

type RobloxAccount struct {
	Cookie string `json:"cookie"`
	Proxy  string `json:"proxy"`
}

func unmarshalAccount(line []byte) (RobloxAccount, error) {
	var acc RobloxAccount
	jsonErr := json.Unmarshal(line, &acc)
	return acc, jsonErr
}

func LoadAccounts() {
	file, err := os.Open(accountFile)
	if err != nil {
		log.Panic(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		account, accountLoadError := unmarshalAccount(scanner.Bytes())
		if accountLoadError == nil {
			if !strings.Contains(account.Proxy, "http://") {
				account.Proxy = "http://" + account.Proxy
			}
			Accounts = append(Accounts, account)
		}
	}
}
