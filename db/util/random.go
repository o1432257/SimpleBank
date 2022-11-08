package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// 設定種子
	rand.Seed(time.Now().UnixNano())
}

// RandomInt 產生介於 min 與 max 的隨機值
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString 產生長度為 n 的隨機字串
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner 產生隨機owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney 產生隨機金額
func RandomMoney() int64 {
	return RandomInt(0, 5000)
}

// RandomCurrency 產生隨機幣種
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "TWD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
