package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

const (
	name = "get_baidu"
)

func main() {
	sync()
	// async()
}

func sync() {
	// 1. 配置 config
	conf := hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  10,   // command 的最大并发量，默认是 10
		SleepWindow:            5000, // 当熔断器被打开后，SleepWindow 的时间就是控制过多久后去尝试服务是否可用了。默认值是5000毫秒
		RequestVolumeThreshold: 20,   // 一个统计窗口10秒内请求数量。达到这个请求数量后才去判断是否要开启熔断。默认值是20
		ErrorPercentThreshold:  50,   // 错误百分比，请求数量大于等于RequestVolumeThreshold并且错误率到达这个百分比后就会启动熔断 默认值是50
	}
	// 2. 配置 command
	hystrix.ConfigureCommand("getProds", conf)

	datatest := []byte{}
	// 3. do 方法, 同步调用,
	/*
		第一个函数为正常的业务逻辑，第二个函数为 fallback，也就是外部系统挂了的时候执行一些动作
		当第一个函数返回error，或者在一系列健康检查的情况下函数无法运行结束，都会触发fallback
	*/
	err := hystrix.Do(name, func() error {
		// 正常业务逻辑，一般是访问其他资源
		res, err := http.Get("www.baidu.com")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		datatest = data
		return nil
	}, func(err error) error {
		// 5. 降级，显示默认值 (不用判断 err 传参)
		/*
			失败处理逻辑，访问其他资源失败时，或者处于熔断开启状态时，会调用这段逻辑
			可以简单构造一个response返回，也可以有一定的策略，比如访问备份资源
			也可以直接返回 err，这样不用和远端失败的资源通信，防止雪崩
		*/
		fmt.Println("xxx", err.Error())
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(datatest))
}

/*
	异步执行Go方法,内部实现是启动了一个gorouting，如果想得到自定义方法的数据，需要你传channel来处理数据，或者输出。
	返回的error也是一个channel
*/
func async() {
	output := make(chan []byte, 1)
	errors := hystrix.Go(name, func() error {
		res, err := http.Get("http://www.baidu.com")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		output <- data
		return nil
	}, nil)

	select {
	case out := <-output:
		fmt.Println("任务成功", string(out))
	case err := <-errors:
		fmt.Println("任务失败", err.Error())
	}
}
