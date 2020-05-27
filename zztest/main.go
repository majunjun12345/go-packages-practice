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
	count := 100
	deviceNumChan = make(chan int, count)
	for i := 0; i < count; i++ {
		deviceNumChan <- i
	}
	close(deviceNumChan)

	for k := 0; k < MAXGOROUTINENUM; k++ {
		wgStart.Add(1)
		go func() {
			defer wgStart.Done()

			for i := range deviceNumChan {
				fmt.Println("======================", i)
			}
		}()
	}
	wgStart.Wait()
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

const (
	Root              = "IoT."
	common            = Root + "Common."
	system            = Root + "System."
	device            = Root + "Device."
	thing             = Root + "Thing."
	asset             = Root + "Asset."
	video             = Root + "Video."
	service           = Root + "Service."
	deviceGroup       = Root + "DeviceGroup."
	property          = Root + "Property."
	alarm             = Root + "Alarm."
	task              = Root + "Task."
	token             = Root + "Token."
	dataExportRecord  = Root + "DataExportRecord."
	objectStorageConf = Root + "DataExport."
	auth              = Root + "Auth."
)

const (
	NotExisted  = "NotExisted"
	NameExisted = "NameExisted"
	IdExisted   = "IdExisted"
	ConfExisted = "ConfExisted"
)

const (
	Success       = system + "Success"
	InternalError = system + "InternalError"
	SystemError   = system + "SystemError"
	RpcTimeOut    = system + "TimeOut"

	ParamsTooLong       = common + "ParamsTooLong"
	ResourceNotExisted  = common + "ResourceNotExisted"
	InternalParamsError = common + "InternalParamsError"
)
const (
	//device
	DeviceNotExisted                 = device + NotExisted
	DeviceNameExisted                = device + NameExisted
	DeviceNamePrefixExisted          = device + "NamePrefixExisted"
	DeviceInvalid                    = device + "Invalid"
	ParentDeviceNotExisted           = device + "ParentDevice" + NotExisted
	InvalidParentDevice              = device + "InvalidParentDevice"
	ExistSubDevice                   = device + "ExistSubDevice"
	DeviceStatusSwitchForbidden      = device + "DeviceStatusSwitchForbidden"
	InvalidDeviceStatus              = device + "InvalidDeviceStatus"
	ParentDeviceExistInSubDevices    = device + "ParentDeviceExistInSubDevices"
	DeviceTreeHeightExceed           = device + "DeviceTreeHeightExceed"
	DeviceAlreadyExistParent         = device + "DeviceAlreadyExistParent"
	UpdateDevicePropertyStatusFailed = device + "UpdateDevicePropertyStatusFailed"
	InvalidDeviceAccessProtocol      = device + "InvalidDeviceAccessProtocol"
	InvalidDeviceColumn              = device + "InvalidDeviceColumn"

	DeviceNameTooLong                  = device + "DeviceNameTooLong"
	TaskNamePrefixExisted              = device + "TaskNamePrefixExisted"
	CredentialHasExpired               = device + "CredentialHasExpired"
	RegisterFailedByTaskStatus         = device + "RegisterFailedByTaskStatus"
	TaskHasExpired                     = device + "TaskHasExpired"
	RegisterFailedByExceededTaskMaxNum = device + "RegisterFailedByExceededTaskMaxNum"
	ResourceIdentifierExisted          = device + "ResourceIdentifierExisted"
	DevicePropertyExisted              = device + "Property" + IdExisted

	//thing
	ThingNotExisted           = thing + NotExisted
	ThingNameExisted          = thing + NameExisted
	ThingIdExisted            = thing + IdExisted
	ThingAssociatedWithDevice = thing + "AssociatedWithDevice"
	ThingPropertyExisted      = thing + "Property" + IdExisted      //模型下属性标识符已存在
	ThingEventExisted         = thing + "Event" + IdExisted         //模型下事件标识符已存在
	ThingEventOutputExisted   = thing + "EventOutput" + IdExisted   //模型下事件输出参数标识已存在
	ThingServiceExisted       = thing + "Service" + IdExisted       //模型下服务标识符已存在
	ThingServiceOutputExisted = thing + "ServiceOutput" + IdExisted //模型下服务输出参数标识已存在
	ThingServiceInputExisted  = thing + "ServiceInput" + IdExisted  //模型下服务输入参数标识已存在
	ThingPropertyNotExisted   = thing + "Property" + NotExisted
	ThingExistDevice          = thing + "ThingExistDevice"

	// task
	TaskExpireTimeInvalid = task + "ExpireTimeInvalid"

	//video
	VideoCreateFail = video + "DeviceCreateFail"
	VideoDeleteFail = video + "DeviceDeleteFail"

	//alarm
	AlarmPolicyCreateFail = alarm + "PolicyCreateFail"

	//asset
	AssetNotExisted    = asset + NotExisted
	AssetNameExisted   = asset + NameExisted
	AssetMissingParams = asset + "MissingParams"
	//group
	CanNotDeleteRootDeviceGroup    = deviceGroup + "CanNotDeleteRootDeviceGroup"
	DeviceGroupNotExisted          = deviceGroup + NotExisted
	DeviceGroupNameExisted         = deviceGroup + NameExisted
	DeviceGroupIllegalResourceType = deviceGroup + "IllegalResourceType"

	//token
	TokenNotExisted = token + NotExisted

	//ServiceInfoExisted = service + "ServiceInfo" + IdExisted

	ParamsInvalid             = common + "ParamsInvalid"
	AccountError              = common + "AccountError"
	PitrixError               = common + "PitrixError"
	IAMError                  = common + "IAMError"
	IllegalAccessKey          = common + "IllegalAccessKey"
	AccessKeyNotActive        = common + "AccessKeyNotActive"
	UserNotFound              = common + "UserNotFound"
	UserNotFinishRegistration = common + "UserNotFinishRegistration"
	UserAccessDenied          = common + "UserAccessDenied"
	SuperUserOnly             = common + "SuperUserOnly"
	ResourceNotFound          = common + "ResourceNotFound"

	InvalidPropertyType = property + "InvalidPropertyType"
	InvalidAccessType   = property + "InvalidAccessType"

	// data-export
	DataExportRecordNotExisted  = dataExportRecord + NotExisted
	ObjectStorageConfNotExisted = objectStorageConf + NotExisted
	ObjectStorageConfExisted    = objectStorageConf + ConfExisted
	AuthNotExisted              = auth + NotExisted
)

var (
	Errors = map[string]string{
		InternalError:                      "内部错误",
		InternalParamsError:                "内部参数错误",
		SystemError:                        "系统错误",
		DeviceNotExisted:                   "设备不存在",
		DeviceNameExisted:                  "设备名称已存在",
		DeviceInvalid:                      "无效设备",
		ParentDeviceNotExisted:             "父设备不存在",
		InvalidParentDevice:                "无效父设备",
		ExistSubDevice:                     "存在子设备",
		InvalidDeviceStatus:                "无效设备状态",
		ParentDeviceExistInSubDevices:      "父设备存在于待添加的子设备列表",
		DeviceTreeHeightExceed:             "设备树高度超过限制",
		DeviceAlreadyExistParent:           "设备已存在父设备",
		UpdateDevicePropertyStatusFailed:   "更新设备属性状态失败",
		DeviceNameTooLong:                  "设备名称长度不能超过限制",
		ThingNotExisted:                    "模型不存在",
		ThingNameExisted:                   "模型名称已存在",
		ThingAssociatedWithDevice:          "不能删除已关联到设备的模型",
		ThingPropertyNotExisted:            "模型属性不存在",
		TaskNamePrefixExisted:              "批次名称前缀已存在",
		CredentialHasExpired:               "凭证已过期",
		RegisterFailedByTaskStatus:         "注册失败",
		TaskHasExpired:                     "任务已过期",
		RegisterFailedByExceededTaskMaxNum: "注册失败, 超过任务限制最大数量",
		ResourceIdentifierExisted:          "资源标识符已存在",
		AssetNotExisted:                    "资产信息不存在",
		DeviceGroupNameExisted:             "设备组已存在",
		DeviceGroupNotExisted:              "设备组不存在",
		CanNotDeleteRootDeviceGroup:        "不能删除根设备组",
		ParamsTooLong:                      "参数值长度超过限制",
		DeviceGroupIllegalResourceType:     "非法设备组资源类型",
		DeviceNamePrefixExisted:            "设备名称前缀已存在",
		ResourceNotFound:                   "资源未发现",
		AccountError:                       "请求账户错误",
		PitrixError:                        "云平台错误",
		IAMError:                           "请求IAM错误",
		AssetMissingParams:                 "缺少参数",
		ThingIdExisted:                     "模型Id已存在",
		//ServiceInfoExisted:                 "服务信息已存在",
		IllegalAccessKey:          "非法AK",
		AccessKeyNotActive:        "不可用的API密匙",
		UserNotFound:              "用户不存在",
		UserNotFinishRegistration: "用户未完成注册",
		UserAccessDenied:          "用户连接被拒绝",
		SuperUserOnly:             "只允许超级用户访问",
		ParamsInvalid:             "参数无效",
		ThingPropertyExisted:      "模型属性已存在",
		ThingEventExisted:         "模型下事件标识符已存在",    //模型下事件标识符已存在
		ThingServiceExisted:       "模型下服务标识符已存在",    //模型下服务标识符已存在
		ThingEventOutputExisted:   "模型下事件输出参数标识已存在", //模型下事件输出参数标识已存在
		ThingServiceOutputExisted: "模型下服务输出参数标识已存在", //模型下服务输出参数标识已存在
		ThingServiceInputExisted:  "模型下服务输入参数标识已存在", //模型下服务输入参数标识已存在

		ThingExistDevice:    "模型下已存在设备",
		VideoCreateFail:     "设备创建失败",
		InvalidDeviceColumn: "无效设备字段",
	}
)
