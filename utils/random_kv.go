package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	randStr = rand.New(rand.NewSource(time.Now().Unix()))
	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

// GetTestKey 获取测试使用的 key
func GetTestKey(i int) []byte {
	return []byte(fmt.Sprintf("bitcask-%09d", i))
}

// GetTestKey 获取测试使用的 key
func GetTestKeyRandom() []byte {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 生成 unsigned int 范围的随机数
	randomNumber := rand.Uint32()
	return []byte(fmt.Sprintf("bitcask-%09d", randomNumber))
}

// RandomValue 生成随机 value，用于测试
func RandomValue(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[randStr.Intn(len(letters))]
	}
	return []byte("bitcask-go-value-" + string(b))
}
