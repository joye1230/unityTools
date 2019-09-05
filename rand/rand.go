package rand

import (
	"math/big"
	"math/rand"
	"time"
	crypt_rand "crypto/rand"
)

//随机N个不重复的数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	//|| (end-start) < count {
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

//生成N长度的随机数
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


//生成N长度的随机数
func RealRandInt64(number int64) int64 {
	randNumber := int64(0)
	n, err := crypt_rand.Int(crypt_rand.Reader, big.NewInt(number))
	if err == nil{
		randNumber = n.Int64()
	}

	return randNumber
}
