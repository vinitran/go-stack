package test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-stack/network"
	"go-stack/transactions"
	"log"
	"math/big"
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

func TestCreateAndBroadcastTx(t *testing.T) {
	senderKey := "1a1ff211642d6935539eacd404e83f7dc2ba08868b98a00cf3fe9bc26b3f7d56"
	privKey, err := transactions.CreateStacksPrivateKey(senderKey)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := transactions.GetPublicKey(privKey)

	transaction, err := transactions.MakeUnsignedSTXTokenTransfer(transactions.SignedTokenTransferOptions{
		TokenTransferOptions: transactions.TokenTransferOptions{
			Recipient:  "ST2H9AJTYQ0KGSAZAXW91TZYGBKAY3C9APZXSSGXW",
			Amount:     big.NewInt(1),
			Fee:        big.NewInt(78043),
			Nonce:      big.NewInt(10),
			Network:    network.StacksNetwork{},
			AnchorMode: 3,
			Memo:       "",
			Sponsored:  false,
		},
		PublicKey: publicKey,
	})

	signer, err := transactions.CreateTransactionSigner(&transaction)
	if err != nil {
		t.Fatal(err)
	}

	err = signer.SignOrigin(privKey)
	if err != nil {
		t.Fatal(err)
	}

	serializeTxBytes, err := transaction.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(hex.EncodeToString(serializeTxBytes))

	tx, _ := json.Marshal(transaction)
	fmt.Println(string(tx))
	url := "https://api.testnet.hiro.so/v2/transactions"

	attachment := []byte(nil) // Set attachment to nil

	// Create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(serializeTxBytes))
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
