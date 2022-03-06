package utils

/**
    utils
    @author: roccoshi
    @desc: 生成随机的Int, String
**/

import (
	"math/rand"
	"time"
)

// String 产生指定长度的随机字符串
func (*IRand) String(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(10 * time.Nanosecond)

	letter := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// Int [min, max]之间的随机int
func (*IRand) Int(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(10 * time.Nanosecond)
	return min + rand.Intn(max-min)
}
