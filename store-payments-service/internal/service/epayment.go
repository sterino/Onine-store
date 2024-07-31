package service

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func getToken() (string, error) {
	tokenURL := "https://testoauth.homebank.kz/epay2/oauth2/token"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("grant_type", "client_credentials")
	writer.WriteField("scope", "webapi usermanagement email_send verification statement statistics payment")
	writer.WriteField("client_id", "test")
	writer.WriteField("client_secret", "yF587AV9Ms94qN2QShFzVR3vFnWkhjbAK3sG")
	writer.WriteField("invoiceId", "00000000000001")
	writer.WriteField("amount", "100")
	writer.WriteField("currency", "KZT")
	writer.WriteField("terminalId", "67e34d63-102f-4bd1-898e-370781d0074d")

	err := writer.Close()
	if err != nil {
		return "", err
	}
	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}
	req, err := http.NewRequest("POST", tokenURL, body)
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status: %s, body: %s", res.Status, data)
	}
	var tokenResponse map[string]interface{}
	err = json.Unmarshal(data, &tokenResponse)
	if err != nil {
		return "", err
	}
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get access token")
	}
	return accessToken, nil
}

func getPublicKey() (*rsa.PublicKey, error) {
	publicKeyURL := "https://testepay.homebank.kz/api/public.rsa"
	resp, err := http.Get(publicKeyURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(body)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to parse RSA public key")
	}
	return rsaPublicKey, nil
}

func encryptData() (string, error) {
	publicKey, err := getPublicKey()
	if err != nil {
		return "", err
	}
	data := map[string]string{
		"hpan":       "4405639704015096",
		"expDate":    "0125",
		"cvc":        "815",
		"terminalId": "67e34d63-102f-4bd1-898e-370781d0074d",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, jsonData)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

type PaymentResponse struct {
	Status    string  `json:"status"`
	Message   string  `json:"message"`
	PaymentID string  `json:"payment_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	InvoiceID string  `json:"invoice_id"`
}

func MakePayment(amount float64) (*PaymentResponse, error) {
	paymentUrl := "https://testepay.homebank.kz/api/payment/cryptopay"
	token, err := getToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %v", err)
	}

	encryptedData, err := encryptData()
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt data: %v", err)
	}

	body := map[string]interface{}{
		"amount":          amount,
		"currency":        "KZT",
		"name":            "JON JONSON",
		"cryptogram":      encryptedData,
		"invoiceId":       "000000001",
		"invoiceIdAlt":    "8564546",
		"description":     "test payment",
		"accountId":       "uuid000001",
		"email":           "jj@example.com",
		"phone":           "77777777777",
		"cardSave":        true,
		"data":            `{\"statement\":{\"name\":\"Arman     Ali\",\"invoiceID\":\"80000016\"}}`,
		"postLink":        "https://testmerchant/order/1123",
		"failurePostLink": "https://testmerchant/order/1123/fail",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %v", err)
	}

	req, err := http.NewRequest("POST", paymentUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %v", err)
	}
	defer resp.Body.Close()
	var paymentResponse PaymentResponse

	if err = json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %s, body: %s", resp.Status, paymentResponse)
	}
	fmt.Println(paymentResponse)
	return &paymentResponse, nil
}
