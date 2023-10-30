package main

import (
	"fmt"
	"net/url"

	"github.com/skip2/go-qrcode"
)

func main() {
	// OTP 密鑰
	secret := []string{
		"VNKNYOWI65RAGQVUNTN37DBPZH2U5XCL",
		"4QQFH5OXP4AQRS5ZKETINQ64WUBB6GG6",
		"D3VOMETPLPHLILMG674YECSTOR5YYB5Y",
		"SIJLL6MZ6D5JKCW325AOCTZU6URBGRGZ",
	}

	// 使用密鑰生成 TOTP 實例
	account := []string{
		"Finance 财务",
		"Login 登录",
		"Operations 运营",
		"Order 订单",
	}
	issuer := "Super365-BR"

	for i := 0; i < 4; i++ {
		// 產生 OTP 認證的 URL
		URL, err := url.Parse("otpauth://totp")
		if err != nil {
			panic(err)
		}
		URL.Path += fmt.Sprintf("/%s:%s", url.PathEscape(issuer), url.PathEscape(account[i]))

		// params := url.Values{}
		// params.Add("secret", secret[i])
		// params.Add("issuer", issuer)

		// URL.RawQuery = params.Encode()

		fmt.Println(URL.String())

		url := fmt.Sprintf("otpauth://totp/%s?issuer=%s&secret=%s", account[i], issuer, secret[i])

		// 使用 go-qrcode 套件產生 QR Code 圖片
		err = qrcode.WriteFile(url, qrcode.Medium, 256, fmt.Sprintf("%s-%s-qr-code.png", issuer, account[i]))
		if err != nil {
			panic(err)
		}

		fmt.Printf("Created QR code for user %s with issuer %s.\n", account[i], issuer)

	}

}
