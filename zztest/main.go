package main

import (
	"archive/tar"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/astaxie/beego"
	"github.com/robfig/cron"
)

const (
	MAXGOROUTINENUM = 10
)

var (
	deviceNumChan chan int
	wgStart       sync.WaitGroup
)

func main() {
	// test1()
	// test2()
	// test3()
	// DefaultValueOfStruct()
	// t()
	// t1()
	// ok := strings.HasSuffix("/Users/majun/sa6/test.tar", ".tar")
	// fmt.Println(ok)
	// testTar("test.tar")

	// p := getCurrentDirectory()
	// fmt.Println("path:", p)

	// rune1()

	// tSelect()

	// One("majun")
	// One("mamengli")
	// INTERVAL := 60 * 60

	// endTimestamp := Get0Timestamp()
	// startTimestamp := endTimestamp - 60*60*24
	// interval := int64(INTERVAL)
	// timeSlices := TimeSegments(startTimestamp, endTimestamp, interval)
	// for _, t := range timeSlices {
	// 	s := fmt.Sprintf("%s-%s", time.Unix(t[0], 0).Format("2006-01-02/15"), time.Unix(t[1], 0).Format("15"))
	// 	fmt.Println(s)
	// }

	// ConR()

	// a := &Animal{
	// 	Name: "cat",
	// 	Age:  10,
	// }
	// data, _ := json.Marshal(a)
	// fmt.Println(string(data))

	// Parse(tokenStr)

	// for i := 0; i < 10000; i++ {
	// 	go func() {
	// 		pub := []byte(publicKey)
	// 		pb, err := jwt.ParseRSAPublicKeyFromPEM(pub) //解析公钥
	// 		if err != nil {
	// 			fmt.Println("ParseRSAPublicKeyFromPEM:", err.Error())
	// 		}
	// 		fmt.Println(pb)
	// 	}()
	// }
	// todayDateStr := time.Now().Format("2006-01-02")

	// t, _ := time.ParseInLocation("2006-01-02", todayDateStr, time.Local)
	// fmt.Println(t.UnixNano() / 1e6)
	// count := 100
	// deviceNumChan = make(chan int, count)
	// for i := 0; i < count; i++ {
	// 	deviceNumChan <- i
	// }
	// close(deviceNumChan)

	// for k := 0; k < MAXGOROUTINENUM; k++ {
	// 	wgStart.Add(1)
	// 	go func() {
	// 		defer wgStart.Done()

	// 		for i := range deviceNumChan {
	// 			fmt.Println("======================", i)
	// 		}
	// 	}()
	// }
	// wgStart.Wait()

	// ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*1)
	// defer cancelFunc()
	// time.Sleep(2 * time.Second)
	// // 覆盖了
	// ctx2, cancelFunc2 := context.WithTimeout(ctx, time.Second*5)
	// defer cancelFunc2()
	// v, ok := ctx2.Deadline()
	// fmt.Println(v, ok)
	// fmt.Println(time.Now().Format(time.RFC3339))

	// time.Sleep(6 * time.Second)

	// cronJob1 := cron.New()
	// cronJob1.AddFunc("@every 1m", func() {
	// 	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 	fmt.Println("hahah")
	// })
	// cronJob1.Start()

	// cronJob2 := cron.New()
	// cronJob2.AddFunc("@every 10s", func() {
	// 	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 	fmt.Println("hahah")
	// })
	// cronJob2.Start()

	cronJob := cron.New()
	cronJob.AddFunc("0 42 21 * * *", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("hahah")
	})
	cronJob.Start()

	select {}
}

// 测试 omitempty
type Animal struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
	School string `json:"school,omitempty"`
}

type Project struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Docs string `json:"docs,omitempty"`
}

var value string

func ConR() {
	value = "mamengli"
	for i := 0; i < 10000; i++ {
		go func() {
			fmt.Println(value)
		}()
	}
	select {}
}

func Get0Timestamp() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix() - 60*60*8
	return timeNumber
}

func TimeSegments(startTimestamp, endTimestamp, interval int64) [][]int64 {
	var s []int64
	chunks := (endTimestamp - startTimestamp) / interval
	if (endTimestamp-startTimestamp)%interval > 0 {
		chunks++
	}
	timeSlices := make([][]int64, chunks)
	for i := 0; int64(i) < chunks; i++ {
		if int64(i) == chunks-1 {
			s = []int64{startTimestamp + interval*int64(i), endTimestamp}
			timeSlices[i] = s
			break
		}
		s = []int64{startTimestamp + interval*int64(i), startTimestamp + interval*int64(i+1)}
		timeSlices[i] = s
	}
	return timeSlices
}

var once sync.Once

func One(name string) {
	once.Do(func() {
		fmt.Println("this is ", name)
	})
}

func testTar(fpath string) {
	f, err := os.Open(fpath)
	if err != nil {
		fmt.Println("11111err:", err)
	}
	defer f.Close()
	tarRead := tar.NewReader(f)
	for {
		header, err := tarRead.Next()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("ERROR: cannot read tar file, error=[%v]\n", err)
			return
		}
		if header.FileInfo().IsDir() {
			continue
		}
		if strings.HasPrefix(filepath.Base(header.Name), ".") {
			continue
		}
	}
}

func t1() {
	fmt.Println("a" < "b")
	fmt.Println("2019-08-31" > "2019-08-31")
}

func test1() {
	state := "aHR0cDovL2xvY2FsaG9zdDo4MDgwL2YvIy90YXNr"
	url, er := base64.RawURLEncoding.DecodeString(state)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(string(url))

	t := time.Now().AddDate(0, -3, 0)
	verifyTime1 := t.Format("2006.01.02")
	fmt.Println(verifyTime1)

	var t1 int64 = 1562337405
	fmt.Println(time.Unix(t1, 0).Format("2006.01.02"))
}

func test2() {
	data, err := json.Marshal([]interface{}{"majun", "mamengli"})
	fmt.Println(data, err, len(data))

	params := []interface{}{}
	dec := json.NewDecoder(bytes.NewReader(data))
	err2 := dec.Decode(&params)
	fmt.Println(err2, params)
	fmt.Println(len(params))
}

type user struct {
	Name string
	Age  int
}

// 测试返回值为地址  error
func test3() (u *user) {
	u.Age = 19
	u.Name = "mamengli"
	return
}

// right
func test4() (u user) {
	u.Age = 19
	u.Name = "mamengli"
	return
}

type Person struct {
	Name string `defaultValue:"mengliam"`
	Age  int    `defaultValue:"21"`
}

func DefaultValueOfStruct() {
	p := Person{
		Name: "mamengli",
	}
	fmt.Printf("info:%v", p)
}

// time.Time 的零值
func t() {
	t := time.Time{}
	fmt.Println(t.IsZero())
}

func RootPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicln("发生错误", err.Error())
	}
	i := strings.LastIndex(s, "\\")
	path := s[0 : i+1]
	return path
}

func getCurrentPath() string {
	_, filename, _, ok := runtime.Caller(1)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "") // the the main function file directory
	} else {
		cwdPath = "./"
	}
	return cwdPath
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Debug(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func rune1() {
	a := "this is a new world, 马梦里"
	b := []rune(a)
	fmt.Println(b, len(b))
	fmt.Println([]byte(a), len(a))
}

func tSelect() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c

		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("recv quit!")
			return
		case syscall.SIGHUP:
		default:
			fmt.Println("111")
			return
		}
	}
}
