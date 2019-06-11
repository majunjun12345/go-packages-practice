package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	TimeFormat()
}

/*
	unix 和　unixNano
*/
func UnixTime() {
	t := time.Now()                            // 2019-06-11 22:17:44.282922603 +0800 CST m=+0.000105742
	fmt.Println(t.Unix())                      // 1560262709
	fmt.Println(t.UTC().Format(time.UnixDate)) // Tue Jun 11 14:19:11 UTC 2019

	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	fmt.Println(timestamp) //1560262815441619852  纳秒
	timestamp = timestamp[:10]
	fmt.Println(timestamp) // 1560262815　秒
}

/*
	时间格式化字符串转换
*/
func TimeFormat() {
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Jun 21, 2017 at 0:00am (PST)") // 2017-06-21 00:00:00 +0000 PST
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2017-Jun-21") // 2017-06-21 00:00:00 +0000 UTC
	fmt.Println(t)

	t, _ = time.Parse("01/02/2006", "06/21/2017")
	fmt.Println(t)
	fmt.Println(t.Unix())
}
