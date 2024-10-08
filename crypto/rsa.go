package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"
)

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	return key, &key.PublicKey
}

func StringifyPrivateKey(key *rsa.PrivateKey) string {
	keyBytes := x509.MarshalPKCS1PrivateKey(key)
	privateKey := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: keyBytes,
		},
	)
	return string(privateKey)
}

func ParsePrivateKey(key string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func StringifyPublicKey(key *rsa.PublicKey) (string, error) {
	keyBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return "", err
	}
	publicKey := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: keyBytes,
		},
	)

	return string(publicKey), nil
}

func ParsePublicKey(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("key type is not RSA")
}

func Encrypt(plain string, publicKey *rsa.PublicKey) (string, error) {
	cipher, err := rsa.EncryptOAEP(
		sha512.New(),
		rand.Reader,
		publicKey,
		[]byte(plain),
		nil)
	return string(cipher), err
}

func Decrypt(cipher string, privateKey *rsa.PrivateKey) (string, error) {
	plain, err := privateKey.Decrypt(nil, []byte(cipher), &rsa.OAEPOptions{Hash: crypto.SHA512})
	return string(plain), err
}

func Sign(msg string, privateKey *rsa.PrivateKey) (string, error) {
	msgHash := sha512.New()
	_, err := msgHash.Write([]byte(msg))
	if err != nil {
		return "", err
	}
	msgHashSum := msgHash.Sum(nil)

	sign, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA512, msgHashSum, nil)
	return string(sign), err
}

func Verify(msg string, sign string, publicKey *rsa.PublicKey) bool {
	msgHash := sha512.New()
	_, err := msgHash.Write([]byte(msg))
	if err != nil {
		return false
	}
	msgHashSum := msgHash.Sum(nil)
	err = rsa.VerifyPSS(publicKey, crypto.SHA512, msgHashSum, []byte(sign), nil)
	return err == nil
}

func main() {
	start := time.Now()
	// Create the keys
	priv, pub := GenerateRsaKeyPair()

	// Export the keys to pem string
	priv_pem := StringifyPrivateKey(priv)
	pub_pem, _ := StringifyPublicKey(pub)

	// Import the keys from pem string
	priv_parsed, _ := ParsePrivateKey(priv_pem)
	pub_parsed, _ := ParsePublicKey(pub_pem)

	// Export the newly imported keys
	priv_parsed_pem := StringifyPrivateKey(priv_parsed)
	pub_parsed_pem, _ := StringifyPublicKey(pub_parsed)

	fmt.Println(priv_parsed_pem)
	fmt.Println(pub_parsed_pem)
	fmt.Println(strings.ReplaceAll(pub_parsed_pem, "\n", ""))
	fmt.Println(priv.E)

	// Check that the exported/imported keys match the original keys
	if priv_pem != priv_parsed_pem || pub_pem != pub_parsed_pem {
		fmt.Println("Failure: Export and Import did not result in same Keys")
	} else {
		fmt.Println("Success")
	}
	fmt.Println(time.Since(start).Milliseconds())

	start = time.Now()
	plain := "afsdjhasfdjhfasdjhkfajhkasdfjafsjasfdjfjhfdjflkjasdhfsdfhadfhlasjfljakfh"
	cipher, _ := Encrypt(plain, pub_parsed)
	decrypted, _ := Decrypt(cipher, priv)
	fmt.Println(plain == decrypted)
	fmt.Println(time.Since(start).Milliseconds())

	sign, _ := Sign(plain, priv)
	fmt.Println(Verify(plain, sign, pub))
	plain = "afsdjhasfdjhfasdjhkfajhkasdfjafsjasfdjfjhfdjflkjasdhfsdfhadfhlasjfljakfh1"
	fmt.Println(Verify(plain, sign, pub))
}
