package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func main() {
	// 生成RSA 512密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		panic(err)
	}

	// 将私钥保存为PEM文件
	pemPrivateFile, err := os.Create("private_key.pem")
	if err != nil {
		panic(err)
	}
	defer pemPrivateFile.Close()

	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	if err := pem.Encode(pemPrivateFile, privateBlock); err != nil {
		panic(err)
	}

	// 将公钥保存为PEM文件
	pemPublicFile, err := os.Create("public_key.pem")
	if err != nil {
		panic(err)
	}
	defer pemPublicFile.Close()

	publicKey := privateKey.PublicKey
	publicBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&publicKey),
	}
	if err := pem.Encode(pemPublicFile, publicBlock); err != nil {
		panic(err)
	}
}
