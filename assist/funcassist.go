package assist

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, length)
	for i := range randomString {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			fmt.Println("Lỗi khi tạo chuỗi ngẫu nhiên:", err)
			return ""
		}
		randomString[i] = charset[num.Int64()]
	}

	return string(randomString)
}