package main

// You can edit this code!
// Click here and start typing.

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	encryptedKey := []byte{196, 231, 38, 149, 234, 51, 142, 186, 230, 214, 96, 243, 229, 153, 103, 74, 117, 241, 237, 135, 91, 170, 216, 167, 235, 154, 180, 28, 125, 48, 82, 44}
	encryptedData := "bj4hUUJtwIZJnnGYO+04JpIxf1Omnsrq+ilLv+eQXSEbvsEKgD9BZ0cnxGBrvd4KHeyHkoiZZbObDkZemFtQkgo0jw1AtscTv4HSIz9OeOFZmNjjA724dA8oyuHY8WWlFddXcArqv4R5a9DO2Qs4e9sk3KIRYn4OJH//lNi8t111J//wcsSTvuer+EO9XeQIZHsZ4sr+fhdU1Jxz7Z1KOM6c2MGuXAqfKj/ebn01XA/LgnaDO8xZ+k7vs19WY0pH33HM0K5K1C+XUVaRGhUtFt2BDgnG5T/MPIyVlfUpJpHJjEdsvTxTN7muK9UsbSonr3XAp9oV/w7xx65N5iEEjvC1CaehAnKOSM6QiyKs8KLiv2vHaabsNNusFZgHUZYUApErbpK9g3UR0744Sy8hVe/l75SmOofZvu59vU509MMc8TdqxssIv2g7JTw="
	plaintext, err := Decode(encryptedKey, encryptedData)
	if err != nil {
		fmt.Errorf("an error occured while decrypting: %s", err.Error())
	}
	fmt.Printf("in hex format: %x", plaintext)
}

func Decode(encryptionKey []byte, encryptedData string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}
