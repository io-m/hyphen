package helpers

import (
	crypto "crypto/rand"
	"encoding/hex"
	"fmt"

	"math/rand"
	"strings"
	"time"
)

var alphabetSmall = "abcdefghijklmnopqrstuvwxyz"
var alphabetCapital = strings.ToUpper(alphabetSmall)
var alphabet = fmt.Sprintf("%s%s", alphabetSmall, alphabetCapital)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano())).Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// GenerateRandomKey generates gibberish sequence of provided length
func GenerateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := crypto.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
