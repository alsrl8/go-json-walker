package generator

import "math/rand"

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func randomValue() interface{} {
	switch rand.Intn(3) {
	case 0:
		return rand.Intn(1000)
	case 1:
		return randomString(rand.Intn(10) + 5)
	default:
		return rand.Intn(2) == 0
	}
}
