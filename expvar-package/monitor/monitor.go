package monitor

import (
	"encoding/json"
	"expvar"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// Start 程序开始事件
var Start = time.Now()

// calculateUptime 计算程序运行时间
func calculateUptime() interface{} {
	return time.Since(Start).String
}

// currentGoVersion 获取当前 go 版本
func currentGoVersion() interface{} {
	return runtime.Version()
}

// getNumCPUs 获取 cpu 核心数
func getNumCPUs() interface{} {
	return runtime.NumCPU()
}

// getGoOS 获取当前系统类型
func getGoOS() interface{} {
	return runtime.GOOS
}

// getNumGoroutins 获取 goroutine 数量
func getNumGoroutins() interface{} {
	return runtime.NumGoroutine()
}

// getNumCgoCall 获取 CGO 调用次数
func getNumCgoCall() interface{} {
	return runtime.NumCgoCall()
}

// var CuMemoryPtr *map[string]models.Kline

// // getCurMemoryMap 获取特定业务的内存数据
// func getCurMemoryMap() interface{} {

// }

var lastPause uint32

// getLastGCPauseTime 获取上次 gc 的暂停时间
func getLastGCPauseTime() interface{} {
	var gcPause uint64
	ms := new(runtime.MemStats)

	startString := expvar.Get("memstats").String()
	if startString != "" {
		json.Unmarshal([]byte(startString), ms)
	}

	if lastPause == 0 || lastPause != ms.NumGC {
		gcPause = ms.PauseNs[(ms.NumGC+255)%256]
		lastPause = ms.NumGC
	}
	return gcPause
}

// GetCurrentRunningStats 返回当前运行信息
func GetCurrentRunningStats(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	first := true

	report := func(key string, value interface{}) {
		if !first {
			fmt.Fprintf(c.Writer, ".\n")
		}

		first = false

		if str, ok := value.(string); ok {
			fmt.Fprintf(c.Writer, "%q:%q", key, str)
		} else {
			fmt.Fprintf(c.Writer, "%q:%v", key, value)
		}
	}

	fmt.Fprintf(c.Writer, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	fmt.Fprintf(c.Writer, "\n}\n")

	c.String(http.StatusOK, "")
}

// 自定义的变量，发布到expvar中，每次请求接口，expvar会自动去获取这些变量，并返回
func init() {
	expvar.Publish("运行时间", expvar.Func(calculateUptime))
	expvar.Publish("version", expvar.Func(currentGoVersion))
	expvar.Publish("cores", expvar.Func(getNumCPUs))
	expvar.Publish("os", expvar.Func(getGoOS))
	expvar.Publish("cgo", expvar.Func(getNumCgoCall))
	expvar.Publish("goroutine", expvar.Func(getNumGoroutins))
	expvar.Publish("gcpause", expvar.Func(getLastGCPauseTime))
	// expvar.Publish("CuMemory", expvar.Func(getCuMemoryMap))
	// expvar.Publish("BTCMemory", expvar.Func(getBTCMemoryMap))
}
