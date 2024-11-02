package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"

	rsautil "test/ras2/rsa"
)

// var pubKey = `-----BEGIN PUBLIC KEY-----
// MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCN+rLx0e5DaYH8wY9XWjhzEwGGs+O3VvXqXfRa/wPCm/h4aXbvoEw12ZlY8K+yNOCA2pnWqQbsaV/MokAnE7zuBMwFojdaQPnudPPX1wAEjVjfJfQWCkhEGWyQjlqN9pPKgjh3ZChLJYhOYBIfD/rtaYBvThIEaqG5M6MCa49z7QIDAQAB
// -----END PUBLIC KEY-----
// `

var pubKey = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCN+rLx0e5DaYH8wY9XWjhzEwGGs+O3VvXqXfRa/wPCm/h4aXbvoEw12ZlY8K+yNOCA2pnWqQbsaV/MokAnE7zuBMwFojdaQPnudPPX1wAEjVjfJfQWCkhEGWyQjlqN9pPKgjh3ZChLJYhOYBIfD/rtaYBvThIEaqG5M6MCa49z7QIDAQAB"

var rsas rsautil.RSASecurity = rsautil.RSASecurity{}

func init() {
	// if err := rsas.SetPublicKey(pubKey); err != nil {
	// 	log.Fatalln(`set public key :`, err)
	// }
}

func main() {
	data := `SLflXrb4TF2tw9d9YLa8fEun+UG44lUg/JZ8R4wQNr/kZVbm2KCws5XYt4Ghii45eub9AgPw+6fYCbRv8vpDxCNtLtA1kRjAKJezBzq8hKertDBnhiKCzUX0zhmliEux6XaAeIr/JaO+UTl4JK0IjzqLLpMEioJ/CzGkEPf3phkychlRZ7SRef4f5zjckQDhQjU4ctQm683KV5ADTUwpFHyoxBPxa/ToUyEe9tzEofJKIi1dhjU+Bg/I1FZnin5FK6WOLEPN4LtYWw2GtzXk4UWrOyPve/4hUSYet7cuN9qRQu35ooKd49tXATZKg74YYbF7jE7uqm/oI2xGw2wKoQ==`

	databyte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	privateKey := "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAJhy+vTpX4k86sYQhxClIigRWk75F1ru4pGZSQXmGT1RkZBhPGk+OQuLUVPObL1izACqApyzIw/nsMr5kSmA7BU/NElcsYjJYlEarpeEyYbeS5jTJ+EHLRYtxZggcF2s9S97UYyEaekDdoUf/eG/zL4rYrKApCrHO92LOQHjuY/xAgMBAAECgYBSenq+AHkYpeWbLRv17HGzXdgPPALfxri770OrtlbXbwcjJwhpJCn2zfQ9NERunkIi5dgt8Vk55K8o0acw3jhbW78jB7BAaLxfckgT/ExqVu9md9+Qh/RVFeY2+EarhKd7BP4tmTnnOxXhl3+qS4cN7uTXhlickiRNN7dqJh0N5QJBANikMUhrrpvHxOTb2O+rSPkc4EN9e3G1LvQYsRDGXphnki8WiwmD1Yuc9BJYkztd8p4UGLGb02Q9nkwjYAh2LE8CQQC0JULaECYrOtoyM8Pbi+UM6oZm25zbRb7ZbFlxvkk/x4dNXNPIHPds+6IKDF3dnKONsPpJ1F9PD0MUgr/Q1C+/AkEAlSxiPxLe1ae2HTyA4W9ZPSe0COTzxnVTEoOaEQn3Awx2LXRhYrjjp1H5AlT5dKyZLl56Lno1ElYXlSfarZjpowJAHszEGk5qiDeeuLibAv1vIv8yDYH81oydLcVVoZncIjh2DKcTWoKBVzPKp5cnsU0ntYENufPCe9zrJiWYsBanNwJAbCRaduFhZjL7umRxn6KGSjeKIvHQhvmF7trjDibnZn967KzIXyQz08HwXrDV0UV2W78semsqKVqrhWBAuXYeHg=="
	// privateKey, err := base64.StdEncoding.DecodeString("MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAJhy+vTpX4k86sYQhxClIigRWk75F1ru4pGZSQXmGT1RkZBhPGk+OQuLUVPObL1izACqApyzIw/nsMr5kSmA7BU/NElcsYjJYlEarpeEyYbeS5jTJ+EHLRYtxZggcF2s9S97UYyEaekDdoUf/eG/zL4rYrKApCrHO92LOQHjuY/xAgMBAAECgYBSenq+AHkYpeWbLRv17HGzXdgPPALfxri770OrtlbXbwcjJwhpJCn2zfQ9NERunkIi5dgt8Vk55K8o0acw3jhbW78jB7BAaLxfckgT/ExqVu9md9+Qh/RVFeY2+EarhKd7BP4tmTnnOxXhl3+qS4cN7uTXhlickiRNN7dqJh0N5QJBANikMUhrrpvHxOTb2O+rSPkc4EN9e3G1LvQYsRDGXphnki8WiwmD1Yuc9BJYkztd8p4UGLGb02Q9nkwjYAh2LE8CQQC0JULaECYrOtoyM8Pbi+UM6oZm25zbRb7ZbFlxvkk/x4dNXNPIHPds+6IKDF3dnKONsPpJ1F9PD0MUgr/Q1C+/AkEAlSxiPxLe1ae2HTyA4W9ZPSe0COTzxnVTEoOaEQn3Awx2LXRhYrjjp1H5AlT5dKyZLl56Lno1ElYXlSfarZjpowJAHszEGk5qiDeeuLibAv1vIv8yDYH81oydLcVVoZncIjh2DKcTWoKBVzPKp5cnsU0ntYENufPCe9zrJiWYsBanNwJAbCRaduFhZjL7umRxn6KGSjeKIvHQhvmF7trjDibnZn967KzIXyQz08HwXrDV0UV2W78semsqKVqrhWBAuXYeHg==")
	if err != nil {
		panic(err)
	}

	b, err := rsautil.RSAPrivateDecode(databyte, string(privateKey))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}

func test1() {
	param := `amount=100&callbackUrl=https://c7-web-dev.paradise-soft.com.tw/pay&cashierType=PC&merchNo=A10751&orderNo=168424426&userId=yhnujm&viewUrl=https://c7-web-dev.paradise-soft.com.tw/pay`

	sign := md5.Sum([]byte(param + "&token=9fcee170-9beb-43f5-83d4-0cd68fdb80e1&sign=0ee7ee8bf7604b3f8c234137e14a6d2f"))

	param += "&sign=" + fmt.Sprintf("%x", sign[:])
	param += "&payChannelCode=card"

	fmt.Println(param)

	encoded, err := rsautil.RSAPublicSign(param, pubKey)
	fmt.Printf("%s %s\n", base64.StdEncoding.EncodeToString(encoded), err)
}
