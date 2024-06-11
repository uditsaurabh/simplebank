package util

import (
	"math/rand"
	"strings"
	"time"
	"fmt"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func RandomPassword(n int) string {
	if n == 0 {
		n = 15
	}
	var str string
	var c string
	for i := 0; i < n; i++ {
		n := rand.Intn(2)
		if n == 0 {
			c = string(alphabets[rand.Intn(len(alphabets))])

		} else {
			c = fmt.Sprintf("%d",numbers[rand.Intn(len(numbers))])
		}
		str = str + c
	}
	fmt.Println("resulting password",str)
	return str
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
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

func RandomCurrencies() string {
	currencies := []string{"Rupees", "Dollars", "Euros"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomCountryCode() int64 {
	return RandomInt(100, 1000)
}
