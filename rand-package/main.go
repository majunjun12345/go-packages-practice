package main

import (
	realRank "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"time"
)

/*
	math/rand: 伪随机
*/

func main() {
	// pseudoRand()

	// rRank()

	// fmt.Println(sessionId())

	// randInt()

}

// 两种写法
func pseudoRand() {
	// 设置随机种子，不然可能出现随机数一样的情况
	rand.Seed(time.Now().UnixNano())

	// 输出指定区间指定类型的随机数
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Int63n(100))

	rand1 := rand.New(rand.NewSource(time.Now().UnixNano())) // 非并发安全
	for i := 0; i < 5; i++ {
		fmt.Println(rand1.Int())
		fmt.Println(rand.Intn(100))
	}
}

func keng() {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().Unix())
		fmt.Printf("current:%d\n", time.Now().UnixNano())
		fmt.Println(rand.Intn(100)) // 每次循环的这两个值都相同，但是同批次的值不同
		fmt.Println(rand.Intn(100))
	}
}

/*
	crypto/rand 包中的随机数是利用当前系统的一些特征，比如内存的使用，文件的使用量，
	不同类型的进程数量等等来进行计算随机数，因此可能重复的几率很低。但是产生速度较慢;

	realRank.Reader 是一个全局的强随机生成器
*/
func rRank() {
	b := make([]byte, 10)
	n, err := realRank.Read(b)
	fmt.Println(n, err, b)
}

// 生成 sessionID
func sessionId() string {
	b := make([]byte, 32)
	_, err := io.ReadFull(realRank.Reader, b)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

// random int [0, 100)
func randInt() {
	b := new(big.Int).SetInt64(int64(100)) //将new(big.Int)设为int64(n)并返回new(big.Int)
	i, err := realRank.Int(realRank.Reader, b)
	fmt.Println(i, err)
}
