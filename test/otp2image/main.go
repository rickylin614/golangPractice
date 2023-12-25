package main

import (
	"fmt"
	"net/url"

	"github.com/skip2/go-qrcode"
)

func main() {
	// OTP 密鑰
	secret := []string{
		"Q2VBMSYKDZY7OCWJRKT5RMOQD6UIWOR7",
		"TT6HK2U3RJJAXSL72NULRTA6ARB4OPMT",
		"M5ZQOELUT3AJQJIHIQJ6ZOJY7PZFLFQZ",
		"6NTZFU6A6FM5FMJFPXME2FWQTE246XZY",
	}

	// 使用密鑰生成 TOTP 實例
	account := []string{
		"Finance 财务",
		"Login 登录",
		"Operations 运营",
		"Order 订单",
	}
	//範例: Super365-VN001越南(Login登錄),Super365-IA印度(Login登錄)
	//品牌名稱-代碼+國家名稱(功能英文+功能中文)
	issuer := "Teen Patti-IN002印度" // 品牌名稱

	for i := 0; i < 4; i++ {
		// 產生 OTP 認證的 URL
		URL, err := url.Parse("otpauth://totp")
		if err != nil {
			panic(err)
		}
		URL.Path += fmt.Sprintf("/%s:%s", url.PathEscape(issuer), url.PathEscape(account[i]))

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
