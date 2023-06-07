package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/farmerx/gorsa"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDKEOImLExwI1r/66Et1BTRJPTzkuBlm7hXobZEqqJS6ANhhovu7jp9oYxDOffnb7notxAOlD2+0tLC+sclBnpZ8c6KnaL31EAzQectX61/uZG/QjcPhv4OLpHvMJYnTPwSaGi9lX11HM/mKRBUHss8OLM0d8s5Km63bUzUJnORcoaSf1DZxviIMIdLXhe+G6QfTfnATGyDGm8FvC9hES9ErN7knMd1ZoGOw3ot/6WIgWEz6s5CMVVJLVLBmyWC8F7vNWfDyDxnt5QYD4Giv37872RZEXKExQeMDwPlku0RTig7DtVMVR9ZiJIc/MNo8NgzI9GGA1SoVAX0LSG+fLRZAgMBAAECggEAcff71dP+ePE4Dkh0bEjGRQA3OHqLHkQqQSLwjuI/E8vQvF0K6ARt8RYA6pnzfRHDJcK6x75M5Sg7v8V6NSPmNnzwJbariGSqoT52iV0Bmpyr29gY/iUBfjY+EU6yIhCFzEwfZ269iCRsDkDy+L81mg3Q5bL7aI7KS6LWRoI5hLq1n7p4lRUGufZJ5dXqWjxmOaRP1U9MoGvZIGMToAEbF0KedyeDCfQK9bLk9nsxdILsv+L+V/m5OV1JHzL6+lz94HP5XBf6u8PEj0qwDrVBwtSZRCDr5oPQ+a406op1X5ItaNk7BcRnCdYhnEEgI4R5XPLgY2Ha/83Z5JRaTdI1QQKBgQD9jDmFyEETUS4zuT+Q/qh2ryNHbXvf2CCwmdsxnXddH0axasYfv7KAFnfYHk3FotdPy0glt23IAlgY3Blzru0v5iNDQQUM1XeyHVVyMZlDx/l/hq8qiksjCdoCfmjMMUfL+gHCArKMp4wPL3dX5kl1EcxZiuTJ8dUVfPk7hZ6szQKBgQDMBTEKtFkS5001JwWtAQ7yOTcDl37XTPotJa3TDY/TlYcMeEYBJsORe9syhV/gEL1B11Y9M5VO0EamtYZjli4tMnIeoypNmksu428hFDKqoFHPfFoi6oyRVb//0Zo2vwGk5atlkkZ74We5dTCQZ+ONHhlLcnlyBLHKItUp/tOlvQKBgGZhG80Qam6BZjWOhMNogN/DomwyIkXlwHVu+tpS6iLZATOcmLmzL0GXUPccbWzzbxuSSuSn55VOlQIcxc43gSDJPc3nUEG5ZKawP5Nnwapj8jPzT4HO674POU+ueWSSNZqArfsIGc6/zMdI9Fiy7VBvSmORuQZuNwuBNiY5szO1AoGAWKQgzJ/S2423U9KieqotDjyhlGPr7is5vwfR4BLSXIFo7ZgMAlAPRFRlxiuAnl9newXOKEUTJ+0B5UiVHQuRAH4cawdFR/YzmCmkjOCzqsKBCP7kQqtZa4OwKbutnv+Z/UT0FOtayFQEtuUdtoAMs7FSGB+ieFFl1YcQynA3EJkCgYEA7S6iwNzmitxAQZ0a3f+QtxdGZ2VhJV+M78j5dqRDjTCwHKbouKgplfNNm6aId3LuIwL7k30mja0GjO2QqCNovurPCj7rmZEPZPpVzLn0vanZ40igFMce3nlKsbWxwwlwQKOUkK2N6BiQDS8dt9GobwWuBkgq5vYJFwQEMGEWRLU=
-----END RSA PRIVATE KEY-----`

const privateKey2 = `
-----BEGIN RSA PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCrLM9baz4dOmpXzZAJVnBWASV4
X6VcuzubwYB2rhO0ypyQybrN6pjh9KlIFwPH/q8tPLng727zmIT63mjGyjpeMvKcIJV3brAmpAas
GMg0k6eNhXMXO8VVOhHtfWtM53mFkfO8ocapV3q9StBqqkXkP6JTk1kC+Ou/SULuuATVYr7BDsZO
TpjMjLPBdXkdeJmvrN087l/38Eysmbxfv1r7pS8Js+N87Gm0Y+V5yJqCymK7daQ0Said73NRtQlv
WWR6o7Il4Lkpvm2D+1nqRpnR/90iVSY+iZFjwcwj5EJd2tv/PO1P5X5WQp7bMavw3SUtNlbqSjX2
c14vL4ngC+9HAgMBAAECggEAad6/KVW9ByrYFxJG4d+afzcaS8csaZxME9usbrrf6VPfXhXG+hD9
GdPO142Ugx/cjOggNyUfxH4y0U66EymEzDpKn5Hvn/zP6jZNAigZu5sPaJJ/SoFnXX0bFX4a3WfV
W22YCLUuQBKhEW6yVSaXL8LDzVk3FoWKmbbgNPGP28G/RqEFiOT8R+NP+gEt2TjW5SrfjR+0pfl1
+iYx+//TO1OaYazR9d3jWeVGGQq5zOsYXGLnrw0gr3FtG1lFI3fv6nVaNRtonjD8rOdwJ5muj1MQ
CckqtOraqGLBzhEOSwNzwzkZltURUzQw9BnpQ5P0KlN1g1JHdoyWZW8yToQpQQKBgQDzPyWwX975
m7zz1P2dwaq5Fc14DFvkHoK+xuq89ONKIUnzHlVgCaOPsElQ4h1VEwmzDcyBEywM8rPN9elu1C+h
yeNIyVgmfS4fSxrcela3nfhcD9TVdqdLmzh/KITuDQN2yLHzu6jqQ2+eDwWhYC3a7MkMREHgU+tC
p3jAT3M+0QKBgQC0JlGLy5UxV751TUfFKap7xwTEV4T4O4KFzY571c+HGb96fLX2d547ZIPgJr1H
O4DwtTtE5K6a8f5EQZ7pO1J2cepHxZDV6dQOBIGsF0qGewHjIwpsgZFb1HOPmDXvIBx1c+aPUwq/
JREmAfnpT1w3E1KaPx0LhXqspt2UduBClwKBgH+geFPsWX+5+RAuhstJHBx0lhQQ1/3/DDaxA+UU
0FTK1pQbLExxgkIYCr/HhABpNsOdz9lBh/FbPoAoMMyXxE4tMXW1Pu0i7Ro6O+PGxM0es3sKS2+d
C/YdgS0mw2hrp4+hic0Kf3w62cuish+T6Tte1bX78lyTa3LTkAXHKf6hAoGAFQloel+io2lwpzgI
CoEMFHowF1F1CJvCaSeQ3Osh2c/q6T3I7egYBS1+rN05Oyk6rGEtc6UUsqlRnNzg1rGDtqskxY7P
k/tDNQPXtKXfoQaaONDttbAHrmaHvBv1KGBikmCfatsypRGKCBsw+Mq4pQHPoo1+Mcs7gQ/XTi/3
ji8CgYAksHG4JuobVu45Li+Uxc+jxPD0eFZnsPxL0AD08gfwPYo4RvRqutbsfyywgVm2idUNkrZJ
hAsq7wZnGVBIx4+P3+cYAEFvbT2ub6/DEOo2ni43RQs8D0XiJEDcfjW7e2LXk0TQCEBqYFJAxUs2
3Q9q2wnbqUSTaKOa53hHFwmPbg==
-----END RSA PRIVATE KEY-----
`

const pubKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyhDiJixMcCNa/+uhLdQU0ST085LgZZu4V6G2RKqiUugDYYaL7u46faGMQzn352+56LcQDpQ9vtLSwvrHJQZ6WfHOip2i99RAM0HnLV+tf7mRv0I3D4b+Di6R7zCWJ0z8EmhovZV9dRzP5ikQVB7LPDizNHfLOSput21M1CZzkXKGkn9Q2cb4iDCHS14XvhukH035wExsgxpvBbwvYREvRKze5JzHdWaBjsN6Lf+liIFhM+rOQjFVSS1SwZslgvBe7zVnw8g8Z7eUGA+Bor9+/O9kWRFyhMUHjA8D5ZLtEU4oOw7VTFUfWYiSHPzDaPDYMyPRhgNUqFQF9C0hvny0WQIDAQAB
-----END PUBLIC KEY-----
`

var (
	ErrDataToLarge     = errors.New("message too long for RSA public key size")
	ErrDataLen         = errors.New("data length error")
	ErrDataBroken      = errors.New("data broken, first byte is not zero")
	ErrKeyPairDismatch = errors.New("data is not encrypted by the private key")
	ErrDecryption      = errors.New("decryption error")
	ErrPublicKey       = errors.New("get public key error")
	ErrPrivateKey      = errors.New("get private key error")
)

// func main() {

// 	source := "ABC"
// 	sign, err := RSASign(source, privateKey)
// 	fmt.Printf("sign:%s err:%v\n", sign, err)
// 	// if err := gorsa.RSA.SetPublicKey(pubKey); err != nil {
// 	// 	fmt.Println("errrr:", err)
// 	// }
// 	// result, err := gorsa.RSA.PubKeyDECRYPT(sign)
// 	// fmt.Printf("result:%s err:%v\n", result, err)
// }

func main() {
	if err := gorsa.RSA.SetPublicKey(pubKey); err != nil {
		log.Fatalln(`set public key :`, err)
	}
	if err := gorsa.RSA.SetPrivateKey(privateKey); err != nil {
		log.Fatalln(`set private key :`, err)
	}

	// data := "URNAMgYU1foVpBjN8kBV+6nN1ckW6bUH35pxRP3uifMM4XxzktL4AlRIAhASuy868WLoQDyYXg9b3N2RNMeie2UFP8lVEIv1r2UAfSYFE80uW6EJdewD2dhVxugZxB7jCt9W+Wgh3eX59OW2KLiyXdJXWqLUA/AOlxYihd5czutrSBUlptorh7byQwzHtJnC63kVcPGPhA8ihYN9p4HBZL/fwayVdmftDoaJ6ayNQ5s11K1UkchR+gmowhQdgLoNm8gGxKH08nnEIXTM8+hKQAnjOWjXwiezSHKSV5rOuEHoJLTfb2gM+hc1ClfuBFVJqzxj0bna+2+epGOyeyLoWQ=="
	data := `SLflXrb4TF2tw9d9YLa8fEun+UG44lUg/JZ8R4wQNr/kZVbm2KCws5XYt4Ghii45eub9AgPw+6fYCbRv8vpDxCNtLtA1kRjAKJezBzq8hKertDBnhiKCzUX0zhmliEux6XaAeIr/JaO+UTl4JK0IjzqLLpMEioJ/CzGkEPf3phkychlRZ7SRef4f5zjckQDhQjU4ctQm683KV5ADTUwpFHyoxBPxa/ToUyEe9tzEofJKIi1dhjU+Bg/I1FZnin5FK6WOLEPN4LtYWw2GtzXk4UWrOyPve/4hUSYet7cuN9qRQu35ooKd49tXATZKg74YYbF7jE7uqm/oI2xGw2wKoQ==`

	databyte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	privateKey, err := base64.StdEncoding.DecodeString("MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAJhy+vTpX4k86sYQhxClIigRWk75F1ru4pGZSQXmGT1RkZBhPGk+OQuLUVPObL1izACqApyzIw/nsMr5kSmA7BU/NElcsYjJYlEarpeEyYbeS5jTJ+EHLRYtxZggcF2s9S97UYyEaekDdoUf/eG/zL4rYrKApCrHO92LOQHjuY/xAgMBAAECgYBSenq+AHkYpeWbLRv17HGzXdgPPALfxri770OrtlbXbwcjJwhpJCn2zfQ9NERunkIi5dgt8Vk55K8o0acw3jhbW78jB7BAaLxfckgT/ExqVu9md9+Qh/RVFeY2+EarhKd7BP4tmTnnOxXhl3+qS4cN7uTXhlickiRNN7dqJh0N5QJBANikMUhrrpvHxOTb2O+rSPkc4EN9e3G1LvQYsRDGXphnki8WiwmD1Yuc9BJYkztd8p4UGLGb02Q9nkwjYAh2LE8CQQC0JULaECYrOtoyM8Pbi+UM6oZm25zbRb7ZbFlxvkk/x4dNXNPIHPds+6IKDF3dnKONsPpJ1F9PD0MUgr/Q1C+/AkEAlSxiPxLe1ae2HTyA4W9ZPSe0COTzxnVTEoOaEQn3Awx2LXRhYrjjp1H5AlT5dKyZLl56Lno1ElYXlSfarZjpowJAHszEGk5qiDeeuLibAv1vIv8yDYH81oydLcVVoZncIjh2DKcTWoKBVzPKp5cnsU0ntYENufPCe9zrJiWYsBanNwJAbCRaduFhZjL7umRxn6KGSjeKIvHQhvmF7trjDibnZn967KzIXyQz08HwXrDV0UV2W78semsqKVqrhWBAuXYeHg==")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x %d\n", privateKey, len(privateKey))
	err = gorsa.RSA.SetPrivateKey(string(privateKey))
	if err != nil {
		panic(err)
	}
	b, err := gorsa.RSA.PriKeyDECRYPT([]byte(databyte))
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// RSA 公鑰加密
func RSAPubEncrytion(pubKey *rsa.PublicKey, source string) ([]byte, error) {
	data := []byte(source)
	encrypted := make([]byte, 0, len(data))
	for i := 0; i < len(data); i += 117 {
		finish := i + 117
		if finish > len(data) {
			finish = len(data)
		}
		part, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, data[i:finish])
		if err != nil {
			return nil, err
		}
		encrypted = append(encrypted, part...)
	}
	return encrypted, nil
}

func RSASign(source, key string) ([]byte, error) {
	block, _ := pem.Decode([]byte(key))
	fmt.Println("block:", block.Type)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block")
	}

	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	priKey, ok := pri.(*rsa.PrivateKey)
	if !ok {
		return nil, err
	}

	data := []byte(source)
	return PrivateEncrypt(priKey, data)
}

func pubKeyDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	k := (pub.N.BitLen() + 7) / 8
	if k != len(data) {
		return nil, ErrDataLen
	}
	m := new(big.Int).SetBytes(data)
	if m.Cmp(pub.N) > 0 {
		return nil, ErrDataToLarge
	}
	m.Exp(m, big.NewInt(int64(pub.E)), pub.N)
	d := leftPad(m.Bytes(), k)
	if d[0] != 0 {
		return nil, ErrDataBroken
	}
	if d[1] != 0 && d[1] != 1 {
		return nil, ErrKeyPairDismatch
	}
	var i = 2
	for ; i < len(d); i++ {
		if d[i] == 0 {
			break
		}
	}
	i++
	if i == len(d) {
		return nil, nil
	}
	return d[i:], nil
}

func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

func PrivateEncrypt(priv *rsa.PrivateKey, data []byte) (enc []byte, err error) {
	k := (priv.N.BitLen() + 7) / 8
	tLen := len(data)
	// rfc2313, section 8:
	// The length of the data D shall not be more than k-11 octets
	if tLen > k-11 {
		return
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], data)
	c := new(big.Int).SetBytes(em)
	if c.Cmp(priv.N) > 0 {
		return
	}
	var m *big.Int
	var ir *big.Int
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		// We have the precalculated values needed for the CRT.
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}

	if ir != nil {
		// Unblind.
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}
	enc = m.Bytes()
	return
}

// Parse PEM encoded PKCS1 or PKCS8 public key
func ParseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, ErrNotRSAPublicKey
	}

	return pkey, nil
}

func OldRsaEn(source string, key string) ([]byte, string, error) {
	pemSource := fmt.Sprintf(`%s`, key)
	pemBytes := []byte(pemSource)

	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "PUBLIC KEY" {

		return nil, "", errors.New("failed to decode PEM block")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {

		return nil, "", err
	}

	pubKey := pub.(*rsa.PublicKey)
	hash := sha256.New()
	msgLen := len(source)
	step := pubKey.Size() - 2*hash.Size() - 2
	var encryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}
		encryptedBlockBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(source)[start:finish])
		if err != nil {

			return nil, "", err
		}
		encryptedBytes = append(encryptedBytes, encryptedBlockBytes...)
	}
	return encryptedBytes, source, nil
}

var (
	ErrKeyMustBePEMEncoded = errors.New("Invalid Key: Key must be PEM encoded PKCS1 or PKCS8 private key")
	ErrNotRSAPrivateKey    = errors.New("Key is not a valid RSA private key")
	ErrNotRSAPublicKey     = errors.New("Key is not a valid RSA public key")
)

// Parse PEM encoded PKCS1 or PKCS8 private key
func ParseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, ErrKeyMustBePEMEncoded
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, ErrNotRSAPrivateKey
	}

	return pkey, nil
}

// RSA 私鑰加密
func RSAPriEncryption(priv *rsa.PrivateKey, source string) (enc []byte, err error) {
	defer func() {
		if errs := recover(); errs != nil {
			err, _ = errs.(error)
		}
	}()
	data := []byte(source)

	k := (priv.N.BitLen() + 7) / 8
	tLen := len(data)
	// rfc2313, section 8:
	// The length of the data D shall not be more than k-11 octets
	if tLen > k-11 {
		return
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], data)
	c := new(big.Int).SetBytes(em)
	if c.Cmp(priv.N) > 0 {
		return
	}
	var m *big.Int
	var ir *big.Int
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		// We have the precalculated values needed for the CRT.
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}

	if ir != nil {
		// Unblind.
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}
	enc = m.Bytes()
	return
}
