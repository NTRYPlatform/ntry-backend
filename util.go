package auth

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/openpgp"
)

// courtesy of SO:https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// var publicKey = "/home/someone/Downloads/pub-key.txt"
// var signatureFile = "/home/someone/Downloads/singed.txt"

// RandString generates random string of size n
func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// VerifySignature verifies thesignature
func VerifySignature(keyRing, signature, target string) (valid bool) {

	valid = false
	keyReader := strings.NewReader(keyRing)

	targetReader := strings.NewReader(target)

	signatureReader := strings.NewReader(signature)

	keyring, err := openpgp.ReadArmoredKeyRing(keyReader)
	if err != nil {
		log.Printf("Read Armored Key Ring: %v ", err.Error())
		return
	}

	_, err = openpgp.CheckArmoredDetachedSignature(keyring, targetReader, signatureReader)
	if err != nil {
		log.Printf("Check Detached Signature: %v ", err.Error())
		return
	}

	valid = true
	return
}

// func main() {
// 	key, _ := ioutil.ReadFile("/home/someone/Downloads/pub-key.txt")
// 	signature, _ := ioutil.ReadFile("/home/someone/Downloads/singed.txt")
// 	log.Print(VerifySignature(string(key), string(signature), "this is test"))
// }
