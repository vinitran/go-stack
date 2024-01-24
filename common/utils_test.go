package common

//func TestHex2Bytes(t *testing.T) {
//	privateKeyHex := "deadbeef"
//	test, err := Hex2Bytes(privateKeyHex)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	assert.Equal(t, test, []uint8{222, 173, 190, 239}, "invalid")
//
//	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
//	if err != nil {
//		t.Fatal(err)
//	}
//	assert.Equal(t, privateKeyBytes, []uint8{222, 173, 190, 239}, "invalid")
//}

//func TestEncodePrivateKey(t *testing.T) {
//	// Decode the hex-encoded private key.
//	pkBytes, err := hex.DecodeString("a11b0a4e1a132305652ee7a8eb7848f6ad" +
//		"5ea381e3ce20a2c086a2e388230811")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	privKey := secp256k1.PrivKeyFromBytes(pkBytes)
//
//	ciphertext, err := hex.DecodeString("35f644fbfb208bc71e57684c3c8b437402ca" +
//		"002047a2f1b38aa1a8f1d5121778378414f708fe13ebf7b4a7bb74407288c1958969" +
//		"00207cf4ac6057406e40f79961c973309a892732ae7a74ee96cd89823913b8b8d650" +
//		"a44166dc61ea1c419d47077b748a9c06b8d57af72deb2819d98a9d503efc59fc8307" +
//		"d14174f8b83354fac3ff56075162")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//privKey.PubKey()
//	// Try decrypting the message.
//	plaintext, err := secp256k1.(privKey, ciphertext)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println(string(plaintext))
//}
