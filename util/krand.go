package util

import (
	"math/rand"
	"time"
)

const (
	// KRANDNUM 纯数字
	KRANDNUM = 0
	// KRANDLOWER 小写字母
	KRANDLOWER = 1
	// KRANDUPPER 大写字母
	KRANDUPPER = 2
	// KRANDALL 数字、大小写字母
	KRANDALL = 3
)

// Krand 随机字符串
func Krand(size int, kind int) []byte {
	kinds, result := [][2]int{[...]int{10, 48}, [...]int{26, 97}, [...]int{26, 65}}, make([]byte, size)
	isall := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isall { // random kind
			kind = rand.Intn(3)
		}
		scope, base := kinds[kind][0], kinds[kind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}
