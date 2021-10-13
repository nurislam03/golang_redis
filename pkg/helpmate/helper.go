package helpmate

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

const (
	Local = "local"
	Dev   = "dev"
	Stg   = "stg"
	Prod  = "prod"
	Test  = "prod"
)

//MaskStringData returns 12 digits masked string with only last 4 digits visible
func MaskStringData(data string) string {
	lenData := len(data)
	initial := "********"
	maskedData := initial + data[(lenData-4):]
	return maskedData
}

//RandNumberAsString return a random number as a string in a specific range
func RandNumberAsString(minSize, maxSize int64) (string, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(maxSize))
	if err != nil {
		return "", err
	}
	randNum := nBig.Int64()
	randNum += minSize

	//convert nBig (big integer) to a string
	data := fmt.Sprintf("%d", randNum)
	return data, nil
}

//SHA256Hash returns hex encoded string of hash value
func SHA256Hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func SHA256HMACHash(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func RandSecret(size uint) string {
	if size == 0 {
		size = 20
	}
	secret := make([]byte, size)
	_, _ = rand.Reader.Read(secret)
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(secret)
}

func FormatBDMobile(m string) string {
	return "+88" + m[len(m)-11:]
}

func IsLocalEnv() bool {
	return viper.GetString("server.env") == Local
}

func IsDevelopmentEnv() bool {
	return viper.GetString("server.env") == Dev
}

func IsStagingEnv() bool {
	return viper.GetString("server.env") == Stg
}

func IsProductionEnv() bool {
	return viper.GetString("server.env") == Prod
}

func IsTestEnv() bool {
	return viper.GetString("server.env") == Test
}

func GetUserSlug(r *http.Request) string {
	userSlug := r.Header.Get("USER-SLUG")
	if userSlug != "" {
		return strings.TrimSpace(userSlug)
	}
	return ""
}

func GetUserPhone(r *http.Request) string {
	userSlug := r.Header.Get("USER-PHONE")
	if userSlug != "" {
		return strings.TrimSpace(userSlug)
	}
	return ""
}
