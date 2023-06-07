package main

import (
	"fmt"

	"github.com/pquerna/otp/totp"
)

func main() {
	// -- 產生新KEY
	key, err := totp.Generate(totp.GenerateOpts{Issuer: "paradise prod ia finance", AccountName: "paradise-soft.com.tw"})
	key2, err := totp.Generate(totp.GenerateOpts{Issuer: "paradise prod ia user", AccountName: "paradise-soft.com.tw"})
	key3, err := totp.Generate(totp.GenerateOpts{Issuer: "paradise prod ia Ops", AccountName: "paradise-soft.com.tw"})
	key4, err := totp.Generate(totp.GenerateOpts{Issuer: "paradise prod ia order", AccountName: "paradise-soft.com.tw"})
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println(key.Secret())  // JD5MKJQJMLCNWRN4TMR43ZPB2VKS6H2I // secret_finance 验证财务OTP
		fmt.Println(key2.Secret()) // R67EHKV5F44NJTWUOD2TX5OFLEBESFNU // secret_user_login 验证用户登录OTP
		fmt.Println(key3.Secret()) // QDL5YBSQNVHXS6C3IDX63AKXRTVDIHTN // secret_ops 验证运营OTP
		fmt.Println(key4.Secret()) // 22LC4FJQVUQRP2ACVMSQXL5T6GBGIGVG // secret_order 验证訂單OTP
	}

	// -- 驗證KEY
	b := totp.Validate("261512", "7IORPDLLFBDGDE2ASK3S3PQV2PAYEQHG")
	fmt.Println("驗證通過?", b)
}
