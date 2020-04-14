package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

func RandId() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	return Md5Encode(strconv.Itoa(int(nano + rand.Int63())))
}

// 获取md5
func Md5Encode(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
