package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
)

type LocalSigner struct {
	privateKey string
}

func Hash256(s string) string {
	hashResult := sha256.Sum256([]byte(s))
	hashString := string(hashResult[:])
	return hashString
}

func Hash256x2(s string) string {
	return Hash256(Hash256(s))
}

func (signer LocalSigner) Sign(message string) string {
	apiSecret, _ := hex.DecodeString(signer.privateKey)
	key, _ := btcec.PrivKeyFromBytes(btcec.S256(), apiSecret)
	sig, _ := key.Sign([]byte(Hash256x2(message)))
	return fmt.Sprintf("%x", sig.Serialize())
}

func GenerateKeyPair() (string, string) {
	apiSecret := make([]byte, 32)
	if _, err := rand.Read(apiSecret); err != nil {
		panic(err)
	}
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), apiSecret)
	apiKey := fmt.Sprintf("%x", privKey.PubKey().SerializeCompressed())
	apiSecretStr := fmt.Sprintf("%x", apiSecret)
	return apiSecretStr, apiKey
}
