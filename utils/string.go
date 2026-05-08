package utils

import "math/rand/v2"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.IntN(62)]
	}
	return string(b)
}
