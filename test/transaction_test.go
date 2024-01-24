package test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net/http"
	"testing"
)

func TestBroadcastTransaction(t *testing.T) {
	dataHex := "808000000004001cf09bd26ec797f7562563ac9ea437f6909342f70000000000000009000000000001bd790100cb88bc0da84438ac53930d52a507b9534317bed02b2a812c1370a60626bfeb8c440d232ec6e8d04c2dbd45fffc8318ec6beafc63fa977c8a2970b4da7012e28c03020000000000051aa2954b5eb8270cabeaef121d7fd05cd5e1b12ab7000000000000000100000000000000000000000000000000000000000000000000000000000000000000"
	// Set your API endpoint
	url := "https://api.testnet.hiro.so/v2/transactions"

	// Simulate the rawTx and attachment (set attachment to nil for this example)
	attachment := []byte(nil) // Set attachment to nil

	// Create a map for JSON payload
	decodeString, err := hex.DecodeString(dataHex)
	if err != nil {
		return
	}

	// Create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(decodeString))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the content type header based on the presence of attachment
	if attachment != nil {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", "application/octet-stream")
	}

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status and body
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}

type Transaction struct {
	Version int `json:"version"`
	Auth    struct {
		AuthType          int `json:"authType"`
		SpendingCondition struct {
			HashMode    int    `json:"hashMode"`
			Signer      string `json:"signer"`
			Nonce       string `json:"nonce"`
			Fee         string `json:"fee"`
			KeyEncoding int    `json:"keyEncoding"`
			Signature   struct {
				Type int    `json:"type"`
				Data string `json:"data"`
			} `json:"signature"`
		} `json:"spendingCondition"`
	} `json:"auth"`
	Payload struct {
		Type        int `json:"type"`
		PayloadType int `json:"payloadType"`
		Recipient   struct {
			Type    int `json:"type"`
			Address struct {
				Type    int    `json:"type"`
				Version int    `json:"version"`
				Hash160 string `json:"hash160"`
			} `json:"address"`
		} `json:"recipient"`
		Amount string `json:"amount"`
		Memo   struct {
			Type    int    `json:"type"`
			Content string `json:"content"`
		} `json:"memo"`
	} `json:"payload"`
	ChainID           int64 `json:"chainId"`
	PostConditionMode int   `json:"postConditionMode"`
	PostConditions    struct {
		Type              int   `json:"type"`
		LengthPrefixBytes int   `json:"lengthPrefixBytes"`
		Values            []any `json:"values"`
	} `json:"postConditions"`
	AnchorMode int `json:"anchorMode"`
}
