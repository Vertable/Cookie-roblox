package client

import (
	"fmt"
	"github.com/Vertable/Cookie-roblox/constants"
	"log"
	"net/http"
	"net/url"
	"time"
)

func CheckAccount(testAccount constants.RobloxAccount) {
	proxyURL, err := url.Parse(testAccount.Proxy)
	if err != nil {
		log.Println(err)
		return
	}
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
		Timeout:   0x40 * time.Second,
	}
	req, _ := http.NewRequest("GET", "https://www.roblox.com/my/settings/json", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	req.Header.Add("cookie", fmt.Sprintf(".ROBLOSECURITY=%s;", testAccount.Cookie))

	resp, err := client.Do(req)
	if err == nil {
		if resp.StatusCode == 0xc8 {
			userResponse, decodeError := constants.UnmarshalResponse(resp.Body)
			if decodeError == nil {
				constants.Working[userResponse.UserId] = true
				fmt.Printf("You have %d valid unique accounts.\n", len(constants.Working))
			}
		}
		if resp.StatusCode == 0x1AD {
			time.Sleep(time.Second)
			CheckAccount(testAccount)
		}
	}

}
