package main

import (
	"database/sql"
	"fmt"
	"runtime"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//Ping和Exec方法在调用完之后，会自动释放连接。没有返回值的函数不会自动释放链接，所以最好使用 exec 来执行 sql 语句
func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=true")
	CheckErr(err)
	defer db.Close()

	err = db.Ping() // open 并不会建立一个连接,只有当你使用的时候才会建立连接,所以这里需要提前 ping 一下,确保连接正常
	if err != nil {
		panic(err.Error())
	}

	// singleQuery(db) // 22s
	// bigQueryTest(db)// 264s

	// testParts()
	segmentQueryTest(db) // 100 41s, 1000 33, 10000 34
}

func connNumTest(db *sql.DB) {
	for i := 0; i < 10000; i++ {
		// var id int
		// db.QueryRow("select id from user where id=?", i).Scan(&id) // scan 后会自动 close，所以连接数一直是 1

		go func() {
			db.QueryRow("select id from testtable where id=?", i) // read tcp 127.0.0.1:57782->127.0.0.1:3306: read: connection reset by peer
		}()

	}
	time.Sleep(time.Second * 10)
}

func testParts() {
	total := 1005
	step := 100
	parts := total / step

	// segments := []int{}
	segments := [][]int{}
	for i := 0; i < parts; i++ {
		segments = append(segments, []int{step*i + 1, step * (i + 1)})
	}
	segments = append(segments, []int{step*parts + 1, total})
	fmt.Println(segments)

	works := make(chan []int, 1000)

	go func() {
		defer close(works)
		for i := 0; i < len(segments); i++ {
			works <- segments[i]
		}
	}()

	for i := range works {
		fmt.Println(i)
	}
}

func segmentQueryTest(db *sql.DB) {
	now := time.Now()
	// day, _ := time.ParseDuration("-24h")
	// LocalTimeFormat3 := "2006-01-02 15:04:05"
	// beginTime := now.Add(day * 7).Format(LocalTimeFormat3)
	// beginTime = "2010-01-01 10:10:24"

	// rows, err := db.Query("select id, username, password, age, sex, birth from user where birth = (select max(birth) from user where birth < ?)", beginTime)

	total := 1000000
	step := 1000
	parts := total / step
	segments := [][]int{}

	for i := 0; i < parts; i++ {
		segments = append(segments, []int{step*i + 1, step * (i + 1)})
	}
	segments = append(segments, []int{step*parts + 1, total})

	works := make(chan []int, 1000)
	wg := sync.WaitGroup{}

	go func() {
		defer close(works)
		for i := 0; i < len(segments); i++ {
			fmt.Println(segments[i])
			works <- segments[i]
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range works {
				rows, err := db.Query("select id, username, password, age, sex, birth from user where id>=? and id<=?", n[0], n[1])
				CheckErr(err)
				for rows.Next() {
					var id int
					var username string
					var password string
					var age int
					var sex int
					var birth time.Time
					err := rows.Scan(&id, &username, &password, &age, &sex, &birth)
					CheckErr(err)
					fmt.Println(id, username, password, age, sex, birth)
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(now))
}

func bigQueryTest(db *sql.DB) {
	now := time.Now()
	// day, _ := time.ParseDuration("-24h")
	// LocalTimeFormat3 := "2006-01-02 15:04:05"
	// beginTime := now.Add(day * 7).Format(LocalTimeFormat3)
	// beginTime = "2010-01-01 10:10:24"

	// rows, err := db.Query("select id, username, password, age, sex, birth from user where birth = (select max(birth) from user where birth < ?)", beginTime)

	works := make(chan int, 1000)
	wg := sync.WaitGroup{}

	go func() {
		defer close(works)
		for i := 1; i <= 1000000; i++ {
			works <- i
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range works {
				var id int
				var username string
				var password string
				var age int
				var sex int
				var birth time.Time
				err := db.QueryRow("select id, username, password, age, sex, birth from user where id=?", n).Scan(&id, &username, &password, &age, &sex, &birth)
				if err == sql.ErrNoRows {
					continue
				}
				CheckErr(err)
				fmt.Println(id, username, password, age, sex, birth)
			}
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(now))
}

func singleQuery(db *sql.DB) {
	now := time.Now()
	rows, err := db.Query("select id, username, password, age, sex, birth from user")
	CheckErr(err)

	for rows.Next() {
		var id int
		var username string
		var password string
		var age int
		var sex int
		var birth time.Time
		err := rows.Scan(&id, &username, &password, &age, &sex, &birth)
		CheckErr(err)
		fmt.Println(id, username, password, age, sex, birth)
	}
	fmt.Println(time.Since(now))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
