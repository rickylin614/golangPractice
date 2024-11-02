package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
)

const filepath = "auth.json"

type PurchaseToken struct {
	PackageName string `json:"packageName"` // app的包名, APP應用在google play商店的唯一識別碼
	ProductId   string `json:"productId"`   // Google play console中設定的商品ID
	Token       string `json:"token"`       // 購買令牌, 此筆交易的唯一識別碼
}

type ValidationResponse struct {
	Kind                        string `json:"kind"`                        // 資源的種類
	PurchaseTimeMillis          string `json:"purchaseTimeMillis"`          // 購買的時間，以毫秒為單位
	PurchaseState               int    `json:"purchaseState"`               // 購買的狀態 0.購買了 1. 已取消 2.待處理
	ConsumptionState            int    `json:"consumptionState"`            // 消費的狀態 0.承諾使用 1.使用
	DeveloperPayload            string `json:"developerPayload"`            // 開發者負載
	OrderId                     string `json:"orderId"`                     // 訂單 ID
	PurchaseType                int    `json:"purchaseType"`                // 購買類型 0.測試 1.促銷 2.獎勵廣告
	AcknowledgementState        int    `json:"acknowledgementState"`        // 確認狀態 0.尚未確認 1.已確認
	PurchaseToken               string `json:"purchaseToken"`               // 購買令牌
	ProductId                   string `json:"productId"`                   // 商品 ID
	Quantity                    int    `json:"quantity"`                    // 數量
	ObfuscatedExternalAccountId string `json:"obfuscatedExternalAccountId"` // 模糊化的外部帳戶 ID
	ObfuscatedExternalProfileId string `json:"obfuscatedExternalProfileId"` // 模糊化的外部個人資料 ID
	RegionCode                  string `json:"regionCode"`                  // 區域代碼
}

type AcknowledgePurchaseRequest struct {
	DeveloperPayload string `json:"developerPayload"`
}

type ConfirmationResponse struct {
	// 確認訂單的回應通常是空的，所以這裡不需要定義任何欄位
}

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message      string `json:"message"`
			Domain       string `json:"domain"`
			Reason       string `json:"reason"`
			Location     string `json:"location"`
			LocationType string `json:"locationType"`
		} `json:"errors"`
	} `json:"error"`
}

func main() {
	ctx := context.Background()

	// 從檔案中讀取服務帳戶金鑰
	jsonKey, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("讀取服務帳戶金鑰失敗:", err)
		return
	}

	// 解析服務帳戶金鑰
	config, err := google.JWTConfigFromJSON(jsonKey, "https://www.googleapis.com/auth/androidpublisher")
	if err != nil {
		fmt.Println("解析服務帳戶金鑰失敗:", err)
		return
	}

	// 創建 OAuth 2.0 客戶端
	client := config.Client(ctx)

	purchaseToken := &PurchaseToken{
		PackageName: "com.easy.subscriptions.classytaxi",
		ProductId:   "0001",
		Token:       "fpckjdfffnajmmlnbkohkpjo.AO-J1OwTRcbejGY8EWzRKX7eIkD3ed3KlONO_Zmx3eHQKF3UneXqdvKBXc-DIASmD2q4tLdpEfrNL8ghGFo2EewJJzItoMxXiPDu2m24Wq6hFNdJFmMTZqc",
	}

	// 取得訂單詳情
	getResponse, err := getOrder(client, purchaseToken)
	if err != nil {
		fmt.Printf("取得失敗: %+v\n", err)
		return
	}
	fmt.Println("取得成功:", getResponse)

	// 確認訂單
	confirmationResponse, err := confirmOrder(client, purchaseToken)
	if err != nil {
		fmt.Printf("驗證失敗: %+v\n", err)
		return
	}
	fmt.Println("確認成功:", confirmationResponse)
}

func getOrder(client *http.Client, purchaseToken *PurchaseToken) (*ValidationResponse, error) {
	url := fmt.Sprintf("https://www.googleapis.com/androidpublisher/v3/applications/%s/purchases/products/%s/tokens/%s",
		purchaseToken.PackageName, purchaseToken.ProductId, purchaseToken.Token)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		fmt.Println(string(body))
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("驗證失敗: %v", errorResponse)
	}

	var validationResponse ValidationResponse
	err = json.Unmarshal(body, &validationResponse)
	if err != nil {
		return nil, err
	}

	return &validationResponse, nil
}

func confirmOrder(client *http.Client, purchaseToken *PurchaseToken) (*ConfirmationResponse, error) {
	url := fmt.Sprintf("https://www.googleapis.com/androidpublisher/v3/applications/%s/purchases/products/%s/tokens/%s:acknowledge",
		purchaseToken.PackageName, purchaseToken.ProductId, purchaseToken.Token)

	acknowledgePurchaseRequest := &AcknowledgePurchaseRequest{
		DeveloperPayload: "yourDeveloperPayload", // developerPayload 是你自己定義的，並且應該是一個可以幫助你識別或驗證購買的唯一字串。這可能是一個隨機生成的唯一 ID，或者是一個包含特定資訊（如用戶 ID、購買時間等）的加密字串。你應該在你的伺服器或應用程式中生成和儲存這個字串。
	}

	requestBody, err := json.Marshal(acknowledgePurchaseRequest)
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		fmt.Println(string(body))
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("確認失敗: %v", errorResponse)
	}

	var confirmationResponse ConfirmationResponse
	err = json.Unmarshal(body, &confirmationResponse)
	if err != nil {
		return nil, err
	}

	return &confirmationResponse, nil
}
