package util

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(ALPHABET)

	for range n {
		c := ALPHABET[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{USD, EUR, PLN}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {

	return fmt.Sprintf("%s@email.com", RandomString(15))
}
