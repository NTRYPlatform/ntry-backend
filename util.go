package notary

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegexUtil struct {
	email *regexp.Regexp
}

type ErrRequired struct {
	error
	arg string
}

type ErrInvalidValue struct {
	error
	arg string
}

func (e *ErrRequired) Error() string {
	return fmt.Sprintf("%s is required!", e.arg)
}

func (e *ErrInvalidValue) Error() string {
	return fmt.Sprintf("%s is invalid!", e.arg)
}

func NewRegexUtil() *RegexUtil {
	r := RegexUtil{}
	if e, err := regexp.Compile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)"); err != nil {
		fmt.Errorf("Couldn't compile regex! %v", err.Error())
	} else {
		r.email = e
	}
	return &r
}

func (r *RegexUtil) MatchEmail(email string) bool {
	return r.email.Match([]byte(email))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("bad password: %v", err)
	}
	return err == nil
}

// courtesy of SO:https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#%^&*()_-"
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
