package main

import "C"
import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/CoboGlobal/cobo-go-api/cobo_custody"
	"github.com/btcsuite/btcd/btcec"
)

func main() {
}

//type LocalSigner struct {
//	PrivateKey string
//}

func Hash256(s string) string {
	hashResult := sha256.Sum256([]byte(s))
	hashString := string(hashResult[:])
	return hashString
}

func Hash256x2(s string) string {
	return Hash256(Hash256(s))
}

// Sign :
//
//export Sign
func Sign(private *C.char, msg *C.char) *C.char {
	signer := cobo_custody.LocalSigner{
		PrivateKey: C.GoString(private),
	}
	sig := signer.Sign(C.GoString(msg))
	fmt.Println(sig)
	return C.CString(sig)
}

// GetPublicKey :
//
//export GetPublicKey
func GetPublicKey(privateKey *C.char) *C.char {
	signer := cobo_custody.LocalSigner{
		PrivateKey: C.GoString(privateKey),
	}
	pub := signer.GetPublicKey()
	fmt.Println(pub)
	return C.CString(pub)
}

//func (signer LocalSigner) Sign(message string) string {
//	apiSecret, _ := hex.DecodeString(signer.PrivateKey)
//	key, _ := btcec.PrivKeyFromBytes(btcec.S256(), apiSecret)
//	sig, _ := key.Sign([]byte(Hash256x2(message)))
//	return fmt.Sprintf("%x", sig.Serialize())
//}
//
//func (signer LocalSigner) GetPublicKey() string {
//	apiSecret, _ := hex.DecodeString(signer.PrivateKey)
//	key, _ := btcec.PrivKeyFromBytes(btcec.S256(), apiSecret)
//	return fmt.Sprintf("%x", key.PubKey().SerializeCompressed())
//}

// GenerateKeyPair :
//
//export GenerateKeyPair
func GenerateKeyPair() (*C.char, *C.char) {
	apiSecret := make([]byte, 32)
	if _, err := rand.Read(apiSecret); err != nil {
		panic(err)
	}
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), apiSecret)
	apiKey := fmt.Sprintf("%x", privKey.PubKey().SerializeCompressed())
	apiSecretStr := fmt.Sprintf("%x", apiSecret)

	fmt.Printf("[GO] apiKey %s\n", apiKey)
	fmt.Printf("[GO] apiSecretStr %s\n\n", apiSecretStr)
	return C.CString(apiSecretStr), C.CString(apiKey)
}
